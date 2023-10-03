package domain

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	RespInternal = &Response{Code: 500, Message: "Internal server error"}
	RespBadReq   = &Response{Code: 400, Message: "Bad request"}
)
