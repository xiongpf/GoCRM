package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysLogErr struct {
	Id         int       `orm:"column(id);auto"`
	ErrTypeid  int       `orm:"column(Err_typeid);null"`
	ErrType    string    `orm:"column(Err_type);size(250);null"`
	ErrTime    time.Time `orm:"column(Err_time);type(datetime);null"`
	ErrUrl     string    `orm:"column(Err_url);size(500);null"`
	ErrMessage string    `orm:"column(Err_message);size(250);null"`
	ErrSource  string    `orm:"column(Err_source);size(500);null"`
	ErrTrace   string    `orm:"column(Err_trace);size(250);null"`
	ErrEmpId   int       `orm:"column(Err_emp_id);null"`
	ErrEmpName string    `orm:"column(Err_emp_name);size(250);null"`
	ErrIp      string    `orm:"column(Err_ip);size(250);null"`
}

func (t *SysLogErr) TableName() string {
	return "Sys_log_Err"
}

func init() {
	orm.RegisterModel(new(SysLogErr))
}

// AddSysLogErr insert a new SysLogErr into database and returns
// last inserted Id on success.
func AddSysLogErr(m *SysLogErr) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysLogErrById retrieves SysLogErr by Id. Returns error if
// Id doesn't exist
func GetSysLogErrById(id int) (v *SysLogErr, err error) {
	o := orm.NewOrm()
	v = &SysLogErr{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysLogErr retrieves all SysLogErr matches certain condition. Returns empty list if
// no records exist
func GetAllSysLogErr(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysLogErr))
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

	var l []SysLogErr
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

// UpdateSysLogErr updates SysLogErr by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysLogErrById(m *SysLogErr) (err error) {
	o := orm.NewOrm()
	v := SysLogErr{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysLogErr deletes SysLogErr by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysLogErr(id int) (err error) {
	o := orm.NewOrm()
	v := SysLogErr{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysLogErr{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
