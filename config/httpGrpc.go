package config

type HttpGrpc struct {
	Host string `mapstructure:"host" json:"host" ini:"host"`
	Port string `mapstructure:"port" json:"port" ini:"port"`
}
