/*
Sekarang kita akan membuat sebuah endpoint unutk keperluan
create data atau membuat data mobil baru.Untuk itu pada file
carController.go, kita perlu mengimport dulu framework Gin yang
telah kita install dengan cara seperti pada gambar dibawah ini.
Kita juga perlu mengimport package net/http untuk penentuan
status code dan package fmt untuk memformat data string.
*/

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct dengan nama Car yang akan kita jadikan sebagai struktur data dari data-data mobilnya,
// dan kita juga memerlukan sebuah variable global dengan tipe data list yang akan menampung seluruh data mobilnya.
type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	//Function CreateCar akan menjadi endpoint untuk proses create data.
	//Function ini perlu menerima sebuah parameter dengan tipe data *gin.Context
	//yang merupakan sebuah tipe data yang telah disediakan oleh framework
	//Gin.*gin.Context mempunyai berbagai macam method yang dapat
	//kita gunakan untuk mendapat request body dari client, mengirim response dan lain-lain.
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		/*
			ShouldBindJSON merupakan sebuah method dari tipe data *gin.Context yang
			digunakan untuk mem-binding data JSON yang kirimkan oleh client sebagai
			request body kepada server. Method ShouldBindJSON menerima sebuah parameter
			yang dimana kita perlu meletakkan pointer dari variable yang akan
			menampung hasil dari data binding tersebut. Method ShouldBindJSON akan
			mereturn sebuah error jika memang terjadi sebuah error, maka dari itu kita
			perlu memvalidasi terlebih dahulu jika terjadi sebuah error.
		*/
		ctx.AbortWithError(http.StatusBadRequest, err)
		/*
			Kemudian method AbortWithError pada  baris 23 digunakan untuk
			melempar error jika memang ada error yang terjadi. Method AbortWithError
			menerima 2 parameter. Parameter pertama adalah status error nya,
			kemudian parameter kedua adalah data error nya.
		*/
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	//Kemudian jika tidak terjadi error ketika proses data binding,
	//maka data request body yang telah ditampung oleh variable newCar,
	//akan kita tambahkan pada variable global CarDatas
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
	/*
	 mengirim data response kepada client dengan menggunakan method JSON.
	 Method JSON digunakan untuk mengirim response kepada client dengan
	 format data JSON. Method ini menerima 2 parameter. Parameter
	 pertama adalah status response nya, dan parameter kedua adalah
	 data response yang di kirimkan kepada client.
	*/
}

// membuat proses untuk mengupdate data mobil.
func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	//method ctx.Param merupakan sebuah method yang digunakan untuk mendapat request parameter yang dikirimkan oleh client. Method ctx.Param menerima satu parameter dimana kita perlu meletekkan nama parameter yang kita buat pada router nya nanti.
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}
	//melooping data-data mobil yang ditampung oleh variable CarDatas guna untuk mencari terlebih dahulu data mobil yang ingin di update oleh client nantinya. Kita mencari data mobilnya berdasarkan id data mobil yang dikirmkan client melalui request parameter. Jika data mobilnya tidak dapat, maka kita langsung melempar error dengan status Data Not Found(404).Namun jika data mobilnya ada, maka kita langsung mengupdate data mobilnya dan menghentikan proses loopingnya.

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been succesfully updated", carID),
	})
}

// GetCar digunakan untuk endpoint mendapatkan satu data mobil
func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

/*
endpoint terakhir kita yaitu endpoint untuk menghapus satu data mobil
*/
func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been succesfully deleted", carID),
	})
}

//tugas
/*tugas buat endpoint yang menampilkan semua data tanpa id*/
func GetAll(ctx *gin.Context) {
	var carData []Car
	carData = CarDatas

	ctx.JSON(http.StatusOK, gin.H{
		"cars": carData,
	})
}
