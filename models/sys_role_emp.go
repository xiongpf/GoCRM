package models

type SysRoleEmp struct {
	RoleID int `orm:"column(RoleID)"`
	EmpID  int `orm:"column(empID)"`
}
