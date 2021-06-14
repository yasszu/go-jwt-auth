package env

type Variables struct {
	PostgresHost     Env
	PostgresPort     Env
	PostgresUser     Env
	PostgresPassword Env
	PostgresDB       Env
	ServerHost       Env
	ServerPort       Env
	JWTSecret        Env
}

func NewVariables() Variables {
	v := Variables{
		PostgresHost: Env{
			Key:   "POSTGRES_HOST",
			Value: "localhost",
		},
		PostgresPort: Env{
			Key:   "POSTGRES_PORT",
			Value: "5432",
		},
		PostgresUser: Env{
			Key:   "POSTGRES_USER",
			Value: "postgres",
		},
		PostgresPassword: Env{
			Key:   "POSTGRES_PASSWORD",
			Value: "password",
		},
		PostgresDB: Env{
			Key:   "POSTGRES_DB",
			Value: "postgres",
		},
		ServerHost: Env{
			Key:   "SERVER_HOST",
			Value: "localhost",
		},
		ServerPort: Env{
			Key:   "SERVER_PORT",
			Value: "8888",
		},
		JWTSecret: Env{
			Key:   "JWT_SECRET",
			Value: "your_secret_key",
		},
	}
	v.load()
	return v
}

func (v *Variables) load() {
	v.PostgresHost.Load()
	v.PostgresPort.Load()
	v.PostgresUser.Load()
	v.PostgresPassword.Load()
	v.PostgresDB.Load()
	v.ServerHost.Load()
	v.ServerPort.Load()
	v.JWTSecret.Load()
}
