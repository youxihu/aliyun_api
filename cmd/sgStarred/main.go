package main

import (
	"aliyun_api/internal/aliYunClient"
	"aliyun_api/internal/describeSecurityGroups"
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

	for accountName, account := range accounts {
		log.Printf("账号：%s，区域：%s\n", accountName, account.RegionId)

		for _, sgId := range account.SecurityGroups {
			runGetStarredRules(
				account.Endpoint,
				account.RegionId,
				account.AccessKeyID,
				account.AccessSecret,
				sgId,
			)
		}
	}
}

func runGetStarredRules(endpoint, regionId, ak, sk, sgId string) {
	client, err := aliYunClient.CreateEcsClient(endpoint, ak, sk)
	if err != nil {
		fmt.Printf("【%s】创建客户端失败: %v\n", sgId, err)
		return
	}
	_, err = describeSecurityGroups.GetStarredRules(client, regionId, sgId)
	if err != nil {
		fmt.Printf("【%s】查询带 * 的规则失败：%v\n", sgId, err)
		return
	}
}
