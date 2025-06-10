package utils

import (
	"fmt"
	"testing"
)

func TestGetPublicIP(t *testing.T) {
	// 测试不带 CIDR 的情况
	ip1, err := GetPublicIP(true)
	if err != nil {
		t.Errorf("获取公网 IP 失败（stripCidr=true）: %v", err)
	} else {
		fmt.Printf("TestGetPublicIP(stripCidr=true): %s\n", ip1)
	}

	// 测试带 CIDR 的情况
	ip2, err := GetPublicIP(false)
	if err != nil {
		t.Errorf("获取公网 IP 失败（stripCidr=false）: %v", err)
	} else {
		fmt.Printf("TestGetPublicIP(stripCidr=false): %s\n", ip2)
	}
}
