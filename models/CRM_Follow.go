package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMFollow struct {
	Id             int       `orm:"column(id);auto"`
	CustomerId     int       `orm:"column(Customer_id);null"`
	CustomerName   string    `orm:"column(Customer_name);size(250);null"`
	Follow         string    `orm:"column(Follow);size(250);null"`
	FollowDate     time.Time `orm:"column(Follow_date);type(datetime);null"`
	FollowTypeId   int       `orm:"column(Follow_Type_id);null"`
	FollowType     string    `orm:"column(Follow_Type);size(250);null"`
	DepartmentId   int       `orm:"column(department_id);null"`
	DepartmentName string    `orm:"column(department_name);size(250);null"`
	EmployeeId     int       `orm:"column(employee_id);null"`
	EmployeeName   string    `orm:"column(employee_name);size(250);null"`
	IsDelete       int       `orm:"column(isDelete);null"`
	DeleteTime     time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMFollow) TableName() string {
	return "CRM_Follow"
}

func init() {
	orm.RegisterModel(new(CRMFollow))
}

// AddCRMFollow insert a new CRMFollow into database and returns
// last inserted Id on success.
func AddCRMFollow(m *CRMFollow) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMFollowById retrieves CRMFollow by Id. Returns error if
// Id doesn't exist
func GetCRMFollowById(id int) (v *CRMFollow, err error) {
	o := orm.NewOrm()
	v = &CRMFollow{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMFollow retrieves all CRMFollow matches certain condition. Returns empty list if
// no records exist
func GetAllCRMFollow(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMFollow))
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

	var l []CRMFollow
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

// UpdateCRMFollow updates CRMFollow by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMFollowById(m *CRMFollow) (err error) {
	o := orm.NewOrm()
	v := CRMFollow{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMFollow deletes CRMFollow by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMFollow(id int) (err error) {
	o := orm.NewOrm()
	v := CRMFollow{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMFollow{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
