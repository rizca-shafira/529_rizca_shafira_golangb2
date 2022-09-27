/*
Pointer -> var yg digunakan utk menyimpan alamat memori var lain
var yg memiliki reference atau alamat memory yg sama saling berhubungan dan nilainya sama
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber

	var first int = 4
	var second *int = &first //ketika ingin memberi value pointer, perlu ampersand
	fmt.Println("first value:", first)
	fmt.Println("first memory:", &first)  //menampilkan alamat memory ver first dgn ampersand &
	fmt.Println("second value:", *second) //menampilkan nilai asli dgn tanda asterisk *
	fmt.Println("second memory:", second)

	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value):", firstPerson)
	fmt.Println("firstPerson(memori):", &firstPerson)
	fmt.Println("secondNumber (value):", *secondPerson)
	fmt.Println("secondNumber(memori):", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe" //cara merubah variabel lwt pointernya

	fmt.Println("firstPerson (value):", firstPerson)
	fmt.Println("firstPerson(memori):", &firstPerson)
	fmt.Println("secondNumber (value):", *secondPerson)
	fmt.Println("secondNumber(memori):", secondPerson)

	//pointer as a parameter
	var a int = 10
	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
