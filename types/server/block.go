package types

// BlockDevicePartition means block disk parition info (such as AWS EBS)
type BlockDevicePartition struct {
	MountPoint    string
	PartitionSize string
}

type BlockDevicePartitionList struct {
	BlockDevicePartitionList []BlockDevicePartition
}

type BlockStorageMapping struct {
	Order                            int
	BlockStorageSnapshotInstanceNo   string
	BlockStorageSnapshotInstanceName string
	BlockStorageSize                 int64
	BlockStorageName                 string
	BlockStorageVolumeType           CommonCode
	Iops                             int64
	Throughput                       int64
	IsEncryptedVolume                bool
}

type BlockStorageMappingList struct {
	BlockStorageMappingList []BlockStorageMapping
}
