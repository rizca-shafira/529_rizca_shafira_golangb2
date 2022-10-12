package database

import (
	"fmt"
	"golang_sesi7/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	//menampung config untuk menghubungkan kepada database
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{}) //method Open digunakan untuk membangun koneksi kepada database. Ketika koneksi telah berhasil terbangun, maka variable dbpada baris 25 akan mengandung referensi dari database dengan tipe data *gorm.DB
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	db.Debug().AutoMigrate(models.User{}, models.Product{})
	//Method Debug pada baris 30 digunakan sebagai debugging atau logger. Kemudian di chaining dengan method AutoMigrate.Method AutoMigrate digunakan untuk memigrasi secara otomatis dari struct-struct yang telah dibuat.

}

func GetDB() *gorm.DB {
	return db
}
