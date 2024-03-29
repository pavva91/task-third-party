package config

var ServerConfigValues ServerConfig

// Model that links to config.yml file
type ServerConfig struct {
	Database struct {
		Connections int    `yaml:"connections" env:"DB_CONNECTIONS" env-description:"Total number of database connections"`
		Name        string `yaml:"name" env:"DB_NAME" env-description:"Database name"`
		Host        string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
		Password    string `yaml:"pass"  env:"DB_PASSWORD" env-description:"db password"`
		Port        string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
		Username    string `yaml:"user"  env:"DB_USERNAME" env-description:"db username"`
		Timezone    string `yaml:"timezone" env:"DB_TIMEZONE" env-description:"Database timezone"`
	} `yaml:"database"`
	Server struct {
		ApiPath            string   `yaml:"api-path"  env:"API_PATH" env-description:"API base path"`
		ApiVersion         string   `yaml:"api-version"  env:"API_VERSION" env-description:"API Version"`
		CorsAllowedClients []string `yaml:"cors-allowed-clients" env:"CORS_ALLOWED_CLIENTS"  env-description:"List of allowed CORS Clients"`
		Host               string   `yaml:"host"  env:"SERVER_HOST" env-description:"server host"`
		Port               string   `yaml:"port" env:"SERVER_PORT"  env-description:"server port"`
		Protocol           string   `yaml:"protocol" env:"SERVER_PROTOCOL"  env-description:"server protocol"`
	} `yaml:"server"`
}
