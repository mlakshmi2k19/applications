package model

import (
	"time"
)

type Data struct {
	Ord  Order
	Itm  Item
	Prod Product
	Cust Customer
	Reg  Region
}

type Order struct {
	ID            string    `gorm:"primaryKey"`
	CustomerID    string    `gorm:"index"`
	RegionName    string    `gorm:"index"`
	DateOfSale    time.Time `gorm:"type:date"`
	PaymentMethod string
	ShippingCost  float64

	Customer Customer `gorm:"foreignKey:CustomerID"`
	Region   Region   `gorm:"foreignKey:RegionName"`
	Items    []Item   `gorm:"foreignKey:OrderID"`
}

type Item struct {
	ID           uint   `gorm:"primaryKey"`
	OrderID      string `gorm:"index;uniqueIndex:idx_order_product"`
	ProductID    string `gorm:"index;uniqueIndex:idx_order_product"`
	QuantitySold int
	Discount     float64

	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

type Product struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Category  string
	UnitPrice float64
}

type Customer struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Email   string `gorm:"unique"`
	Address string

	Orders []Order `gorm:"foreignKey:CustomerID"`
}

type Region struct {
	Name string `gorm:"primaryKey"`

	Orders []Order `gorm:"foreignKey:RegionName"`
}
