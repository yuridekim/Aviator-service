package pkg

type ServerInterface interface {
	Get(url string) error
	List(url string) error
	Create(url string, payload interface{}) error
	Delete(url string) error
}

func (client *NcpClient) Create(url string, payload interface{}) error {
	return nil
}

func (client *NcpClient) Get(url string) error {
	return nil
}

func (client *NcpClient) List(url string) error {
	return nil
}

func (client *NcpClient) Delete(url string) error {
	return nil
}
