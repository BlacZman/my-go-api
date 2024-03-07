package services

import "os"

const (
	DSN         string = "DSN"
	DB_HOST     string = "DB_HOST"
	DB_USER            = "DB_USER"
	DB_PASSWORD        = "DB_PASSWORD"
	DB_NAME            = "DB_NAME"
	DB_PORT            = "DB_PORT"
	DB_SSL_MODE        = "DB_SSL_MODE"
	DB_TIMEZONE        = "DB_TIMEZONE"
)

type DatabaseConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	TimeZone string
}

type AppConfigService struct {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn string
	db  DatabaseConfig
}

func NewAppConfigService() AppConfigService {
	return AppConfigService{
		dsn: HandleEnvironmentVariable(DSN, ""),
		db: DatabaseConfig{
			host:     HandleEnvironmentVariable(DB_HOST, ""),
			user:     HandleEnvironmentVariable(DB_USER, ""),
			password: HandleEnvironmentVariable(DB_PASSWORD, ""),
			dbname:   HandleEnvironmentVariable(DB_NAME, ""),
			port:     HandleEnvironmentVariable(DB_PORT, ""),
			sslmode:  HandleEnvironmentVariable(DB_SSL_MODE, ""),
			TimeZone: HandleEnvironmentVariable(DB_TIMEZONE, ""),
		},
	}
}

func HandleEnvironmentVariable(env string, defaultValue string) string {
	value, err := os.LookupEnv(env)
	if err {
		return defaultValue
	}

	return value
}
