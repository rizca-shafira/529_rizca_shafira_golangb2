//Defer #1
/*
Defer -> keyword bhs go yg digunakan utk memanggil sebuah function yg dimana proses eksekusi akan ditahan hingga block sebuah function selesai
fmt.Println pertama dipanggil dgn menggunakan defer. Text yg berada pada posisi paling bawah merupakan text yg ditampilkan menggunakan defer
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("defer function starts to excecute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")

	callDeferFunc() //callDeferFunc dipanggil sebelum hai everyone
	fmt.Println("Hai everyone!!")

	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc() //memanggil function deferFunc dgn keyword defer
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

//tulisan Defer func starts to execute tetap dijalankan sblm hai everyone, krn defer dipanggil di function callDeferFunc, bukan di main.
//jd akan dijalankan setelah semua kode di callDeferFunc sudah dijalankan
//function exit berasal dr package os utk menghentikan program scr paksa
//Tulisan Invoke with defer tdk ditampilkan, krn stlh fmt.Println, program dihentikan paksa oleh function os.Exit
//os.Exit menerima satu parameter berupa nilai dgn tipe data int yg digunakan sbg status exit nya
