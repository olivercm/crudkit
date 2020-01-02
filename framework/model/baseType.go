package model

type BaseType struct {
	Success   bool        `json:"success"`
	Total     int         `json:"total"`
	Data      interface{} `json:"body"`
	Datas     interface{} `json:"bodys"`
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	Current   int         `json:"current"`
	Extra     string      `json:"extra"`
}
