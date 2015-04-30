package models

import "time"

type SysOnline struct {
	UserID      int       `orm:"column(UserID);null"`
	UserName    string    `orm:"column(UserName);size(50);null"`
	LastLogTime time.Time `orm:"column(LastLogTime);type(datetime);null"`
}
