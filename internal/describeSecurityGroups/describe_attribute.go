package describeSecurityGroups

import (
	"fmt"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"strings"
)

// Attribute 查询指定区域的安全组属性
func Attribute(client *ecs20140526.Client, regionId, securityGroupId string) error {
	req := &ecs20140526.DescribeSecurityGroupAttributeRequest{
		RegionId:        tea.String(regionId),
		SecurityGroupId: tea.String(securityGroupId),
	}

	runtime := &service.RuntimeOptions{}

	resp, err := client.DescribeSecurityGroupAttributeWithOptions(req, runtime)
	if err != nil {
		return err
	}

	body := resp.Body
	if body == nil {
		return fmt.Errorf("response body is nil")
	}

	log.Println("入方向规则：")
	for _, perm := range body.Permissions.Permission {
		log.Printf("  - 协议: %s, 端口范围: %s, 来源IP: %s, 描述: %s\n",
			tea.StringValue(perm.IpProtocol),
			tea.StringValue(perm.PortRange),
			tea.StringValue(perm.SourceCidrIp),
			tea.StringValue(perm.Description),
		)
	}

	return nil
}

// GetStarredRules 查询安全组并返回描述以 * 开头的规则列表
func GetStarredRules(client *ecs20140526.Client,
	regionId,
	securityGroupId string) ([]*ecs20140526.DescribeSecurityGroupAttributeResponseBodyPermissionsPermission, error) {
	req := &ecs20140526.DescribeSecurityGroupAttributeRequest{
		RegionId:        tea.String(regionId),
		SecurityGroupId: tea.String(securityGroupId),
	}

	runtime := &service.RuntimeOptions{}
	resp, err := client.DescribeSecurityGroupAttributeWithOptions(req, runtime)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	if body == nil || body.Permissions == nil || body.Permissions.Permission == nil {
		return nil, fmt.Errorf("响应为空或无权限信息")
	}

	var result []*ecs20140526.DescribeSecurityGroupAttributeResponseBodyPermissionsPermission
	for _, perm := range body.Permissions.Permission {
		desc := tea.StringValue(perm.Description)
		if strings.HasPrefix(desc, "*") {
			result = append(result, perm)
		}
	}

	if len(result) == 0 {
		log.Printf("【%s】未找到带 * 描述的规则\n", securityGroupId)
		return result, nil
	}

	log.Printf("【%s】共找到 %d 条带 * 的规则：\n", securityGroupId, len(result))
	for _, rule := range result {
		log.Printf("  协议:%s 端口范围:%s 描述:%s RuleId: %s NicType: %s\n",
			tea.StringValue(rule.IpProtocol),
			tea.StringValue(rule.PortRange),
			tea.StringValue(rule.Description),
			tea.StringValue(rule.SecurityGroupRuleId), // 打印 RuleId
			tea.StringValue(rule.NicType),             // 打印 NicType
		)
	}

	return result, nil
}
