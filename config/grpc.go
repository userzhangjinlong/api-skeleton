package config

type Grpc struct {
	Host string `mapstructure:"host" json:"host" ini:"host"`
	Port string `mapstructure:"Port" json:"Port" ini:"Port"`
}
