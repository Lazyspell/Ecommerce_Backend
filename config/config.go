package config

type AppConfig struct {
	InProduction bool
	Port         int
	Env          string
	DB           struct {
		DbHost string
		DbName string
		DbUser string
		DbPass string
		DbPort string
		DbSSL  string
	}

	Jwt struct {
	}
}

type Application struct {
	config AppConfig
}
