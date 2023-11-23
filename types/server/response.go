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
