package config

type System struct {
	App      App      `json:"app" ini:"app"`
	Database Database `json:"database" ini:"database"`
	Redis    Redis    `json:"redis" ini:"redis"`
	Log      Log      `json:"log" ini:"log"`
	Proxy    Proxy    `json:"proxy" ini:"proxy"`
}

type App struct {
	Debug string `mapstructure:"debug" json:"debug" ini:"debug"`
}

type Database struct {
	Charset  string `mapstructure:"charset" json:"charset" ini:"charset"`
	Host     string `mapstructure:"host" json:"host" ini:"host"`
	Port     string `mapstructure:"port" json:"port" ini:"port"`
	Name     string `mapstructure:"name" json:"name" ini:"name"`
	Username string `mapstructure:"username" json:"username" ini:"username"`
	Password string `mapstructure:"password" json:"password" ini:"password"`
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
