package config

type Etcd struct {
	Host string `mapstructure:"host" json:"host" ini:"host"`
	Node string `mapstructure:"node" json:"node" ini:"node"`
}
