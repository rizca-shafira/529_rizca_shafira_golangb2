/*
Kondisional digunakan untuk mengatur alur program/kpn harus mengeksekusi program
Ada 2 macam kondisional : if else & switch
*/

/*
Temporary variable -> bisa menggunakan variable yg hanya bisa digunakan pada suatu scope block kondisional
*/

/*
Switch
*/

package main

import "fmt"

func main() {
	var currentYear = 2021
	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat SIM")
	} else {
		fmt.Println("Kamu sudah boleh membuat SIM")
	}

	var score = 8
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//Switch with relational operator
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		fmt.Println("study harder")
	}

	//Fallthrough keyword -> melanjutkan pengecekan ke case selanjutnya meskipun suatu case sdh terpenuhi kondisinya
	score = 6
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough //case ini sdh terpenuhi tp lanjut ngecek case selanjutnya
	case score < 5:
		fmt.Println("It's ok!") //case ini jg terpenuhi
	default:
		fmt.Println("Study harder")
	}

	//Nested Conditions -> di dalam proses kondisional terdapat kondisional lg
	//contoh : di dalam if ada switch atau ada if lain
}
