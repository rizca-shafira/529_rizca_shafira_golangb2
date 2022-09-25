/*
Ailase -> fitur yg digunakan sbg nama alternatif dari tipe data yang sudah ada
tipe data dgn nama berbeda memiliki arti tipe datanya berbeda
kecuali : tipe data byte = uint8 dan rune = uint32
*/

package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte
	b = a //no error krn tipe datanya sama
	_ = b

	//mendeklarasi tipe data dgn nama yg dibuat sendiri
	//terdapat tipe data bernama second yg merupakan alias uint
	//var hour punya tipe data second
	//saat menggunakan flag %T, terlihat bahwa tipe data hour uint
	type second = uint
	var hour second = 3600
	fmt.Printf("hour type:%T\n", hour) // hour type : uint
}
