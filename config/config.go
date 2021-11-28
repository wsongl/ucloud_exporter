package config

type Config struct {
	Credential
	ItemMetrics
	ProductMetrics
}

type Credential struct {
	Access_key_id string
	Access_key_secret string
	Regions []map[string]string
}

type ItemMetrics struct {
	Items map[string][]map[string]string
}

type ProductMetrics struct {
	products []string
}

// 全局配置变量
var ConfigVar = &Config{}
