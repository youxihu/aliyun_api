package str

type Config struct {
	// bbz account
	BbzAccessKeyID    string   `yaml:"bbz_access_key_id"`
	BbzAccessSecret   string   `yaml:"bbz_access_secret"`
	BbzRegionID       string   `yaml:"bbz_region_id"`
	BbzEndpoint       string   `yaml:"bbz_endpoint"`
	BbzSecurityGroups []string `yaml:"bbz_security_groups"`

	// bbx account (main region)
	BbxAccessKeyID    string   `yaml:"bbx_access_key_id"`
	BbxAccessSecret   string   `yaml:"bbx_access_secret"`
	BbxRegionID       string   `yaml:"bbx_region_id"`
	BbxEndpoint       string   `yaml:"bbx_endpoint"`
	BbxSecurityGroups []string `yaml:"bbx_security_groups"`

	// bbx account (hk region)
	BbxHKRegionID       string   `yaml:"bbx_hk_region_id"`
	BbxHKEndpoint       string   `yaml:"bbx_hk_endpoint"`
	BbxHKSecurityGroups []string `yaml:"bbx_hk_security_groups"`
}
