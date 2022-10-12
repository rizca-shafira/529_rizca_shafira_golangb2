// Untuk menginstall Gorm, maka kita perlu menjalankan perintah: go get -u gorm.io/gorm.
package main

import (
	"errors"
	"fmt"
	"golang_sesi7/database"
	"golang_sesi7/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	//createUser("johndoe@gmail.com")
	//getUserById(1)
	//updateUserById(1, "johnjohn@gmail.com")
	//createProduct(1, "YLO", "YYYY")
	//getUserWithProducts()
	deleteProductById(1)
}

func createUser(email string) {
	//function bernama createUser yang digunakan untuk membuat data User baru
	db := database.GetDB()

	User := models.User{
		Email: email,
	}
	err := db.Create(&User).Error
	/*
		menggunakan method Create. Kemudian method Create ini memerlukan
		data sebagai argumentnya yang dimana data tersebut perlu memiliki
		tipe data yang sama dengan yang ingin kita buat.Misalkan seperti
		pada gambar disebelah kanan, penulis berhendak untuk membuat data
		User baru, alhasil data yang diberikan pada method Create tersebut
		adalah data dengan tipe data struct User dari folder models yang
		telah penulis buat sebelumnya dan juga sudah kita bahas.Method
		untuk CRUD pada Gorm dengan property Error agar kita bisa langsung
		dapat memeriksa errornya jika memang ada.
	*/

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}
	fmt.Println("New User Data:", User)
}

func getUserById(id uint) {
	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, "id = ?", id).Error
	/*
		Untuk mendapatkan suatu data dari suatu table, maka kita dapat menggunakan
		method First.Method First dapat menerima 3 parameter, parameter pertamanya
		adalah pointer terhadap data yang ingin dicari. Karena sekarang penulis
		berhendak untuk mencari data User, maka penulis memberikan tipe data
		struct User sebagai argumen pertama dari method First.Kemudian parameter
		keduanya adalah condition dari query nya, dan yang terakhiir adalah data
		dari conditionnya.Method First akan mengembalikan error berupa ErrRecordNotFound
		jika tidak ada data yang ditemukan.
	*/
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data:%+v\n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email:%+v\n", user.Email)
}

/*
Untuk mengupdate data menggunakan Gorm, digunakan method Model agar hasil dari update nya dapat
langsung discan sehingga kita dapat langsung mengetahui hasilnya.
Kemudian jika kita ingin membuat condition nya, maka kita dapat menggunakan
method Where sehingga method Model dapat dichaining dengan method Where.
Dan yang terkakhir kita bisa langsung menentukan data apa yang ingin di update.
Selain menggunakan method Model, kita juga dapat menggunakan method Table
seperti yang telah dijelaskan pada dokumentasi dari Gorm
*/

/*
gorm telah menyediakan hooks yang dapat kita gunakan, salah satunya method BeforeCreate
before create terkeksekusi sebelum melakukan create data.
*/

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()
	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}
	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New Product Data:", Product)
}

// eager load
func getUserWithProducts() {
	db := database.GetDB()
	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

// untuk menghapus, bisa menggunakan method delete
func deleteProductById(id uint) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}
	fmt.Printf("Product with id %d has been succesfully deleted", id)
}
