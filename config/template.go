package config

type Template struct {
	App struct {
		Env   string `mapstructure:"env"`
		Debug bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	DB struct {
		Host             string `mapstructure:"host"`
		Port             int    `mapstructure:"port"`
		User             string `mapstructure:"user"`
		Password         string `mapstructure:"password"`
		ConnectionString string `mapstructure:"connection_string"`
		MaxOpenConns     int    `mapstructure:"max_open_connections"`
		MaxIdleConns     int    `mapstructure:"max_idle_connections:"`
	} `mapstructure:"db"`
}
