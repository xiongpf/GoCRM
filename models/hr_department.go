package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HrDepartment struct {
	Id         int       `orm:"column(id);auto"`
	DName      string    `orm:"column(d_name);size(50);null"`
	Parentid   int       `orm:"column(parentid);null"`
	Parentname string    `orm:"column(parentname);size(50);null"`
	DType      string    `orm:"column(d_type);size(50);null"`
	DIcon      string    `orm:"column(d_icon);size(50);null"`
	DFuzeren   string    `orm:"column(d_fuzeren);size(50);null"`
	DTel       string    `orm:"column(d_tel);size(50);null"`
	DFax       string    `orm:"column(d_fax);size(50);null"`
	DAdd       string    `orm:"column(d_add);size(255);null"`
	DEmail     string    `orm:"column(d_email);size(50);null"`
	DMiaoshu   string    `orm:"column(d_miaoshu);size(255);null"`
	DOrder     int       `orm:"column(d_order);null"`
	IsDelete   int       `orm:"column(isDelete);null"`
	DeleteTime time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *HrDepartment) TableName() string {
	return "hr_department"
}

func init() {
	orm.RegisterModel(new(HrDepartment))
}

// AddHrDepartment insert a new HrDepartment into database and returns
// last inserted Id on success.
func AddHrDepartment(m *HrDepartment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHrDepartmentById retrieves HrDepartment by Id. Returns error if
// Id doesn't exist
func GetHrDepartmentById(id int) (v *HrDepartment, err error) {
	o := orm.NewOrm()
	v = &HrDepartment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHrDepartment retrieves all HrDepartment matches certain condition. Returns empty list if
// no records exist
func GetAllHrDepartment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HrDepartment))
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

	var l []HrDepartment
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

// UpdateHrDepartment updates HrDepartment by Id and returns error if
// the record to be updated doesn't exist
func UpdateHrDepartmentById(m *HrDepartment) (err error) {
	o := orm.NewOrm()
	v := HrDepartment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHrDepartment deletes HrDepartment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHrDepartment(id int) (err error) {
	o := orm.NewOrm()
	v := HrDepartment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HrDepartment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
