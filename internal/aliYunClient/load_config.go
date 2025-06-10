package aliYunClient

import (
	"aliyun_api/internal/str"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(path string) (*str.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg str.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
