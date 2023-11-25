package pkg

type NcpService struct {
	token  string
	Server ServerInterface
}

func NewNcpService(token string) *NcpService {
	return &NcpService{token: token}
}

func (n *NcpService) GetToken() string {
	return n.token
}
