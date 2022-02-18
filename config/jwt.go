package config

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" ini:"secret"`
	Issuer string `mapstructure:"issuer" json:"issuer" ini:"issuer"`
	Expire string `mapstructure:"expire" json:"expire" ini:"expire"`
}
