package http

type RespCode int16

const (
	CodeSuccess RespCode = 1111
	CodeFail    RespCode = 1112
)

var CodeMessageMap = map[RespCode]string{
	CodeSuccess: "Success",
	CodeFail:    "Failed",
}

func (code RespCode) Msg() string {
	return CodeMessageMap[code]
}
