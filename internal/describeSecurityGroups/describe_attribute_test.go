// describe_test.go
package describeSecurityGroups

import (
	"aliyun_api/internal/aliYunClient"
	"fmt"
	"log"
	"testing"
)

const (
	ConfigPath = "/home/youxihu/secret/aiops/aliyun_sg/test.sg.yaml"
)

func TestDescribeSecurityGroupAttribute(t *testing.T) {
	cfg, err := aliYunClient.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	accounts := aliYunClient.GetAccounts(cfg)

	for accountName, account := range accounts {
		fmt.Printf("账号：%s，区域：%s\n", accountName, account.RegionId)

		for _, sgId := range account.SecurityGroups {
			runDescribeAttributeTest(
				account.Endpoint,
				account.RegionId,
				account.AccessKeyID,
				account.AccessSecret,
				sgId,
			)
		}
	}
}

func runDescribeAttributeTest(endpoint, regionId, ak, sk, sgId string) {
	client, err := aliYunClient.CreateEcsClient(endpoint, ak, sk)
	if err != nil {
		fmt.Printf("【%s】创建客户端失败: %v\n", sgId, err)
		return
	}

	fmt.Printf("【%s】开始查询...\n", sgId)
	err = Attribute(client, regionId, sgId)
	if err != nil {
		fmt.Printf("【%s】查询失败: %v\n", sgId, err)
		return
	}
}

func TestGetStarredRules(t *testing.T) {
	cfg, err := aliYunClient.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	accounts := aliYunClient.GetAccounts(cfg)

	for accountName, account := range accounts {
		fmt.Printf("账号：%s，区域：%s\n", accountName, account.RegionId)

		for _, sgId := range account.SecurityGroups {
			runGetStarredRulesTest(
				account.Endpoint,
				account.RegionId,
				account.AccessKeyID,
				account.AccessSecret,
				sgId,
			)
		}
	}
}

func runGetStarredRulesTest(endpoint, regionId, ak, sk, sgId string) {
	client, err := aliYunClient.CreateEcsClient(endpoint, ak, sk)
	if err != nil {
		fmt.Printf("【%s】创建客户端失败: %v\n", sgId, err)
		return
	}
	_, err = GetStarredRules(client, regionId, sgId)
	if err != nil {
		fmt.Printf("【%s】查询带 * 的规则失败：%v\n", sgId, err)
		return
	}
}
