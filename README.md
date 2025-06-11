# 动态公网 IP 变更导致安全组规则失效的解决方案

## 项目背景
在使用阿里云 ECS（弹性计算服务）时，由于动态公网 IP 的特性，公司的公网 IP 地址可能会频繁变更。这种变更会导致安全组中的入方向规则失效，进而影响业务的正常运行。

为了解决这一问题，开发这个自动化工具，通过调用阿里云 ECS 的 OpenAPI 接口，实现对安全组规则的实时更新，确保业务不因网络问题中断、停止。

---
## 项目已实现(待拓展)

- [查询安全组基本信息列表](https://next.api.aliyun.com/document/Ecs/2014-05-26/DescribeSecurityGroups)
- [查询安全组及其组内规则信息](https://next.api.aliyun.com/document/Ecs/2014-05-26/DescribeSecurityGroupAttribute)
- `根据自定义需求`[修改安全组入方向规则](https://next.api.aliyun.com/document/Ecs/2014-05-26/ModifySecurityGroupRule)
- 实时获取公司公网 IP 并更新至安全组规则中
- 根据回调信息实时发送事件通知与告警

---

## 技术栈

- **语言**：`Go、Shell`
- **API**：更多需求可根据[阿里云API](https://next.api.aliyun.com/document/Ecs/2014-05-26)自行调用
- **本项目使用方法参阅**：[Makefile](https://github.com/youxihu/aliyun_api/blob/master/Makefile)

---

## 实现示意图
![自动化流程完成 示意图](https://img.picui.cn/free/2025/06/11/68493eb53c393.png)