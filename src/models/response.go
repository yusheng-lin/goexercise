package models

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PagingResponse struct {
	TotalRows  int64       `json:"totalrows"`
	TotalPages int64       `json:"totalpages"`
	PageNo     int         `json:"pageno"`
	Data       interface{} `json:"data"`
	Rows       int         `json:"rows"`
}
