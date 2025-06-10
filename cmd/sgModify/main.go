package main

import (
	"aliyun_api/internal/aliYunClient"
	"aliyun_api/internal/modifySecurityGroups"
	"aliyun_api/utils"
	"fmt"
	"log"
)

const (
	ConfigPath = "/home/youxihu/secret/aiops/aliyun_sg/test.sg.yaml"
)

func main() {
	cfg, err := aliYunClient.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	accounts := aliYunClient.GetAccounts(cfg)

	var allSuccess = true // 标记是否全部成功

	for accountName, account := range accounts {
		log.Printf("正在处理账号：%s\n", accountName)
		if len(account.SecurityGroups) == 0 {
			log.Printf("【%s】未配置任何安全组，跳过处理\n", accountName)
			continue
		}
		for _, sgId := range account.SecurityGroups {
			success := runModify(
				account.Endpoint,
				account.RegionId,
				account.AccessKeyID,
				account.AccessSecret,
				sgId,
			)
			if !success {
				allSuccess = false
			}
		}
	}

	if allSuccess {
		fmt.Println("ModifySecurityGroupsRules=True")
	}
}

func runModify(endpoint, regionId, ak, sk, sgId string) bool {
	client, err := aliYunClient.CreateEcsClient(endpoint, ak, sk)
	if err != nil {
		log.Printf("【%s】创建客户端失败: %v\n", sgId, err)
		return false
	}
	publicIP, err := utils.GetPublicIP(true)
	if err != nil {
		log.Fatalf("【%s】获取本机公网 IP 失败：%v", sgId, err)
	}

	err = modifySecurityGroups.ModifyStarredRules(client, regionId, sgId, publicIP)
	if err != nil {
		log.Printf("【%s】修改失败: %v\n", sgId, err)
		return false
	}
	return true
}
