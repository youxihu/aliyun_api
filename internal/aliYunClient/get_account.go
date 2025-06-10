package aliYunClient

import "aliyun_api/internal/str"

type Account struct {
	AccessKeyID    string
	AccessSecret   string
	SecurityGroups []string
	RegionId       string
	Endpoint       string
}

func newBBZAccount(cfg *str.Config) Account {
	return Account{
		AccessKeyID:    cfg.BbzAccessKeyID,
		AccessSecret:   cfg.BbzAccessSecret,
		SecurityGroups: cfg.BbzSecurityGroups,
		RegionId:       cfg.BbzRegionID,
		Endpoint:       cfg.BbzEndpoint,
	}
}

func newBBXAccount(cfg *str.Config) Account {
	return Account{
		AccessKeyID:    cfg.BbxAccessKeyID,
		AccessSecret:   cfg.BbxAccessSecret,
		SecurityGroups: cfg.BbxSecurityGroups,
		RegionId:       cfg.BbxRegionID,
		Endpoint:       cfg.BbxEndpoint,
	}
}

func newBBXHKAccount(cfg *str.Config) Account {
	return Account{
		AccessKeyID:    cfg.BbxAccessKeyID,
		AccessSecret:   cfg.BbxAccessSecret,
		SecurityGroups: cfg.BbxHKSecurityGroups,
		RegionId:       cfg.BbxHKRegionID,
		Endpoint:       cfg.BbxHKEndpoint,
	}
}

func GetAccounts(cfg *str.Config) map[string]Account {
	return map[string]Account{
		"bbz":    newBBZAccount(cfg),
		"bbx":    newBBXAccount(cfg),
		"bbx-hk": newBBXHKAccount(cfg),
	}
}
