package config

type OssConfig struct {
	ApiKey      string `mapstructure:"key" json:"key" yaml:"key"`
	ApiSecrect  string `mapstructure:"secrect" json:"secrect" yaml:"secrect"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url" yaml:"callback_url"`
	UploadDir   string `mapstructure:"upload_dir" json:"upload_dir" yaml:"upload_dir"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name" yaml:"name"`
	Host       string       `mapstructure:"host" json:"host" yaml:"host"`
	PORT       int          `mapstructure:"port" json:"port" yaml:"port"`
	Tags       []string     `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul" yaml:"consul"`
	OssInfo    OssConfig    `mapstructure:"oss" json:"oss" yaml:"oss"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" yaml:"host"`
	Port      uint64 `mapstructure:"port" yaml:"port"`
	Namespace string `mapstructure:"namespace" yaml:"namespace"`
	User      string `mapstructure:"user" yaml:"user"`
	Password  string `mapstructure:"password" yaml:"password"`
	DataId    string `mapstructure:"dataid" yaml:"dataid"`
	Group     string `mapstructure:"group" yaml:"group"`
}
