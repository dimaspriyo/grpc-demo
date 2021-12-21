package database

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	Schme    string `mapstructure:"schema"`
}
