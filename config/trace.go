package config

type Trace struct {
	Servicename string `mapstructure:"servicename" json:"servicename" ini:"servicename"`
	Agenthost   string `mapstructure:"agenthost" json:"agenthost" ini:"agenthost"`
	Port        string `mapstructure:"port" json:"port" ini:"port"`
}
