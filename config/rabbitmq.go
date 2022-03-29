package config

type Rabbitmq struct {
	Username string `mapstructure:"username" json:"username" ini:"username"`
	Password string `mapstructure:"password" json:"password" ini:"password"`
	Node1    string `mapstructure:"node1" json:"node1" ini:"node1"`
	Vhost    string `mapstructure:"vhost" json:"vhost" ini:"vhost"`
}
