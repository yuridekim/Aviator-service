package pkg

type Server struct {
	Token string
}

//go:generate mockgen -destination=mocks/mock_Server.go -package=pkg github.com/pkg/server ServerInterface
type ServerInterface interface {
	Get(url string) error
	List(url string) error
	Create(url string, payload interface{}) error
	Update(url string) error
	Delete(url string) error
}

func (client *Server) Create(url string, payload interface{}) error {
	return nil
}

func (client *Server) Get(url string) error {
	return nil
}

func (client *Server) List(url string) error {
	return nil
}

func (client *Server) Delete(url string) error {
	return nil
}

func (client *Server) Update(url string) error {
	return nil
}
