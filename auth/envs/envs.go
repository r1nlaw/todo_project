package envs

var ServerEnvs Envs

type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	JWT_SECRET        string
	AUTH_PORT         string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

// Инициализация значений ENV
func LoadEnvs() error {
	ServerEnvs.JWT_SECRET = ""

	ServerEnvs.POSTGRES_USER = "yura"
	ServerEnvs.POSTGRES_HOST = "localhost"
	ServerEnvs.AUTH_PORT = "9104"
	ServerEnvs.POSTGRES_PASSWORD = "yura"
	ServerEnvs.POSTGRES_PORT = "9103"
	ServerEnvs.POSTGRES_NAME = "postgres"
	ServerEnvs.POSTGRES_USE_SSL = "disable"

	return nil
}
