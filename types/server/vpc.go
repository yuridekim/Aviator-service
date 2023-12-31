package types

import "time"

type Vpc struct {
	VpcNo         string     `xml:"vpcNo"`
	VpcName       string     `xml:"vpcName"`
	Ipv4CidrBlock string     `xml:"ipv4CidrBlock"`
	VpcStatus     CommonCode `xml:"vpcStatus"`
	RegionCode    string     `xml:"regionCode"`
	CreateDate    time.Time  `xml:"createDate"`
}

type VpcList struct {
	TotalRows int   `xml:"totalRows"`
	VpcList   []Vpc `xml:"vpcList>vpc"`
}
