package pkg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// URL        = "vm.cloudclub.io"
	URL        = "http://175.45.204.141/createServerImage"
	ACCESS_KEY = "yFDf3Gt4byTEOu66SFaA"
)

func makeSignature() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampStr := fmt.Sprintf("%d", timestamp)

	secretKey := "QoISqp2DopwMUokLYWSpoxYhQXi0aR9luGaCNkVe"

	method := "GET"

	message := method + " " + URL + "\n" + timestampStr + "\n" + ACCESS_KEY
	messageBytes := []byte(message)
	secretKeyBytes := []byte(secretKey)

	hash := hmac.New(sha256.New, secretKeyBytes)
	hash.Write(messageBytes)
	signingKey := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return signingKey
}

func SetHeader() {
	url := URL
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
	accessKey := ACCESS_KEY
	signature := makeSignature()

	body := map[string]interface{}{
		"regionCode":                            "KR",
		"serverImageName":                       "test-1",
		"blockStorageList.1.order":              "0",
		"blockStorageList.1.snapshotInstanceNo": "1111",
	}

	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-ncp-apigw-timestamp", timestamp)
	req.Header.Set("x-ncp-iam-access-key", accessKey)
	req.Header.Set("x-ncp-apigw-signature-v2", signature)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func main() {
	SetHeader()
}
