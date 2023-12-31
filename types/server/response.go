package types

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
