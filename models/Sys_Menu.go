package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SysMenu struct {
	Id          int    `orm:"column(Menu_id);auto"`
	MenuName    string `orm:"column(Menu_name);size(255);null"`
	Parentid    int    `orm:"column(parentid);null"`
	Parentname  string `orm:"column(parentname);size(255);null"`
	AppId       int    `orm:"column(App_id);null"`
	MenuUrl     string `orm:"column(Menu_url);size(255);null"`
	MenuIcon    string `orm:"column(Menu_icon);size(50);null"`
	MenuHandler string `orm:"column(Menu_handler);size(50);null"`
	MenuOrder   int    `orm:"column(Menu_order);null"`
	MenuType    string `orm:"column(Menu_type);size(50);null"`
}

func (t *SysMenu) TableName() string {
	return "Sys_Menu"
}

func init() {
	orm.RegisterModel(new(SysMenu))
}

// AddSysMenu insert a new SysMenu into database and returns
// last inserted Id on success.
func AddSysMenu(m *SysMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysMenuById retrieves SysMenu by Id. Returns error if
// Id doesn't exist
func GetSysMenuById(id int) (v *SysMenu, err error) {
	o := orm.NewOrm()
	v = &SysMenu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysMenu retrieves all SysMenu matches certain condition. Returns empty list if
// no records exist
func GetAllSysMenu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysMenu))
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

	var l []SysMenu
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

// UpdateSysMenu updates SysMenu by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysMenuById(m *SysMenu) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysMenu deletes SysMenu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysMenu(id int) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysMenu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
