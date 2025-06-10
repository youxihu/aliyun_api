#!/bin/bash
# 加载相关函数
source "/home/sys_bash_send/sendIP.sh"

# 设置文件路径
IP_FILE="/home/youxihu/alarm/Get-publicIpAndModify/old_ip.txt"
MAX_RETRIES=3                      # 最大重试次数
RETRY_DELAY=2                      # 重试间隔(秒)
SG_MODIFY_PATH="/home/youxihu/alarm/Get-publicIpAndModify/sgModify"
E_TIME=$(date +%Y/%m/%d--%H:%M)


check_and_update_ip() {
    # 检查文件是否存在，如果不存在则创建
    [ ! -e "$IP_FILE" ] && touch "$IP_FILE"

    # 获取旧IP地址
    local old_ip
    old_ip=$(<"$IP_FILE")

    # 获取新IP地址（支持重试机制）
    local new_ip=""
    local retry_count=0

    echo "$E_TIME 正在获取公网IP地址..."
    while [ -z "$new_ip" ] && [ $retry_count -lt $MAX_RETRIES ]; do
        new_ip=$(curl -s icanhazip.com)
        if [ -z "$new_ip" ]; then
            retry_count=$((retry_count + 1))
            echo "  获取IP失败，正在重试 (${retry_count}/${MAX_RETRIES})..."
            sleep $RETRY_DELAY
        fi
    done

    # 检查最终是否获取到IP
    if [ -z "$new_ip" ]; then
        echo "[错误] 无法获取IP地址，请检查网络连接或服务是否可用"
        sendDingIP "### **告警通知: $IPIP**\n\
#### 状态: IP获取失败\n\
- 错误: 无法获取IP地址\n\
- 执行时间: $(date +'%Y-%m-%d %H:%M:%S')\n\
- 备注: 请检查网络连接或服务是否可用\n"
        return 1
    fi

    echo "当前公网IP: $new_ip"
    echo "历史公网IP: ${old_ip:-无记录}"

    # 检查IP是否变化
    if [ "$old_ip" != "$new_ip" ]; then
        echo "检测到IP变化，正在更新..."
        echo "$new_ip" > "$IP_FILE"

        # 执行安全组修改
        echo "正在修改安全组规则..."
        local modify_output
        modify_output=$("$SG_MODIFY_PATH" 2>&1 | tee -a sgModify.log)
        local modify_success
        modify_success=$(grep -w 'ModifySecurityGroupsRules=True' <<< "$modify_output")

        # 发送通知
        if [ "$modify_success" = "ModifySecurityGroupsRules=True" ]; then
            echo "安全组规则修改成功"
            sendDingIP "### **告警通知: $IPIP**\n\
### 状态: 已处理\n\
- IP地址: $new_ip\n\
- 执行时间: $(date +'%Y-%m-%d %H:%M:%S')\n\
- 备注：已通过自动化运维系统修改安全组\n\
- [登录阿里云查看修改情况](https://ecs.console.aliyun.com/home#/)"
        else
            echo "[警告] 安全组规则修改失败"
            sendDingIP "### **告警通知: $IPIP**\n\
#### 状态: 待处理\n\
- IP地址: $new_ip\n\
- 执行时间: $(date +'%Y-%m-%d %H:%M:%S')\n\
- 备注: 自动化修改安全组出现错误\n\
- [请及时登录阿里云或查看自动化运维系统运行日志](https://ecs.console.aliyun.com/home#/)"
        fi
    else
        echo "IP地址无变化，无需处理"
    fi
}

main() {
    check_and_update_ip
}

main
