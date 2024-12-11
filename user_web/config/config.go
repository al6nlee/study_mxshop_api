package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	UserSrv UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	Name    string        `mapstructure:"name" json:"name"`
	PORT    int           `mapstructure:"port" json:"port"`
}
