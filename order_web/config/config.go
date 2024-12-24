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

type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key" yaml:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key" yaml:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url" yaml:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url" yaml:"return_url"`
}

type ServerConfig struct {
	GoodsSrvInfo     GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv" yaml:"goods_srv"`
	OrderSrvInfo     GoodsSrvConfig `mapstructure:"order_srv" json:"order_srv" yaml:"order_srv"`
	InventorySrvInfo GoodsSrvConfig `mapstructure:"inventory_srv" json:"inventory_srv" yaml:"inventory_srv"`
	Name             string         `mapstructure:"name" json:"name" yaml:"name"`
	Host             string         `mapstructure:"host" json:"host" yaml:"host"`
	PORT             int            `mapstructure:"port" json:"port" yaml:"port"`
	Tags             []string       `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo          JWTConfig      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ConsulInfo       ConsulConfig   `mapstructure:"consul" json:"consul" yaml:"consul"`
	AliPayInfo       AlipayConfig   `mapstructure:"alipay" json:"alipay" yaml:"alipay"`
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
