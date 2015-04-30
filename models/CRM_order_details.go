package models

type CRMOrderDetails struct {
	OrderId     int     `orm:"column(order_id);null"`
	ProductId   int     `orm:"column(product_id);null"`
	ProductName string  `orm:"column(product_name);size(250);null"`
	Price       float32 `orm:"column(price);null"`
	Quantity    int     `orm:"column(quantity);null"`
	Unit        string  `orm:"column(unit);size(250);null"`
	Amount      float32 `orm:"column(amount);null"`
}
