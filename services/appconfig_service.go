package services

import "os"

const (
	DSN string = "DSN"
)

type AppConfigService struct {
}

func NewAppConfigService() AppConfigService {
	return AppConfigService{}
}

func HandleEnvironmentVariable(env string, defaultValue string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		return defaultValue
	}

	return value
}

// dsn := "host=localhost user=postgres password=postgres dbname=postgres port=9920 sslmode=disable TimeZone=Asia/Bangkok"
func (a AppConfigService) DSN() string {
	return HandleEnvironmentVariable(DSN, "Not setup env yet")
}
