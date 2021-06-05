package env

type JWT struct {
	Secret Env
}

func NewJWT() JWT {
	e := JWT{Secret: Env{
		Key:   "JWT_SECRET",
		Value: "your_secret_key",
	}}
	e.load()
	return e
}

func (e *JWT) load() {
	e.Secret.Load()
}
