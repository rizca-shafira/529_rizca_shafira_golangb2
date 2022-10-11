//Gin adalah sebauh framework untuk bahasa Go yang digunakan untuk keperluan http routing.
//Dengan menggunakan Gin, maka akan sangat mempermudah kita dalam membuat Rest API yang sudah pasti memerlukan fitur routing
//Kali ini kita akan membuat sebuah Rest API sederhana yang akan melakukan CRUD data mobil

//go mod init example/gin
//jalankan perintah go get -u github.com/gin-gonic/gin
//File main.go menjadi entry point dari servernya.
//Folder routers akan menjadi tempat kita dalam menaruh konfigurasi dari routingnya
//sedangkan controllers akan menjadi tempat kita untuk menaruh endpoint-endpoint yang kita perlukan.

package main

import "gin_framework/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
