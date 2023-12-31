package types

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	// types "github.com/cloud-club/Aviator-service/types/server"
)

// 필수가 아닌 필드(필수 여부: No)는 주석 처리 해두었음.
// 필요할 때 주석 해제
// 필수가 아닌 필드 중 (필수 여부: Conditional)는 주석 처리 안 했음.
// 필요할 때 주석 처리

type RequestInterface interface {
	RequestString() string
	MapResponse(responseBody []byte) (interface{}, error)
}
type CreateServerRequest struct {
	//RegionCode                        string               `json:"regionCode"`
	// ServerImageProductCode 와 MemberServerImageInstanceNo 둘 중 하나는 무조건 필수 기재
	MemberServerImageInstanceNo string `json:"memberServerImageInstanceNo"` // Conditional
	ServerImageProductCode      string `json:"serverImageProductCode"`      // Conditional
	ServerImageNo               string `json:"serverImageNo"`               // Conditional
	VpcNo                       string `json:"vpcNo"`
	SubnetNo                    string `json:"subnetNo"`
	//ServerProductCode                 string               `json:"serverProductCode"`
	ServerSpecCode string `json:"serverSpecCode"` // Conditional
	//IsEncryptedBaseBlockStorageVolume bool                 `json:"isEncryptedBaseBlockStorageVolume"`
	//FeeSystemTypeCode                 string               `json:"feeSystemTypeCode"`
	//ServerCreateCount                 int                  `json:"serverCreateCount"`
	//ServerCreateStartNo               int                  `json:"serverCreateStartNo"`
	//ServerName                        string               `json:"serverName"`

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// NetworkInterfaceList.N의 값들
	// NetworkInterfaceList를 먼저 호출한 뒤, N 번째 NetworkInterface 정보에서 필요한 부분들 추출해서 사용

	//기본 네트워크 인터페이스로 설정하려면 0 입력
	// min:0, max:2
	NetworkInterfaceOrder int `json:"networkInterfaceList.N.networkInterfaceOrder"`

	// 사용자가 직접 생성한 네트워크 인터페이스를 추가하려고 하는 경우 해당 네트워크 인터페이스 번호 입력
	NetworkInterfaceNo string `json:"networkInterfaceNo"`

	// 새로 생성할 네트워크 인터페이스의 서브넷 또는 추가하려고 하는 기존 네트워크 인터페이스의 서브넷 결정
	// 기본 네트워크 인터페이스인 경우에는 자동으로 할당
	NetworkInterfaceSubnetNo string `json:"networkInterfaceSubnetNo"`

	//NetworkInterfaceIp 	  string               `json:"networkInterfaceIp"`

	// 네트워크 인터페이스를 새로 생성하는 경우 반드시 적용할 ACG 결정
	// 최대 3개의 ACG 적용 가능
	// accessControlGroupNo는 getAccessControlGroupList 액션을 통해 획득 가능
	AccessControlGroupNoListN []string `json:"accessControlGroupNoList"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	// PlacementGroupNo           string               `json:"placementGroupNo"`
	// IsProtectServerTermination bool                 `json:"isProtectServerTermination"`
	// ServerDescription          string               `json:"serverDescription"`
	// InitScriptNo               string               `json:"initScriptNo"`
	// LoginKeyName               string               `json:"loginKeyName"`
	// AssociateWithPublicIp      bool                 `json:"associateWithPublicIp"`
	RaidTypeName string `json:"raidTypeName"` // Conditional

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// BlockDevicePartitionList.N의 값들
	// BlockDevicePartitionList를 먼저 호출한 뒤, N 번째 BlockDevicePartition 정보에서 필요한 부분들 추출해서 사용
	//BlockDevicePartitionMountPoint string               `json:"blockDevicePartitionMountPoint"`
	//BlockDevicePartitionSize       string               `json:"blockDevicePartitionSize"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// BlockStorageMappingList.N의 값들
	// BlockStorageMappingList를 먼저 호출한 뒤, N 번째 BlockStorageMapping 정보에서 필요한 부분들 추출해서 사용
	//BlockStorageMappingOrder                      int                  `json:"blockStorageMappingList"`
	//BlockStorageMappingSnapshotInstanceNo         string               `json:"blockStorageMappingSnapshotInstanceNo"`
	//BlockStorageMappingBlockStorageSize           string               `json:"blockStorageMappingBlockStorageSize"`
	//BlockStorageMappingBlockStorageName           string               `json:"blockStorageMappingBlockStorageName"`
	//BlockStorageMappingBlockStorageVolumeTypeCode string               `json:"blockStorageMappingBlockStorageVolumeTypeCode"`
	//BlockStorageMappingEncrypted                  string               `json:"blockStorageMappingEncrypted"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	//ResponseFormatType string `json:"responseFormatType"`
}

type GetProductRequest struct {
	ServerImageProductCode string `json:"serverImageProductCode"`
}

