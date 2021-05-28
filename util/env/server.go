package env

type Server struct {
	Host Env
	Port Env
}

func NewServer() Server {
	e := Server{
		Host: Env{
			Key:   "SERVER_HOST",
			Value: "localhost",
		},
		Port: Env{
			Key:   "SERVER_PORT",
			Value: "8888",
		},
	}
	e.load()
	return e
}

func (e *Server) load() {
	e.Port.Load()
	e.Host.Load()
}
