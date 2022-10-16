package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Order_ID      int       `json:"order_id" gorm:"primaryKey;autoIncrement;"`
	Customer_Name string    `json:"customer_name" gorm:"not null;type:varchar(50)"`
	Ordered_At    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
}

type Item struct {
	gorm.Model
	Item_ID     int    `json:"item_id" gorm:"primaryKey;autoIncrement;"`
	Item_Code   string `json:"item_code" gorm:"not null;type:varchar(10)"`
	Description string `json:"description" gorm:"not null;type:varchar(200)"`
	Quantity    int    `json:"quantity" gorm:"not null;"`
	Order_ID    int    `json:"Order_Id" gorm:"index"`
	Items       []Item `json:"item" gorm:"foreignKey:Order_Id"`
}
