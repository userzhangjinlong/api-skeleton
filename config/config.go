package config

type System struct {
	App      App      `json:"app" ini:"app"`
	Database Database `json:"database" ini:"database"`
	Redis    Redis    `json:"redis" ini:"redis"`
	Log      Log      `json:"log" ini:"log"`
	Proxy    Proxy    `json:"proxy" ini:"proxy"`
	Trace    Trace    `json:"trace" ini:"trace"`
}

type App struct {
	Debug string `mapstructure:"debug" json:"debug" ini:"debug"`
}

type Database struct {
	Charset      string `mapstructure:"charset" json:"charset" ini:"charset"`
	MaxIdleConns string `mapstructure:"maxidleconns" json:"maxidleconns" ini:"maxidleconns"`
	MaxOpenConns string `mapstructure:"maxopenconns" json:"maxopenconns" ini:"maxopenconns"`
	//默认Default库
	Host     string `mapstructure:"host" json:"host" ini:"host"`
	Port     string `mapstructure:"port" json:"port" ini:"port"`
	Name     string `mapstructure:"name" json:"name" ini:"name"`
	Username string `mapstructure:"username" json:"username" ini:"username"`
	Password string `mapstructure:"password" json:"password" ini:"password"`
	//information_schema库
	HostSchema     string `mapstructure:"hostschema" json:"hostschema" ini:"hostschema"`
	PortSchema     string `mapstructure:"portschema" json:"portschema" ini:"portschema"`
	NameSchema     string `mapstructure:"nameschema" json:"nameschema" ini:"nameschema"`
	UsernameSchema string `mapstructure:"usernameschema" json:"usernameschema" ini:"usernameschema"`
	PasswordSchema string `mapstructure:"passwordschema" json:"passwordschema" ini:"passwordschema"`
}

type Redis struct {
	Root string `mapstructure:"root" json:"root" ini:"root"`
	Auth string `mapstructure:"auth" json:"auth" ini:"auth"`
	Port string `mapstructure:"port" json:"port" ini:"port"`
	Db   string `mapstructure:"db" json:"db" ini:"db"`
}

type Log struct {
	Path string `mapstructure:"path" json:"path" ini:"path"`
	//LogFile bool   `mapstructure:"log-file" json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
}

type Proxy struct {
	Port       string `mapstructure:"port" json:"port" ini:"port"`
	TrustProxy string `mapstructure:"trustProxy" json:"trustProxy" ini:"trustProxy"`
}

type Trace struct {
	Servicename string `mapstructure:"servicename" json:"servicename" ini:"servicename"`
	Agenthost   string `mapstructure:"agenthost" json:"agenthost" ini:"agenthost"`
	Port        string `mapstructure:"port" json:"port" ini:"port"`
}
