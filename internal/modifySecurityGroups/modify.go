package modifySecurityGroups

import (
	"fmt"
	"log"
	"strings"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	tea "github.com/alibabacloud-go/tea/tea"

	ds "aliyun_api/internal/describeSecurityGroups"
)

// ModifyStarredRules 修改所有描述以 * 开头的规则，设置 SourceCidrIp 为本机公网 IP
func ModifyStarredRules(client *ecs20140526.Client, regionId, sgId string, publicIP string) error {
	var allErrors []string // 收集所有错误信息

	rules, err := ds.GetStarredRules(client, regionId, sgId)
	if err != nil {
		return fmt.Errorf("【%s】查询安全组规则失败：%v", sgId, err)
	}

	if len(rules) == 0 {
		fmt.Printf("【%s】未找到带 * 描述的规则，跳过修改\n", sgId)
		return nil
	}

	for _, rule := range rules {
		currentSourceCidrIp := tea.StringValue(rule.SourceCidrIp)
		if currentSourceCidrIp == publicIP {
			log.Printf("【%s】规则 (%s/%s) 访问来源已是 %s，无需修改\n",
				sgId,
				tea.StringValue(rule.IpProtocol),
				tea.StringValue(rule.PortRange),
				publicIP,
			)
			continue
		}

		modifyReq := &ecs20140526.ModifySecurityGroupRuleRequest{
			RegionId:            tea.String(regionId),
			SecurityGroupId:     tea.String(sgId),
			SecurityGroupRuleId: rule.SecurityGroupRuleId,
			SourceCidrIp:        tea.String(publicIP),
		}

		runtime := &service.RuntimeOptions{}
		_, err := client.ModifySecurityGroupRuleWithOptions(modifyReq, runtime)
		if err != nil {
			errorMsg := fmt.Sprintf("【%s】修改规则 (%s/%s) 失败: %v",
				sgId,
				tea.StringValue(rule.IpProtocol),
				tea.StringValue(rule.PortRange),
				err,
			)
			allErrors = append(allErrors, errorMsg)
			continue
		}

		fmt.Printf("【%s】规则 (%s/%s)修改成功,访问来源: %s-->%s\n",
			sgId,
			tea.StringValue(rule.IpProtocol),
			tea.StringValue(rule.PortRange),
			tea.StringValue(rule.SourceCidrIp),
			publicIP,
		)
	}

	// 如果有任意一条错误，就返回错误
	if len(allErrors) > 0 {
		return fmt.Errorf("部分规则修改失败:\n%s", strings.Join(allErrors, "\n"))
	}

	return nil
}
