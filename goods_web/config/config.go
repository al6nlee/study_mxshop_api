package config

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type ServerConfig struct {
	GoodsSrv   GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv" yaml:"goods_srv"`
	Name       string         `mapstructure:"name" json:"name" yaml:"name"`
	Host       string         `mapstructure:"host" json:"host" yaml:"host"`
	PORT       int            `mapstructure:"port" json:"port" yaml:"port"`
	Tags       []string       `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo    JWTConfig      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ConsulInfo ConsulConfig   `mapstructure:"consul" json:"consul" yaml:"consul"`
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
