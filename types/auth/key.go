package auth

type KeyService struct {
	accessKey string
	secretKey string
}

func NewKeyService(accessKey string, secretKey string) *KeyService {
	return &KeyService{accessKey: accessKey, secretKey: secretKey}
}

func (k *KeyService) GetAccessKey() string {
	return k.accessKey
}

func (k *KeyService) GetSecretKey() string {
	return k.secretKey
}
