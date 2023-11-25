package pkg

import (
	"fmt"
	"net/http"
)

func GetCommonHeader(request *http.Request) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("cache-control", "max-age=0")
}

func SetAuthToken(request *http.Request, token string) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}
