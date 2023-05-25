package models

import "goexercise/enum"

type Login struct {
	Acct string `json:"acct" example:"ben"`
	Pwd  string `json:"pwd" example:"pass.123"`
}

type DeleteAccount struct {
	Acct string
}

type Paging struct {
	// enum: 0 asc, 1 desc
	Sort enum.Sort `json:"sort" example:"0"`
	// enum: 0 account, 1 fullname, 2 created
	SortBy enum.UserSort `json:"sortby" example:"0"`
	Take   int           `json:"take" example:"50"`
	PageNo int           `json:"pageno" example:"1"`
}

func (page *Paging) Validate() bool {
	return page.PageNo > 0 && page.Take > 0
}
