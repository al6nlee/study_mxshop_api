package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	Expire   int    `mapstructure:"expire" json:"expire"`
}

type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key" json:"key"`
	ApiSecrect string `mapstructure:"secrect" json:"secrect"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	UserSrv    UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	Name       string        `mapstructure:"name" json:"name"`
	Host       string        `mapstructure:"host" json:"host"`
	PORT       int           `mapstructure:"port" json:"port"`
	Tags       []string      `mapstructure:"tags" json:"tags"`
	JWTInfo    JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSmsInfo AliSmsConfig  `mapstructure:"sms" json:"sms"`
	RedisInfo  RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo ConsulConfig  `mapstructure:"consul" json:"consul"`
}
