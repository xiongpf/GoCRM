package models

import "time"

type CRMContractAttachment struct {
	ContractId int       `orm:"column(contract_id);null"`
	PageId     string    `orm:"column(page_id);size(250);null"`
	FileId     string    `orm:"column(file_id);size(250);null"`
	FileName   string    `orm:"column(file_name);size(250);null"`
	RealName   string    `orm:"column(real_name);size(250);null"`
	FileSize   int       `orm:"column(file_size);null"`
	CreateId   int       `orm:"column(create_id);null"`
	CreateName string    `orm:"column(create_name);size(250);null"`
	CreateDate time.Time `orm:"column(create_date);type(datetime);null"`
}
