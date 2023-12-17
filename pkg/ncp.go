package pkg

import (
	"errors"
	"github.com/cloud-club/Aviator-service/types/auth"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

const (
	JwtTokenExpireTime = time.Minute * 30
	SignKey            = "CloudClubAviator"
)

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

func (n *NcpService) CreateToken(username string, name string, role []string) error {
	claim := auth.AuthTokenClaims{
		Username: username,
		Name:     name,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "CloudClub",
			ExpiresAt: jwt.At(time.Now().Add(JwtTokenExpireTime)),
		},
	}

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &claim)
	signedAuthToken, err := authToken.SignedString([]byte(SignKey))
	if err != nil {
		return err
	}
	n.token = signedAuthToken
	return nil
}

func (n *NcpService) VerifyToken(token string, claim *auth.AuthTokenClaims) (bool, error) {
	t, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		if _, isOk := t.Method.(*jwt.SigningMethodHMAC); !isOk {
			return nil, errors.New("deny to provide JWT Token")
		}
		return []byte(SignKey), nil
	})
	return t.Valid, err
}
