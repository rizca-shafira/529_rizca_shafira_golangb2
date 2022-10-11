package routers

import (
	"gin_framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	//variable router kita gunakan sebagai penampung untuk engine dari router yang kita dapatkan dari pemanggilan function gin.Default.
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	//meregistrasikan endpoint UpdateCar pada function StartServer yang terletak pada file carRouter.go di dalam folder routers.
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	return router
}

/*
Function StartServer akan kita gunakan untuk menjalankan
server dari aplikasi kita. Function ini mengembalikan
suatu data dengan tipe data struct *gin.Engine yang berasal
dari Gin, dan kita gunakan untuk menjalankan server, sebagai
multiplexer dari routing dan lain-lain.
*/

/*
 method POST untuk menghubungkan client dengan endpoint CreateCar.
 Kita memberikan 2 parameter terhadapt method POST ini. Parameter
 pertama adalah route path nya, sedangkan parameter kedua memerlukan
 handler atau endpoint nya. Endpoint harus merupakan sebuah function
 yang menerima satu parameter dengan tipe data *gin.Context
*/
