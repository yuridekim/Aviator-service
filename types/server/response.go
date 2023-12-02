package types

import "time"

// ServerInstanceResponse is same as compute server in Naver Cloud (such as AWS EC2)
type ServerInstanceResponse struct {
	ServerInstanceNo               string
	ServerName                     string
	ServerDescription              string
	CpuCount                       int
	MemorySize                     int64
	PlatformType                   CommonCode
	LoginKeyName                   string
	PublicIpInstanceNo             string
	PublicIp                       string
	ServerInstanceStatus           CommonCode
	ServerInstanceOperation        CommonCode
	ServerInstanceStatusName       string
	CreateDate                     time.Time
	Uptime                         time.Time
	ServerImageProductCode         string
	ServerProductCode              string
	IsProtectServerTermination     bool
	ZoneCode                       string
	RegionCode                     string
	VpcNo                          string
	SubnetNo                       string
	NetworkInterfaceNoList         NetworkInterfaceNoList
	InitScriptNo                   string
	ServerInstanceType             CommonCode
	BaseBlockStorageDiskType       CommonCode
	BaseBlockStorageDiskDetailType CommonCode
	PlacementGroupNo               string
	PlacementGroupName             string
	MemberServerImageInstanceNo    string
	BlockDevicePartitionList       []BlockDevicePartition
	HypervisorType                 CommonCode
	ServerImageNo                  string
	ServerSpecCode                 string
}

type ServerInstanceListResponse struct {
	TotalCount   int
	InstanceList []ServerInstanceResponse
}

type ServerInstance struct {
	ServerInstanceNo               string     `xml:"serverInstanceNo"`
	ServerName                     string     `xml:"serverName"`
	CpuCount                       int        `xml:"cpuCount"`
	MemorySize                     int64      `xml:"memorySize"`
	PlatformType                   CommonCode `xml:"platformType"`
	LoginKeyName                   string
	PublicIpInstanceNo             string
	PublicIp                       string
	ServerInstanceStatus           CommonCode             `xml:"serverInstanceStatus"`
	ServerInstanceOperation        CommonCode             `xml:"serverInstanceOperation"`
	ServerInstanceStatusName       string                 `xml:"serverInstanceStatusName"`
	CreateDate                     time.Time              `xml:"createDate"`
	Uptime                         time.Time              `xml:"uptime"`
	ServerImageProductCode         string                 `xml:"serverImageProductCode"`
	ServerProductCode              string                 `xml:"serverProductCode"`
	IsProtectServerTermination     bool                   `xml:"isProtectServerTermination"`
	ZoneCode                       string                 `xml:"zoneCode"`
	RegionCode                     string                 `xml:"regionCode"`
	VpcNo                          string                 `xml:"vpcNo"`
	SubnetNo                       string                 `xml:"subnetNo"`
	NetworkInterfaceNoList         NetworkInterfaceNoList `xml:"networkInterfaceNoList>networkInterfaceNo"`
	InitScriptNo                   string                 `xml:"initScriptNo"`
	ServerInstanceType             CommonCode             `xml:"serverInstanceType"`
	BaseBlockStorageDiskType       CommonCode             `xml:"baseBlockStorageDiskType"`
	BaseBlockStorageDiskDetailType CommonCode             `xml:"baseBlockStorageDiskDetailType"`
	PlacementGroupNo               string                 `xml:"placementGroupNo"`
	PlacementGroupName             string                 `xml:"placementGroupName"`
	BlockDevicePartitionList       []BlockDevicePartition `xml:"blockDevicePartitionList>blockDevicePartition"`
	HypervisorType                 CommonCode             `xml:"hypervisorType"`
	ServerImageNo                  string                 `xml:"serverImageNo"`
	ServerSpecCode                 string                 `xml:"serverSpecCode"`
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
