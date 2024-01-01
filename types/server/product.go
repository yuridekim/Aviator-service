package types

type Product struct {
	ProductCode             string     `xml:"productCode"`
	ProductName             string     `xml:"productName"`
	ProductType             CommonCode `xml:"productType"`
	ProductDescription      string     `xml:"productDescription"`
	InfraResourceType       CommonCode `xml:"infraResourceType"`
	InfraResourceDetailType CommonCode `xml:"infraResourceDetailType"`
	CpuCount                int        `xml:"cpuCount"`
	MemorySize              int64      `xml:"memorySize"`
	BaseBlockStorageSize    int64      `xml:"baseBlockStorageSize"`
	PlatformType            CommonCode `xml:"platformType"`
	OsInformation           string     `xml:"osInformation"`
	DiskType                CommonCode `xml:"diskType"`
	DbKindCode              string     `xml:"dbKindCode"`
	AddBlockStorageSize     int64      `xml:"addBlockStorageSize"`
	GenerationCode          string     `xml:"generationCode"`
}

type ProductList struct {
	TotalRows   int       `xml:"totalRows"`
	ProductList []Product `xml:"productList>product"`
}
