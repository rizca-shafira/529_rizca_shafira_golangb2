package models

import "github.com/jinzhu/gorm"

type Reload struct {
	gorm.Model
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
