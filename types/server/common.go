package types

type CommonCode struct {
	Code     string // server status : int | creat | run | nstop
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
