package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMInvoice struct {
	Id             int       `orm:"column(id);auto"`
	CustomerId     int       `orm:"column(Customer_id);null"`
	CustomerName   string    `orm:"column(Customer_name);size(250);null"`
	InvoiceNum     string    `orm:"column(invoice_num);size(250);null"`
	InvoiceTypeId  int       `orm:"column(invoice_type_id);null"`
	InvoiceType    string    `orm:"column(invoice_type);size(250);null"`
	InvoiceAmount  float32   `orm:"column(invoice_amount);null"`
	InvoiceContent string    `orm:"column(invoice_content);size(250);null"`
	InvoiceDate    time.Time `orm:"column(invoice_date);type(datetime);null"`
	CDepid         int       `orm:"column(C_depid);null"`
	CDepname       string    `orm:"column(C_depname);size(250);null"`
	CEmpid         int       `orm:"column(C_empid);null"`
	CEmpname       string    `orm:"column(C_empname);size(250);null"`
	CreateId       int       `orm:"column(create_id);null"`
	CreateName     string    `orm:"column(create_name);size(250);null"`
	CreateDate     time.Time `orm:"column(create_date);type(datetime);null"`
	OrderId        int       `orm:"column(order_id);null"`
	IsDelete       int       `orm:"column(isDelete);null"`
	DeleteTime     time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMInvoice) TableName() string {
	return "CRM_invoice"
}

func init() {
	orm.RegisterModel(new(CRMInvoice))
}

// AddCRMInvoice insert a new CRMInvoice into database and returns
// last inserted Id on success.
func AddCRMInvoice(m *CRMInvoice) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMInvoiceById retrieves CRMInvoice by Id. Returns error if
// Id doesn't exist
func GetCRMInvoiceById(id int) (v *CRMInvoice, err error) {
	o := orm.NewOrm()
	v = &CRMInvoice{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMInvoice retrieves all CRMInvoice matches certain condition. Returns empty list if
// no records exist
func GetAllCRMInvoice(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMInvoice))
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

	var l []CRMInvoice
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

// UpdateCRMInvoice updates CRMInvoice by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMInvoiceById(m *CRMInvoice) (err error) {
	o := orm.NewOrm()
	v := CRMInvoice{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMInvoice deletes CRMInvoice by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMInvoice(id int) (err error) {
	o := orm.NewOrm()
	v := CRMInvoice{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMInvoice{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
