package pkg

import (
	"crypto"
	"fmt"
	"net/http"
	"strconv"
	"time"

	hmac "github.com/NaverCloudPlatform/ncloud-sdk-go-v2/hmac"
)

func GetCommonHeader(request *http.Request) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("cache-control", "max-age=0")
}

func SetAuthToken(request *http.Request, token string) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func SetNCPHeader(req *http.Request, accessKey, secretKey string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	signer := hmac.NewSigner(secretKey, crypto.SHA256)
	signature, _ := signer.Sign(http.MethodGet, req.URL.String(), accessKey, timestamp)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-ncp-apigw-timestamp", timestamp)
	req.Header.Set("x-ncp-iam-access-key", accessKey)
	req.Header.Set("x-ncp-apigw-signature-v2", signature)
}