func (ssr GetProductRequest) RequestString() string {
	v := url.Values{}
	s := reflect.ValueOf(ssr)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

func (ssr GetProductRequest) MapResponse(responseBody []byte) (interface{}, error) {
	v := &GetProductResponse{}

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

type ListServerRequest struct {
	RegionCode string `json:"regionCode"`
}

func (ssr ListServerRequest) RequestString() string {
	v := url.Values{}
	s := reflect.ValueOf(ssr)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

func processTimestamp(input []byte) (resultReponse []byte) {
	regex := regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})([+-]\d{4})`)

	// Find all occurrences of timestamps in the input
	matches := regex.FindAllSubmatchIndex(input, -1)

	if len(matches) > 0 {
		// Create a copy of the original input
		modifiedInput := make([]byte, len(input)+len(matches))
		copy(modifiedInput, input)

		for i, match := range matches {
			start, end := match[2]-4*i, match[3]-4*i
			appendingTS := []byte{90} // represents 'Z' in ascii
			originalTS := make([]byte, 19)
			copy(originalTS, modifiedInput[start:end][:19])
			timestamp := append(originalTS, appendingTS...)
			modifiedInput = append(modifiedInput[:start], append([]byte(timestamp), modifiedInput[end+5:]...)...)
		}
		return modifiedInput
	} else {
		fmt.Println("Timestamps not found in the input")
	}
	return input
}

func (ssr ListServerRequest) MapResponse(responseBody []byte) (interface{}, error) {
	v := &ListServerResponse{}

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

type StopServerRequest struct {
	ServerNo string `json:"serverInstanceNoList.1"` // limiting only to a single server instance
}

func (ssr StopServerRequest) MapResponse(responseBody []byte) (interface{}, error) {
	v := &StopServerResponse{}

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

func (ssr StopServerRequest) RequestString() string {
	v := url.Values{}
	s := reflect.ValueOf(ssr)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

type DeleteServerRequest struct {
	ServerNo string `json:"serverInstanceNoList.1"` // limiting only to a single server instance
}

func (ssr DeleteServerRequest) MapResponse(responseBody []byte) (interface{}, error) {
	v := &DeleteServerResponse{}

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

func (ssr DeleteServerRequest) RequestString() string {
	v := url.Values{}
	s := reflect.ValueOf(ssr)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

type UpdateServerRequest struct {
	ServerInstanceNo  string `json:"serverInstanceNo"`
	ServerProductCode string `json:"serverProductCode"` //conditional
	// ServerSpecCode    string `json:"serverSpecCode"`    //conditional
}

func (ssr UpdateServerRequest) MapResponse(responseBody []byte) (interface{}, error) {
	v := &UpdateServerResponse{}

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

func (ssr UpdateServerRequest) RequestString() string {
	v := url.Values{}
	s := reflect.ValueOf(ssr)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

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

type ListServerResponse struct {
	ReturnCode         int              `xml:"returnCode"`
	ReturnMessage      string           `xml:"returnMessage"`
	TotalRows          int              `xml:"totalRows"`
	ServerInstanceList []ServerInstance `xml:"serverInstanceList>serverInstance"`
}

type UpdateServerResponse struct {
	RequestId          string           `xml:"requestId"`
	ReturnCode         int              `xml:"returnCode"`
	ReturnMessage      string           `xml:"returnMessage"`
	TotalRows          int              `xml:"totalRows"`
	ServerInstanceList []ServerInstance `xml:"serverInstanceList>serverInstance"`
}

type DeleteServerResponse struct {
	ReturnCode         int              `xml:"returnCode"`
	ReturnMessage      string           `xml:"returnMessage"`
	TotalRows          int              `xml:"totalRows"`
	ServerInstanceList []ServerInstance `xml:"serverInstanceList>serverInstance"`
}

type StopServerResponse struct {
	ReturnCode         int              `xml:"returnCode"`
	ReturnMessage      string           `xml:"returnMessage"`
	TotalRows          int              `xml:"totalRows"`
	ServerInstanceList []ServerInstance `xml:"serverInstanceList>serverInstance"`
}

type ProductInstance struct {
	ProductCode          string     `xml:"productCode" json:"productCode"`
	ProductName          string     `xml:"productName" json:"productName"`
	ProductType          CommonCode `xml:"productType" json:"productType"`
	ProductDescription   string     `xml:"productDescription" json:"productDescription"`
	InfraResourceType    CommonCode `xml:"infraResourceType" json:"infraResourceType"`
	CpuCount             string     `xml:"cpuCount" json:"cpuCount"`
	MemorySize           string     `xml:"memorySize" json:"memorySize"`
	BaseBlockStorageSize string     `xml:"baseBlockStorageSize" json:"baseBlockStorageSize"`
	OsInformation        string     `xml:"osInformation" json:"osInformation"`
	DiskType             CommonCode `xml:"diskType" json:"diskType"`
	DbKindCode           string     `xml:"dbKindCode" json:"dbKindCode"`
	AddBlockStorageSize  string     `xml:"addBlockStorageSize" json:"addBlockStorageSize"`
	GenerationCode       string     `xml:"generationCode" json:"generationCode"`
}
type GetProductResponse struct {
	RequestId     string            `xml:"requestId"`
	ReturnCode    int               `xml:"returnCode"`
	ReturnMessage string            `xml:"returnMessage"`
	TotalRows     int               `xml:"totalRows"`
	ProductList   []ProductInstance `xml:"productList>product"`
}
