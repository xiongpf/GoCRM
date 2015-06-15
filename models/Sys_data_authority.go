package models

import "time"

type SysDataAuthority struct {
	RoleId     int       `orm:"column(Role_id);null"`
	OptionId   int       `orm:"column(option_id);null"`
	SysOption  string    `orm:"column(Sys_option);size(250);null"`
	SysView    int       `orm:"column(Sys_view);null"`
	SysAdd     int       `orm:"column(Sys_add);null"`
	SysEdit    int       `orm:"column(Sys_edit);null"`
	SysDel     int       `orm:"column(Sys_del);null"`
	CreateId   int       `orm:"column(Create_id);null"`
	CreateDate time.Time `orm:"column(Create_date);type(datetime);null"`
}
