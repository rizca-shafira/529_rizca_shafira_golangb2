package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null; unique; type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
Gorm juga telah menyediakn tags yang dapat kita gunakan,
seperti halnya property ID diberikan tag gorm:”primaryKey”.
Ini digunakan untuk menjadikan kolom id tersebut menjadi
primary key dari table User.Kemudian kita juga dapat
memberikan constraint kepada suatu kolom, seperti halnya
pada property email. Terdapat tag not null, dan unique.
Hal ini dilakukan agar kolom email dari table User menjadi
unique column dan tidak boleh null alias not null.Property
Email dari struct memiliki tipe data string, dan secara
otomatis tipe datanya akan menjadi TEXT pada database.
Namun kita dapat mengganti tipe datanya seperti contohnya
varchar, maka kita dapat melalukannya dengan memberikan tag type:<tipe_data>
*/

/*
cara menginstall driver Postgresql:
go get gorm.io/driver/postgres.
*/
