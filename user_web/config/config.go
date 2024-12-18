package config

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Expire   int    `mapstructure:"expire" json:"expire" yaml:"expire"`
}

type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key" json:"key" yaml:"key"`
	ApiSecrect string `mapstructure:"secrect" json:"secrect" yaml:"secrect"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type ServerConfig struct {
	UserSrv    UserSrvConfig `mapstructure:"user_srv" json:"user_srv" yaml:"user_srv"`
	Name       string        `mapstructure:"name" json:"name" yaml:"name"`
	Host       string        `mapstructure:"host" json:"host" yaml:"host"`
	PORT       int           `mapstructure:"port" json:"port" yaml:"port"`
	Tags       []string      `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo    JWTConfig     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	AliSmsInfo AliSmsConfig  `mapstructure:"sms" json:"sms" yaml:"sms"`
	RedisInfo  RedisConfig   `mapstructure:"redis" json:"redis" yaml:"redis"`
	ConsulInfo ConsulConfig  `mapstructure:"consul" json:"consul" yaml:"consul"`
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
