package types

type CommonCode struct {
	Code     string `xml:"code"` // server status : int | creat | run | nstop
	CodeName string `xml:"codeName"`
}
