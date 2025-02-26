package server

type Server interface {
	Start()
}

type server struct {
}

func NewServer() Server {
	return &server{}
}
