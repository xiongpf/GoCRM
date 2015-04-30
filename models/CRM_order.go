package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMOrder struct {
	Id             int       `orm:"column(id);auto"`
	Serialnumber   string    `orm:"column(Serialnumber);size(250);null"`
	CustomerId     int       `orm:"column(Customer_id);null"`
	CustomerName   string    `orm:"column(Customer_name);size(250);null"`
	OrderDate      time.Time `orm:"column(Order_date);type(datetime);null"`
	PayTypeId      int       `orm:"column(pay_type_id);null"`
	PayType        string    `orm:"column(pay_type);size(250);null"`
	OrderDetails   string    `orm:"column(Order_details);size(250);null"`
	OrderStatusId  int       `orm:"column(Order_status_id);null"`
	OrderStatus    string    `orm:"column(Order_status);size(250);null"`
	OrderAmount    float32   `orm:"column(Order_amount);null"`
	CreateId       int       `orm:"column(create_id);null"`
	CreateDate     time.Time `orm:"column(create_date);type(datetime);null"`
	CDepId         int       `orm:"column(C_dep_id);null"`
	CDepName       string    `orm:"column(C_dep_name);size(250);null"`
	CEmpId         int       `orm:"column(C_emp_id);null"`
	CEmpName       string    `orm:"column(C_emp_name);size(250);null"`
	FDepId         int       `orm:"column(F_dep_id);null"`
	FDepName       string    `orm:"column(F_dep_name);size(250);null"`
	FEmpId         int       `orm:"column(F_emp_id);null"`
	FEmpName       string    `orm:"column(F_emp_name);size(250);null"`
	ReceiveMoney   float32   `orm:"column(receive_money);null"`
	ArrearsMoney   float32   `orm:"column(arrears_money);null"`
	InvoiceMoney   float32   `orm:"column(invoice_money);null"`
	ArrearsInvoice float32   `orm:"column(arrears_invoice);null"`
	IsDelete       int       `orm:"column(isDelete);null"`
	DeleteTime     time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMOrder) TableName() string {
	return "CRM_order"
}

func init() {
	orm.RegisterModel(new(CRMOrder))
}

// AddCRMOrder insert a new CRMOrder into database and returns
// last inserted Id on success.
func AddCRMOrder(m *CRMOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMOrderById retrieves CRMOrder by Id. Returns error if
// Id doesn't exist
func GetCRMOrderById(id int) (v *CRMOrder, err error) {
	o := orm.NewOrm()
	v = &CRMOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMOrder retrieves all CRMOrder matches certain condition. Returns empty list if
// no records exist
func GetAllCRMOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMOrder))
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

	var l []CRMOrder
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

// UpdateCRMOrder updates CRMOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMOrderById(m *CRMOrder) (err error) {
	o := orm.NewOrm()
	v := CRMOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMOrder deletes CRMOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMOrder(id int) (err error) {
	o := orm.NewOrm()
	v := CRMOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
