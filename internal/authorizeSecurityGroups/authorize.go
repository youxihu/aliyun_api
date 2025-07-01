package authorizeSecurityGroups

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"strings"
)

// AddSecurityGroupRule 添加一条安全组规则
func AddSecurityGroupRule(
	accessKeyId string,
	accessKeySecret string,
	regionId string,
	sgID string,
	sourceCidrIp string,
	portRange string,
	ipProtocol string,
	policy string,
) error {

	// 创建ECS客户端
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return fmt.Errorf("创建ECS客户端失败: %v", err)
	}

	// 构建请求
	request := ecs.CreateAuthorizeSecurityGroupRequest()
	request.Scheme = "https"
	request.SecurityGroupId = sgID
	request.RegionId = regionId

	// 设置权限规则
	permission := ecs.AuthorizeSecurityGroupPermissions{
		SourceCidrIp: sourceCidrIp,
		PortRange:    portRange,
		IpProtocol:   strings.ToUpper(ipProtocol),
		Policy:       strings.Title(strings.ToLower(policy)), // 确保是 Accept / Drop
	}
	request.Permissions = &[]ecs.AuthorizeSecurityGroupPermissions{permission}

	// 发送请求
	response, err := client.AuthorizeSecurityGroup(request)
	if err != nil {
		return fmt.Errorf("调用API失败: %v", err)
	}

	fmt.Printf("RequestId: %s\n", response.RequestId)
	fmt.Println("✅ 安全组规则添加成功。")
	return nil
}
