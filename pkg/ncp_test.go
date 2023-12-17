package pkg_test

import (
	"errors"
	"fmt"
	"github.com/cloud-club/Aviator-service/pkg"
	"github.com/cloud-club/Aviator-service/types/auth"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateAndVerifyToken_success(t *testing.T) {
	ncp := pkg.NewNcpService("ncp service token")
	ncp.Server = pkg.NewServerService("ncp server token")

	err := ncp.CreateToken("admin", "CloudClubAdmin", []string{"admin"})
	if err != nil {
		t.Fatal(err)
	}
	isSuccess, err := ncp.VerifyToken(ncp.GetToken(), &auth.AuthTokenClaims{})
	if err != nil {
		t.Fatal(err)
	}

	if isSuccess {
		t.Log(ncp.GetToken())
	} else {
		t.Fatal(errors.New(fmt.Sprintf("token verify failed : %s", ncp.GetToken())))
	}
}

func TestCreateAndVerifyToken_fail_token_is_invalid(t *testing.T) {
	ncp := pkg.NewNcpService("ncp service token")
	ncp.Server = pkg.NewServerService("ncp server token")

	err := ncp.CreateToken("admin", "CloudClubAdmin", []string{"admin"})
	if err != nil {
		t.Fatal(err)
	}
	isSuccess, err := ncp.VerifyToken(ncp.GetToken()+"is failed", &auth.AuthTokenClaims{})
	log.Println(err)
	if isSuccess {
		t.Fail()
	} else {
		assert.Error(t, err, errors.New("token signature is invalid"))
	}
}
