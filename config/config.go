package config

type System struct {
	App          App          `json:"app" ini:"app"`
	Database     Database     `json:"database" ini:"database"`
	Redis        Redis        `json:"redis" ini:"redis"`
	Log          Log          `json:"log" ini:"log"`
	Proxy        Proxy        `json:"proxy" ini:"proxy"`
	WebSocket    WebSocket    `json:"web-socket" int:"web-socket"`
	Trace        Trace        `json:"trace" ini:"trace"`
	Jwt          Jwt          `json:"jwt" ini:"jwt"`
	RedisCluster RedisCluster `json:"redis-cluster" ini:"redis-cluster"`
	Nsq          Nsq          `json:"nsq" ini:"nsq"`
	Kafka        Kafka        `json:"kafka" ini:"kafka"`
	Rabbitmq     Rabbitmq     `json:"rabbitmq" ini:"rabbitmq"`
	Grpc         Grpc         `json:"grpc" ini:"grpc"`
	Etcd         Etcd         `json:"etcd" ini:"etcd"`
}

type App struct {
	Debug string `mapstructure:"debug" json:"debug" ini:"debug"`
}

type Database struct {
	Charset      string `mapstructure:"charset" json:"charset" ini:"charset"`
	MaxIdleConns string `mapstructure:"maxidleconns" json:"maxidleconns" ini:"maxidleconns"`
	MaxOpenConns string `mapstructure:"maxopenconns" json:"maxopenconns" ini:"maxopenconns"`
	//默认Default库 用户中心
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
	//IM库
	HostIm     string `mapstructure:"hostim" json:"hostim" ini:"hostim"`
	PortIm     string `mapstructure:"portim" json:"portim" ini:"portim"`
	NameIm     string `mapstructure:"nameim" json:"nameim" ini:"nameim"`
	UserNameIm string `mapstructure:"usernameim" json:"usernameim" ini:"usernameim"`
	PasswordIm string `mapstructure:"passwordim" json:"passwordim" ini:"passwordim"`
	//IM读库
	HostImRead     string `mapstructure:"hostimread" json:"hostimread" ini:"hostimread"`
	PortImRead     string `mapstructure:"portimread" json:"portimread" ini:"portimread"`
	NameImRead     string `mapstructure:"nameimread" json:"nameimread" ini:"nameimread"`
	UserNameImRead string `mapstructure:"usernameimread" json:"usernameimread" ini:"usernameimread"`
	PasswordImRead string `mapstructure:"passwordimread" json:"passwordimread" ini:"passwordimread"`
}

type Redis struct {
	Root string `mapstructure:"root" json:"root" ini:"root"`
	Auth string `mapstructure:"auth" json:"auth" ini:"auth"`
	Port string `mapstructure:"port" json:"port" ini:"port"`
	Db   string `mapstructure:"db" json:"db" ini:"db"`
}

type RedisCluster struct {
	Root      string `mapstructure:"root" json:"root" ini:"root"`
	Auth      string `mapstructure:"auth" json:"auth" ini:"auth"`
	PortOne   string `mapstructure:"portone" json:"portone" ini:"portone"`
	PortTwo   string `mapstructure:"porttwo" json:"porttwo" ini:"porttwo"`
	PortThree string `mapstructure:"portthree" json:"portthree" ini:"portthree"`
	PortFour  string `mapstructure:"portfour" json:"portfour" ini:"portfour"`
	PortFive  string `mapstructure:"portfive" json:"portfive" ini:"portfive"`
	PortSix   string `mapstructure:"portsix" json:"portsix" ini:"portsix"`
}

type Log struct {
	Path string `mapstructure:"path" json:"path" ini:"path"`
	//LogFile bool   `mapstructure:"log-file" json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
}

type Proxy struct {
	Port       string `mapstructure:"port" json:"port" ini:"port"`
	TrustProxy string `mapstructure:"trustProxy" json:"trustProxy" ini:"trustProxy"`
}

type WebSocket struct {
	Port string `mapstructure:"port" json:"port" ini:"port"`
	Host string `mapstructure:"host" json:"host" ini:"host"`
}
