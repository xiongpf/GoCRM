package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CRMContract struct {
	Id                   int       `orm:"column(id);auto"`
	ContractName         string    `orm:"column(Contract_name);size(250);null"`
	Serialnumber         string    `orm:"column(Serialnumber);size(250);null"`
	CustomerId           int       `orm:"column(Customer_id);null"`
	CustomerName         string    `orm:"column(Customer_name);size(250);null"`
	CDepid               int       `orm:"column(C_depid);null"`
	CDepname             string    `orm:"column(C_depname);size(250);null"`
	CEmpid               int       `orm:"column(C_empid);null"`
	CEmpname             string    `orm:"column(C_empname);size(250);null"`
	ContractAmount       float32   `orm:"column(Contract_amount);null"`
	PayCycle             int       `orm:"column(Pay_cycle);null"`
	StartDate            string    `orm:"column(Start_date);size(250);null"`
	EndDate              string    `orm:"column(End_date);size(250);null"`
	SignDate             string    `orm:"column(Sign_date);size(250);null"`
	CustomerContractor   string    `orm:"column(Customer_Contractor);size(250);null"`
	OurContractorDepid   int       `orm:"column(Our_Contractor_depid);null"`
	OurContractorDepname string    `orm:"column(Our_Contractor_depname);size(250);null"`
	OurContractorId      int       `orm:"column(Our_Contractor_id);null"`
	OurContractorName    string    `orm:"column(Our_Contractor_name);size(250);null"`
	CreaterId            int       `orm:"column(Creater_id);null"`
	CreaterName          string    `orm:"column(Creater_name);size(250);null"`
	CreateTime           time.Time `orm:"column(Create_time);type(datetime);null"`
	MainContent          string    `orm:"column(Main_Content);size(250);null"`
	Remarks              string    `orm:"column(Remarks);size(250);null"`
	FileSerialnumber     string    `orm:"column(File_serialnumber);size(250);null"`
	IsDelete             int       `orm:"column(isDelete);null"`
	DeleteTime           time.Time `orm:"column(Delete_time);type(datetime);null"`
}

func (t *CRMContract) TableName() string {
	return "CRM_contract"
}

func init() {
	orm.RegisterModel(new(CRMContract))
}

// AddCRMContract insert a new CRMContract into database and returns
// last inserted Id on success.
func AddCRMContract(m *CRMContract) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCRMContractById retrieves CRMContract by Id. Returns error if
// Id doesn't exist
func GetCRMContractById(id int) (v *CRMContract, err error) {
	o := orm.NewOrm()
	v = &CRMContract{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRMContract retrieves all CRMContract matches certain condition. Returns empty list if
// no records exist
func GetAllCRMContract(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRMContract))
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

	var l []CRMContract
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

// UpdateCRMContract updates CRMContract by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRMContractById(m *CRMContract) (err error) {
	o := orm.NewOrm()
	v := CRMContract{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRMContract deletes CRMContract by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRMContract(id int) (err error) {
	o := orm.NewOrm()
	v := CRMContract{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRMContract{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
