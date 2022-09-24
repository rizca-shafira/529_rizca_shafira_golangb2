/*
Introduction
Go -> bahasa pemrograman open source yg dibuat oleh google. Kelebihan golang diantaranya :
1. Clean & simple : syntax kompak dan hemat baris
2. Concurrent : ada fungsi thread
3. Fast : proses compile cepat seperti C dan didesain untuk multi-core computer
*/

/*
Installation (Windows)
- Dapatkan installer dari golang.co.id
- eksekusi "go version" utk mengecek versi Go
- path bisa dilihat dengan cd>echo %GOPATH%
*/

/*
Go Root & Go Path
- setelah menginstall akan muncul environment variable GOROOT tempat go terinstall
- Pada umumnya, seluruh codingan Go diletakkan di satu workspace
- Go Workspace : direktori dimana path file disimpan  di dlm environment variable bernama GOPATH
*/

/*
Setup Go Workspace
- go workspace punya 3 sub folder :
	- bin : berisi exe hasil build
	- pkg : berisi file hasil compile
	- src : berisi project go
- go workspace berada di sub direktori home/go
- tambahkan path dari sub direktori tsb ke path variable bernama GOPATH
- periksa apakah setup path benar dgn perintah echo $GOPATH
*/

/*
Go Modules
- go modules -> manajemen dependensi bahasa go (seperti npm di js)
- dengan go modules, bisa membuat project go di luar go workspace
*/

/*
Initializing Project
mkdir projek-pertama //utk membuat go modules. nama go modules lbh baik sama dgn folder project
cd project-pertama
go mod init project-pertama
*/

/*
Command
- untuk run : go run main.go
- jika ada banyak package main di satu dir, eksekusinya tuliskan semua nama file sbg argumen go run : go run main.go lib.go
- go build digunakan untuk compile program.
- go run . akan menjalankan package main di dir tsb
- go get untuk mendownload package
- utk mendownload package terbaru gunakan flag -u
- command go get hrs dijalankan di dir project, kalau tdk akan diunduh ke go path
*/

/*
setiap project harus punya minimal 1 package main. file dgn package main akan dieksekusi pertama kali.
import digunakan untuk mengimport/memasukkan package lain ke program
main function dipanggil pertama kali pada eksekusi program. hrs ada di program.
*/

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
