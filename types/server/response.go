package types

import (
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
