package modifySecurityGroups

import (
	"aliyun_api/internal/aliYunClient"
	"aliyun_api/utils"
	"fmt"
	"testing"
)

const (
	ConfigPath = "/home/youxihu/secret/aiops/aliyun_sg/test.sg.yaml"
)

func TestModifySecurityGroup(t *testing.T) {
	cfg, err := aliYunClient.LoadConfig(ConfigPath)
	if err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	accounts := aliYunClient.GetAccounts(cfg)

	for accountName, account := range accounts {
		t.Logf("正在处理账号：%s", accountName)

		for _, sgId := range account.SecurityGroups {
			runModifyTest(
				account.Endpoint,
				account.RegionId,
				account.AccessKeyID,
				account.AccessSecret,
				sgId,
			)
		}
	}
}

func runModifyTest(endpoint, regionId, ak, sk, sgId string) {
	client, err := aliYunClient.CreateEcsClient(endpoint, ak, sk)
	if err != nil {
		fmt.Printf("【%s】创建客户端失败: %v\n", sgId, err)
		return
	}
	publicIP, err := utils.GetPublicIP(true)
	if err != nil {
		_ = fmt.Errorf("【%s】获取本机公网 IP 失败：%v", sgId, err)
	}
	err = ModifyStarredRules(client, regionId, sgId, publicIP)
	if err != nil {
		fmt.Printf("【%s】修改失败: %v\n", sgId, err)
		return
	}
}
