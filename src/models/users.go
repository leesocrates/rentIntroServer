package models

type Users struct {
	Openid   string `xorm:"not null pk VARCHAR(32)"`
	Unionid  string `xorm:"VARCHAR(32)"`
	Nickname string `xorm:"VARCHAR(32)"`
	Headimg  string `xorm:"VARCHAR(256)"`
	Province string `xorm:"VARCHAR(32)"`
	City     string `xorm:"VARCHAR(32)"`
}


