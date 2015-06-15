package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMReceive struct {
	Id                   int       `orm:"column(id);auto"`
	CustomerId           int       `orm:"column(Customer_id);null"`
	CustomerName         string    `orm:"column(Customer_name);size(250);null"`
	ReceiveNum           string    `orm:"column(Receive_num);size(250);null"`
	PayTypeId            int       `orm:"column(Pay_type_id);null"`
	PayType              string    `orm:"column(Pay_type);size(250);null"`
	ReceiveAmount        float32   `orm:"column(Receive_amount);null"`
	ReceiveDate          time.Time `orm:"column(Receive_date);type(datetime);null"`
	CDepid               int       `orm:"column(C_depid);null"`
	CDepname             string    `orm:"column(C_depname);size(250);null"`
	CEmpid               int       `orm:"column(C_empid);null"`
	CEmpname             string    `orm:"column(C_empname);size(250);null"`
	CreateId             int       `orm:"column(create_id);null"`
	CreateName           string    `orm:"column(create_name);size(250);null"`
	CreateDate           time.Time `orm:"column(create_date);type(datetime);null"`
	Companyid            int       `orm:"column(companyid);null"`
	OrderId              int       `orm:"column(order_id);null"`
	Remarks              string    `orm:"column(remarks);size(250);null"`
	IsDelete             int       `orm:"column(isDelete);null"`
	DeleteTime           time.Time `orm:"column(Delete_time);type(datetime);null"`
	ReceiveDirectionId   int       `orm:"column(receive_direction_id);null"`
	ReceiveDirectionName string    `orm:"column(receive_direction_name);size(250);null"`
	ReceiveTypeId        int       `orm:"column(receive_type_id);null"`
	ReceiveTypeName      string    `orm:"column(receive_type_name);size(250);null"`
	ReceiveReal          float32   `orm:"column(receive_real);null"`
}

func (t *CRMReceive) TableName() string {
	return "CRM_receive"
}

func init() {
	orm.RegisterModel(new(CRMReceive))
}

// AddCRMReceive insert a new CRMReceive into database and returns
// last inserted Id on success.
func AddCRMReceive(m *CRMReceive) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMReceiveById retrieves CRMReceive by Id. Returns error if
// Id doesn't exist
func GetCRMReceiveById(id int) (v *CRMReceive, err error) {
	o := orm.NewOrm()
	v = &CRMReceive{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMReceive retrieves all CRMReceive matches certain condition. Returns empty list if
// no records exist
func GetAllCRMReceive(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMReceive))
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

	var l []CRMReceive
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

// UpdateCRMReceive updates CRMReceive by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMReceiveById(m *CRMReceive) (err error) {
	o := orm.NewOrm()
	v := CRMReceive{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMReceive deletes CRMReceive by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMReceive(id int) (err error) {
	o := orm.NewOrm()
	v := CRMReceive{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMReceive{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
