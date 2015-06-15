package models

import "time"

type SysAuthority struct {
	RoleId     int       `orm:"column(Role_id)"`
	AppIds     string    `orm:"column(App_ids);size(250);null"`
	MenuIds    string    `orm:"column(Menu_ids);size(250);null"`
	ButtonIds  string    `orm:"column(Button_ids);size(250);null"`
	CreateId   int       `orm:"column(Create_id);null"`
	CreateDate time.Time `orm:"column(Create_date);type(datetime);null"`
}
