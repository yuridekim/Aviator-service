package pkg_test

import (
	"github.com/cloud-club/Aviator-service/pkg"
	"testing"
)

func TestList(t *testing.T) {
	ncp := pkg.NewNcpService("ncp service token")
	ncp.Server = pkg.NewServerService("ncp server token")

	// given
	tests := []struct {
		name          string
		url           string
		payload       interface{}
		expectedError string
		expectedData  interface{}
	}{
		{
			name:          "성공",
			url:           "http://localhost:8080",
			payload:       nil,
			expectedError: "",
			expectedData:  nil,
		},
		{
			name:          "(실패) url을 입력 안함",
			url:           "",
			payload:       nil,
			expectedError: "please input url",
			expectedData:  nil,
		},
	}

	for i := range tests {
		t.Logf("%s : running scenario %d", tests[i].name, i)
		t.Run(tests[i].name, func(t *testing.T) {
			err := ncp.Server.List(tests[i].url)
			if err != nil {
				if err.Error() != tests[i].expectedError {
					t.Fatalf("expected error : %v, got : %v", tests[i].expectedError, err)
				}
			} else {
				if tests[i].expectedError != "" {
					t.Fatalf("expected error : %v, got : %v", tests[i].expectedError, err)
				}
			}
		})
	}

}
