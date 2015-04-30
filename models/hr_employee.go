package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HrEmployee struct {
	Id           int       `orm:"column(id);auto"`
	Uid          string    `orm:"column(uid);size(50);null"`
	Pwd          string    `orm:"column(pwd);size(50);null"`
	Name         string    `orm:"column(name);size(50);null"`
	Idcard       string    `orm:"column(idcard);size(50);null"`
	Birthday     string    `orm:"column(birthday);size(50);null"`
	DId          int       `orm:"column(d_id);null"`
	Dname        string    `orm:"column(dname);size(50);null"`
	Postid       int       `orm:"column(postid);null"`
	Post         string    `orm:"column(post);size(250);null"`
	Email        string    `orm:"column(email);size(50);null"`
	Sex          string    `orm:"column(sex);size(50);null"`
	Tel          string    `orm:"column(tel);size(50);null"`
	Status       string    `orm:"column(status);size(50);null"`
	Zhiwuid      int       `orm:"column(zhiwuid);null"`
	Zhiwu        string    `orm:"column(zhiwu);size(50);null"`
	Sort         int       `orm:"column(sort);null"`
	EntryDate    string    `orm:"column(EntryDate);size(50);null"`
	Address      string    `orm:"column(address);size(255);null"`
	Remarks      string    `orm:"column(remarks);size(255);null"`
	Education    string    `orm:"column(education);size(50);null"`
	Level        string    `orm:"column(level);size(50);null"`
	Professional string    `orm:"column(professional);size(50);null"`
	Schools      string    `orm:"column(schools);size(50);null"`
	Title        string    `orm:"column(title);size(50);null"`
	IsDelete     int       `orm:"column(isDelete);null"`
	DeleteTime   time.Time `orm:"column(Delete_time);type(datetime);null"`
	Portal       string    `orm:"column(portal);size(250);null"`
	Theme        string    `orm:"column(theme);size(250);null"`
	Canlogin     int       `orm:"column(canlogin);null"`
}

func (t *HrEmployee) TableName() string {
	return "hr_employee"
}

func init() {
	orm.RegisterModel(new(HrEmployee))
}

// AddHrEmployee insert a new HrEmployee into database and returns
// last inserted Id on success.
func AddHrEmployee(m *HrEmployee) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHrEmployeeById retrieves HrEmployee by Id. Returns error if
// Id doesn't exist
func GetHrEmployeeById(id int) (v *HrEmployee, err error) {
	o := orm.NewOrm()
	v = &HrEmployee{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHrEmployee retrieves all HrEmployee matches certain condition. Returns empty list if
// no records exist
func GetAllHrEmployee(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HrEmployee))
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

	var l []HrEmployee
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

// UpdateHrEmployee updates HrEmployee by Id and returns error if
// the record to be updated doesn't exist
func UpdateHrEmployeeById(m *HrEmployee) (err error) {
	o := orm.NewOrm()
	v := HrEmployee{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHrEmployee deletes HrEmployee by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHrEmployee(id int) (err error) {
	o := orm.NewOrm()
	v := HrEmployee{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HrEmployee{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
