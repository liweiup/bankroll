package config

type Else struct {
	WxAppid   string `mapstructure:"wx-appid" json:"wx-appid" yaml:"wx-appid"`
	WxSecret  string `mapstructure:"wx-secret" json:"wx-secret" yaml:"wx-secret"`
	WxPreurl  string `mapstructure:"wx-preurl" json:"wx-preurl" yaml:"wx-preurl"`
}

