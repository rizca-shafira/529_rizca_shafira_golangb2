//Error -> tipe data go yg digunakan untuk mereturn sebuah error jika ada error yg terjadi
//error umumnya dieturn pd posisi terakhir dr nilai2 return sebuah function
//contoh : pd funct strconv.Atoi untuk mengkonversi tipe data string menjadi int, func strconv.Atoi akan mereturn error jika nilai dgn tipe data string tsb tdk bisa dikonversi menjadi int.

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	defer catchErr()
	var number int
	var err error
	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("Password has to have more than 4 characters")
	}

	return "Valid password", nil
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

//ketika tdk ada error yg direturn, variabel err menjadi nil krn zero value tipe data error nil
//method Error menampilkan pesan jika tjd sebuah error

//kita jg bisa membuat pesan error kita sendiri dgn menggunakan method New.
//fmt.Scanln digunakan utk menangkap input yg ditulis di terminal. lalu dpt diberikan nilai pointer pd var password agar inputan dpt disimpan di var password

/*
Panic digunakan utk menampilkan stack trace error sekaligus menghentikan flow goroutine (termasuk goroutine main)
setelah ada panic, apapun setelahnya tdk akan dieksekusi, kecuali proses yg sdh didefer sebelumnya (akan muncul sebelum panic error)
informasi error yg ditampilkan panic adalah stack trace error jd sgt mendetail
*/

/*
Recover -> function digunakan utk menangkap error yg terjadi pd sebuah goroutine
function catchErr menangkap error panic dgn menggunakan function recover
dgn recover, pesan dr panic error tdk ditampilkan krn sudah ditangkap recover
*/
