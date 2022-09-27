/*
Pd go, tiap folder akan dianggap sbg suatu package
Kita dpt menggunakan var atau tipe data dr package lain, asalkan var atau tipe data tsb sdh tereksport dr package nya
cara eksport adlh dgn mengawali penulisan vat atau tipe data yg diekspor dgn huruf kapital atau upper case
*/

package main

import "golang_sesi3/helpers"

func main() {
	helpers.Greet()
	//helpers.greet()
	//Error ini terjadi karena function greet bukan merupakan suatu function yang ter-eksport dari package helpers karena penulisan function nya dimulai dari huruf kecil, maka otomatis function greet tidak ter-eksport
	//# command-line-arguments .\exported.go:13:10: greet not exported by package helpers
	//di method ada Invokegreet yang digunakan untuk memanggil function greet yang tidak ter-eksport dari package helpers. Maka dari itu, jika kita ingin memanggil function greet dari file main.go, maka kita perlu memanggil method Invokegreet.
	var person = helpers.Person{}
	person.Invokegreet()
	//terdapat suatu function yang akan dieksekusi terlebih dahulu sebelum function main. Function tersebut bernama init.
	//Karena file init.go merupakan file yang berada dalam root direktori project kita, maka file init.go akan dianggap sebagai bagian dari package main.

	//Jika kita ingin menjalankan seluruh file yang berada di root direktori kita, maka kita menggunakan perintah go run *.go.
	//func init dieksekusi sblm main
}
