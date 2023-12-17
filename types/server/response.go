package types

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"time"
)

// ServerInstance is same as compute server in Naver Cloud (such as AWS EC2)
// 필수가 아닌 필드(필수 여부: No)는 주석 처리 해두었음.
// 필요할 때 주석 해제
type ServerInstance struct {
	ServerInstanceNo string `xml:"serverInstanceNo" json:"serverInstanceNo"`
	ServerName       string `xml:"serverName" json:"serverName"`
	//ServerDescription              string
	CpuCount     int        `xml:"cpuCount" json:"cpuCount"`
	MemorySize   int64      `xml:"memorySize" json:"memorySize"`
	PlatformType CommonCode `xml:"platformType" json:"platformType"`
	LoginKeyName string     `xml:"loginKeyName" json:"loginKeyName"`
	//PublicIpInstanceNo             string
	//PublicIp                       string
	ServerInstanceStatus       CommonCode             `xml:"serverInstanceStatus" json:"serverInstanceStatus"`
	ServerInstanceOperation    CommonCode             `xml:"serverInstanceOperation" json:"serverInstanceOperation"`
	ServerInstanceStatusName   string                 `xml:"serverInstanceStatusName" json:"serverInstanceStatusName"`
	CreateDate                 time.Time              `xml:"createDate" json:"createDate"`
	Uptime                     time.Time              `xml:"uptime" json:"uptime"`
	ServerImageProductCode     string                 `xml:"serverImageProductCode" json:"serverImageProductCode"`
	ServerProductCode          string                 `xml:"serverProductCode" json:"serverProductCode"`
	IsProtectServerTermination bool                   `xml:"isProtectServerTermination" json:"isProtectServerTermination"`
	ZoneCode                   string                 `xml:"zoneCode" json:"zoneCode"`
	RegionCode                 string                 `xml:"regionCode" json:"regionCode"`
	VpcNo                      string                 `xml:"vpcNo" json:"vpcNo"`
	SubnetNo                   string                 `xml:"subnetNo" json:"subnetNo"`
	NetworkInterfaceNoList     NetworkInterfaceNoList `xml:"networkInterfaceNoList" json:"networkInterfaceNoList"`
	//InitScriptNo                   string
	ServerInstanceType             CommonCode `xml:"serverInstanceType" json:"serverInstanceType"`
	BaseBlockStorageDiskType       CommonCode `xml:"baseBlockStorageDiskType" json:"baseBlockStorageDiskType"`
	BaseBlockStorageDiskDetailType CommonCode `xml:"baseBlockStorageDiskDetailType" json:"baseBlockStorageDiskDetailType"`
	//PlacementGroupNo               string
	//PlacementGroupName             string
	//MemberServerImageInstanceNo    string
	//BlockDevicePartitionList       []BlockDevicePartition // Assuming BlockDevicePartition is a defined struct
	HypervisorType CommonCode `xml:"hypervisorType" json:"hypervisorType"`
	ServerImageNo  string     `xml:"serverImageNo" json:"serverImageNo"`
	ServerSpecCode string     `xml:"serverSpecCode" json:"serverSpecCode"`
}

type CreateServerResponse struct {
	RequestId          string           `xml:"requestId"`
	ReturnCode         int              `xml:"returnCode"`
	ReturnMessage      string           `xml:"returnMessage"`
	TotalRows          int              `xml:"totalRows"`
	ServerInstanceList []ServerInstance `xml:"serverInstanceList>serverInstance"`
}

type GetServerResponse struct{}

type ListServerResponse struct{}

type UpdateServerResponse struct{}

type DeleteServerResponse struct{}

// Response XML을 Go struct로 변환하는 함수
// responseBody: API 서버로부터 받은 XML response body
// v: interface{} 타입의 struct 포인터
// 예시: mapResponse(responseBody, &CreateServerResponse{})
func MapResponse(responseBody []byte, v interface{}) (interface{}, error) {
	rv := reflect.ValueOf(v)                    // reflect를 이용해 v가 어떤 구조체 타입인지 알아냄
	if rv.Kind() != reflect.Ptr || rv.IsNil() { // 만약 v가 포인터가 아니거나 nil이면 에러 반환
		return nil, fmt.Errorf("non-nil pointer expected")
	}

	err := xml.Unmarshal(responseBody, v) // responseBody를 v로 매핑. 만약 CreateServerResponse 타입이면 CreateServerResponse로 매핑
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}
