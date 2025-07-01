package authorizeSecurityGroups

import (
	"aliyun_api/internal/aliYunClient"
	"fmt"
	"log"
	"testing"
)

const (
	ConfigPath = "/home/youxihu/secret/aiops/aliyun_sg/sg.yaml"
)

func TestAuthorizeSecurityGroups(t *testing.T) {

	// 加载配置
	cfg, err := aliYunClient.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	accessKeyId := cfg.BbxAccessKeyID
	accessKeySecret := cfg.BbxAccessSecret
	regionId := cfg.BbxRegionID
	// 添加规则
	err = AddSecurityGroupRule(
		accessKeyId,
		accessKeySecret,
		regionId,
		"sg-bp180i48w441ojbrndk3",
		"192.168.0.0/24",
		"22/22",
		"tcp",
		"Accept",
	)
	if err != nil {
		fmt.Println("添加规则失败:", err)
	}
}
