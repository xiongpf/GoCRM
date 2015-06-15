package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMContact struct {
	Id            int       `orm:"column(id);auto"`
	CName         string    `orm:"column(C_name);size(250);null"`
	CSex          string    `orm:"column(C_sex);size(10);null"`
	CDepartment   string    `orm:"column(C_department);size(250);null"`
	CPosition     string    `orm:"column(C_position);size(250);null"`
	CBirthday     string    `orm:"column(C_birthday);size(250);null"`
	CTel          string    `orm:"column(C_tel);size(250);null"`
	CFax          string    `orm:"column(C_fax);size(250);null"`
	CEmail        string    `orm:"column(C_email);size(250);null"`
	CMob          string    `orm:"column(C_mob);size(250);null"`
	CQQ           string    `orm:"column(C_QQ);size(250);null"`
	CAdd          string    `orm:"column(C_add);size(250);null"`
	CHobby        string    `orm:"column(C_hobby);size(250);null"`
	CRemarks      string    `orm:"column(C_remarks);size(250);null"`
	CCustomerid   int       `orm:"column(C_customerid);null"`
	CCustomername string    `orm:"column(C_customername);size(250);null"`
	CCreateId     int       `orm:"column(C_createId);null"`
	CCreateDate   time.Time `orm:"column(C_createDate);type(datetime);null"`
	IsDelete      int       `orm:"column(isDelete);null"`
	DeleteTime    time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMContact) TableName() string {
	return "CRM_Contact"
}

func init() {
	orm.RegisterModel(new(CRMContact))
}

// AddCRMContact insert a new CRMContact into database and returns
// last inserted Id on success.
func AddCRMContact(m *CRMContact) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMContactById retrieves CRMContact by Id. Returns error if
// Id doesn't exist
func GetCRMContactById(id int) (v *CRMContact, err error) {
	o := orm.NewOrm()
	v = &CRMContact{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMContact retrieves all CRMContact matches certain condition. Returns empty list if
// no records exist
func GetAllCRMContact(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMContact))
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

	var l []CRMContact
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

// UpdateCRMContact updates CRMContact by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMContactById(m *CRMContact) (err error) {
	o := orm.NewOrm()
	v := CRMContact{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMContact deletes CRMContact by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMContact(id int) (err error) {
	o := orm.NewOrm()
	v := CRMContact{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMContact{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
