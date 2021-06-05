package env

type Postgres struct {
	Host     Env
	Port     Env
	User     Env
	Password Env
	DB       Env
}

func NewPostgres() Postgres {
	e := Postgres{
		Host: Env{
			Key:   "POSTGRES_HOST",
			Value: "localhost",
		},
		Port: Env{
			Key:   "POSTGRES_PORT",
			Value: "5432",
		},
		User: Env{
			Key:   "POSTGRES_USER",
			Value: "postgres",
		},
		Password: Env{
			Key:   "POSTGRES_PASSWORD",
			Value: "password",
		},
		DB: Env{
			Key:   "POSTGRES_DB",
			Value: "postgres",
		},
	}
	e.load()
	return e
}

func (e *Postgres) load() {
	e.Host.Load()
	e.Port.Load()
	e.Password.Load()
	e.User.Load()
	e.DB.Load()
}
