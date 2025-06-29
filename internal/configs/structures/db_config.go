package structures

type DbConfig struct {
	Host               string `mapstructure:"DB_HOST"`
	User               string `mapstructure:"DB_USER"`
	Password           string `mapstructure:"DB_PASSWORD"`
	Db                 string `mapstructure:"DB_NAME"`
	Port               string `mapstructure:"DB_PORT"`
	MaxOpenConnections int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	Logging            bool   `mapstructure:"DB_LOGGING"`
}
