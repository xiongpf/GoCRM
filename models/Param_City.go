package models

import "time"

type ParamCity struct {
	Id_RENAME  int       `orm:"column(id)"`
	Parentid   int       `orm:"column(parentid);null"`
	City       string    `orm:"column(City);size(250);null"`
	CityOrder  int       `orm:"column(City_order);null"`
	CreateId   int       `orm:"column(Create_id);null"`
	CreateDate time.Time `orm:"column(Create_date);type(datetime);null"`
	UpdateId   int       `orm:"column(Update_id);null"`
	UpdateDate time.Time `orm:"column(Update_date);type(datetime);null"`
}
