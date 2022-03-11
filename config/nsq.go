package config

type Nsq struct {
	Host       string `mapstructure:"host" json:"host" ini:"host"`
	Node1      string `mapstructure:"node1" json:"node1" ini:"node1"`
	Node2      string `mapstructure:"node2" json:"node2" ini:"node2"`
	Node3      string `mapstructure:"node3" json:"node3" ini:"node3"`
	LookUpNode string `mapstructure:"loolupnode" json:"loolupnode" ini:"loolupnode"`
	Password   string `mapstructure:"password" json:"password" ini:"password"`
}
