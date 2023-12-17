package pkg

import (
	"crypto"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	hmac "github.com/NaverCloudPlatform/ncloud-sdk-go-v2/hmac"
)

const (
	GET_URL    = "https://ncloud.apigw.ntruss.com/vserver/v2/getAccessControlGroupList"
	URL        = "https://ncloud.apigw.ntruss.com/vserver/v2/createServerInstances"
	ACCESS_KEY = "b7e6Eq3fmVMGKBCCSLbi"
	SECRET_KEY = "S6ewbCjNSCk5kQLRDHvqXDGPqTUDwDn2LLhmIKma"
)

func SetHeader(req *http.Request) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	signer := hmac.NewSigner(SECRET_KEY, crypto.SHA256)
	signature, _ := signer.Sign("GET", req.URL.String(), ACCESS_KEY, timestamp)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-ncp-apigw-timestamp", timestamp)
	req.Header.Set("x-ncp-iam-access-key", ACCESS_KEY)
	req.Header.Set("x-ncp-apigw-signature-v2", signature)
}

func TestCreate() {
	url := URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	q := req.URL.Query()
	q.Add("vpcNo", "47526")
	q.Add("subnetNo", "106268")
	q.Add("networkInterfaceList.1.networkInterfaceOrder", "0")
	q.Add("serverImageProductCode", "SW.VSVR.OS.LNX64.UBNTU.SVR1804.B050")
	q.Add("networkInterfaceList.1.accessControlGroupNoList.1", "132668")
	req.URL.RawQuery = q.Encode()
	fmt.Println("url: " + req.URL.String())

	SetHeader(req)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println("Response:", resp)
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

func main() {
	TestCreate()
}
