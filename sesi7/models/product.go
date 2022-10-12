/*
Gorm telah menyediakan berbagai macam assosiasi,
seperti one to one, one to many, dan many to many.
Misalkan contohnya selain kita ingin membuat table User,
kita juga ingin membuat table Products. Table Products
ini akan memiliki foreign key dari ID User, sehingga
assosiasinya akan menjadi one to many(Satu User bisa
memiliki banyak Product, dan satu Product hanya
dapat memiliki satu Product).
*/

package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
By default, property UserIDtersebut akan secara otomatis
menjadi foreign key karena, ketika suatu struct memiliki
property yang dimana nama propertynya merupakan gabungan
dari nama struct lain dengan primary key dari struct tersebut,
maka Gorm secara otomatis akan  menjadikan property tersebut
menjadi foreign key dari struct lainnya.Nama structPrimary Key

*/

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")
	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}
	return
}

/*
eager loading untuk join statement
method Preload dan kita perlu memberikan nama table
untuk parameter method Preload. Walaupun nama tablenya
adalah products jika kita melihat pada
database, tapi tetap huruf awal dari nama table
untuk parameter Preload harus menggunakan uppercase.

*/
