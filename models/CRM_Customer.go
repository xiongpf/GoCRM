package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMCustomer struct {
	Id               int       `orm:"column(id);auto"`
	Serialnumber     string    `orm:"column(Serialnumber);size(250);null"`
	Customer         string    `orm:"column(Customer);size(250);null"`
	Address          string    `orm:"column(address);size(250);null"`
	Tel              string    `orm:"column(tel);size(250);null"`
	Fax              string    `orm:"column(fax);size(250);null"`
	Site             string    `orm:"column(site);size(250);null"`
	IndustryId       int       `orm:"column(industry_id);null"`
	Industry         string    `orm:"column(industry);size(250);null"`
	ProvincesId      int       `orm:"column(Provinces_id);null"`
	Provinces        string    `orm:"column(Provinces);size(250);null"`
	CityId           int       `orm:"column(City_id);null"`
	City             string    `orm:"column(City);size(250);null"`
	CustomerTypeId   int       `orm:"column(CustomerType_id);null"`
	CustomerType     string    `orm:"column(CustomerType);size(250);null"`
	CustomerLevelId  int       `orm:"column(CustomerLevel_id);null"`
	CustomerLevel    string    `orm:"column(CustomerLevel);size(250);null"`
	CustomerSourceId int       `orm:"column(CustomerSource_id);null"`
	CustomerSource   string    `orm:"column(CustomerSource);size(250);null"`
	DesCripe         string    `orm:"column(DesCripe);size(4000);null"`
	Remarks          string    `orm:"column(Remarks);size(4000);null"`
	DepartmentId     int       `orm:"column(Department_id);null"`
	Department       string    `orm:"column(Department);size(250);null"`
	EmployeeId       int       `orm:"column(Employee_id);null"`
	Employee         string    `orm:"column(Employee);size(250);null"`
	Privatecustomer  string    `orm:"column(privatecustomer);size(50);null"`
	Lastfollow       time.Time `orm:"column(lastfollow);type(datetime);null"`
	CreateId         int       `orm:"column(Create_id);null"`
	CreateName       string    `orm:"column(Create_name);size(250);null"`
	CreateDate       time.Time `orm:"column(Create_date);type(datetime);null"`
	IsDelete         int       `orm:"column(isDelete);null"`
	DeleteTime       time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMCustomer) TableName() string {
	return "CRM_Customer"
}

func init() {
	orm.RegisterModel(new(CRMCustomer))
}

// AddCRMCustomer insert a new CRMCustomer into database and returns
// last inserted Id on success.
func AddCRMCustomer(m *CRMCustomer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMCustomerById retrieves CRMCustomer by Id. Returns error if
// Id doesn't exist
func GetCRMCustomerById(id int) (v *CRMCustomer, err error) {
	o := orm.NewOrm()
	v = &CRMCustomer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMCustomer retrieves all CRMCustomer matches certain condition. Returns empty list if
// no records exist
func GetAllCRMCustomer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMCustomer))
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

	var l []CRMCustomer
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

// UpdateCRMCustomer updates CRMCustomer by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMCustomerById(m *CRMCustomer) (err error) {
	o := orm.NewOrm()
	v := CRMCustomer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMCustomer deletes CRMCustomer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMCustomer(id int) (err error) {
	o := orm.NewOrm()
	v := CRMCustomer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMCustomer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
