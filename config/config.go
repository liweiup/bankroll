package config

type Server struct {
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
}

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config.yaml"
	LayoutDate = "2006-01-02"
	LayoutTime = "2006-01-02 15:04:05"
	RedisKey = "BK:"
	HolidaySet = RedisKey+"HOLIDAY"
)
