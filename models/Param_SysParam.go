package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ParamSysParam struct {
	Id          int       `orm:"column(id);auto"`
	Parentid    int       `orm:"column(parentid);null"`
	ParamsName  string    `orm:"column(params_name);size(250);null"`
	ParamsOrder int       `orm:"column(params_order);null"`
	CreateId    int       `orm:"column(Create_id);null"`
	CreateDate  time.Time `orm:"column(Create_date);type(datetime);null"`
	UpdateId    int       `orm:"column(Update_id);null"`
	UpdateDate  time.Time `orm:"column(Update_date);type(datetime);null"`
}

func (t *ParamSysParam) TableName() string {
	return "Param_SysParam"
}

func init() {
	orm.RegisterModel(new(ParamSysParam))
}

// AddParamSysParam insert a new ParamSysParam into database and returns
// last inserted Id on success.
func AddParamSysParam(m *ParamSysParam) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetParamSysParamById retrieves ParamSysParam by Id. Returns error if
// Id doesn't exist
func GetParamSysParamById(id int) (v *ParamSysParam, err error) {
	o := orm.NewOrm()
	v = &ParamSysParam{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllParamSysParam retrieves all ParamSysParam matches certain condition. Returns empty list if
// no records exist
func GetAllParamSysParam(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParamSysParam))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ParamSysParam
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateParamSysParam updates ParamSysParam by Id and returns error if
// the record to be updated doesn't exist
func UpdateParamSysParamById(m *ParamSysParam) (err error) {
	o := orm.NewOrm()
	v := ParamSysParam{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteParamSysParam deletes ParamSysParam by Id and returns error if
// the record to be deleted doesn't exist
func DeleteParamSysParam(id int) (err error) {
	o := orm.NewOrm()
	v := ParamSysParam{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ParamSysParam{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
