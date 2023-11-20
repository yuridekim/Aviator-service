package types

import "time"

type CommonCode struct {
	Code     string // instance status : int | creat | run | nstop
	CodeName string
}

// BlockDevicePartition means block disk parition info (such as AWS EBS)
type BlockDevicePartition struct {
	MountPoint    string
	PartitionSize string
}

// NetworkInterfaceNoList is same as NIC number list
type NetworkInterfaceNoList struct {
	NetworkInterfaceNoList []string
}

// ServerInstance is same as compute server instance in Naver Cloud (such as AWS EC2)
type ServerInstance struct {
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

type ServerInstanceList struct {
	TotalCount   int
	InstanceList []ServerInstance
}
