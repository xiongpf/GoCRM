package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PersonalCalendar struct {
	Id            int       `orm:"column(Id);auto"`
	EmpId         int       `orm:"column(emp_id);null"`
	EmpName       string    `orm:"column(emp_name);size(250);null"`
	Companyid     int       `orm:"column(companyid);null"`
	Subject       string    `orm:"column(Subject);size(250);null"`
	Location      string    `orm:"column(Location);size(250);null"`
	MasterId      int       `orm:"column(MasterId);null"`
	Description   string    `orm:"column(Description);size(250);null"`
	CalendarType  int8      `orm:"column(CalendarType);null"`
	StartTime     time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime       time.Time `orm:"column(EndTime);type(datetime);null"`
	IsAllDayEvent uint64    `orm:"column(IsAllDayEvent);size(1);null"`
	HasAttachment uint64    `orm:"column(HasAttachment);size(1);null"`
	Category      string    `orm:"column(Category);size(250);null"`
	InstanceType  int8      `orm:"column(InstanceType);null"`
	Attendees     string    `orm:"column(Attendees);size(250);null"`
	AttendeeNames string    `orm:"column(AttendeeNames);size(250);null"`
	OtherAttendee string    `orm:"column(OtherAttendee);size(250);null"`
	UPAccount     string    `orm:"column(UPAccount);size(250);null"`
	UPName        string    `orm:"column(UPName);size(250);null"`
	UPTime        time.Time `orm:"column(UPTime);type(datetime);null"`
	RecurringRule string    `orm:"column(RecurringRule);size(250);null"`
}

func (t *PersonalCalendar) TableName() string {
	return "Personal_Calendar"
}

func init() {
	orm.RegisterModel(new(PersonalCalendar))
}

// AddPersonalCalendar insert a new PersonalCalendar into database and returns
// last inserted Id on success.
func AddPersonalCalendar(m *PersonalCalendar) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPersonalCalendarById retrieves PersonalCalendar by Id. Returns error if
// Id doesn't exist
func GetPersonalCalendarById(id int) (v *PersonalCalendar, err error) {
	o := orm.NewOrm()
	v = &PersonalCalendar{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPersonalCalendar retrieves all PersonalCalendar matches certain condition. Returns empty list if
// no records exist
func GetAllPersonalCalendar(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PersonalCalendar))
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

	var l []PersonalCalendar
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

// UpdatePersonalCalendar updates PersonalCalendar by Id and returns error if
// the record to be updated doesn't exist
func UpdatePersonalCalendarById(m *PersonalCalendar) (err error) {
	o := orm.NewOrm()
	v := PersonalCalendar{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePersonalCalendar deletes PersonalCalendar by Id and returns error if
// the record to be deleted doesn't exist
func DeletePersonalCalendar(id int) (err error) {
	o := orm.NewOrm()
	v := PersonalCalendar{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PersonalCalendar{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
