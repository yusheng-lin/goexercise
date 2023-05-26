package models

import "time"

type User struct {
	Acct       string
	Fullname   string
	Updated_At time.Time
	Created_At time.Time
}

type Account struct {
	Acct     string
	Fullname string
	Pwd      string
}
