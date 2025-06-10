// internal/aliYunClient/client.go
package aliYunClient

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

// CreateEcsClient 使用指定 AK 创建 ECS 客户端
func CreateEcsClient(endpoint, accessKeyID, accessSecret string) (*ecs20140526.Client, error) {
	cred, err := credentials.NewCredential(&credentials.Config{
		Type:            tea.String("access_key"),
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessSecret),
	})
	if err != nil {
		return nil, err
	}

	config := &client.Config{
		Credential: cred,
		Endpoint:   tea.String(endpoint),
	}

	return ecs20140526.NewClient(config)
}
