package models

import "time"

type ParamSysParamType struct {
	Id_RENAME   int       `orm:"column(id)"`
	ParamsName  string    `orm:"column(params_name);size(250);null"`
	ParamsOrder int       `orm:"column(params_order);null"`
	IsDelete    int       `orm:"column(isDelete);null"`
	DeleteTime  time.Time `orm:"column(Delete_time);type(datetime);null"`
}
