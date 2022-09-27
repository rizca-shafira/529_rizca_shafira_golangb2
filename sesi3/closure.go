/*
closure -> anonymous function atau function tanpa nama yg dpt disimpan sbg sebuah variabel maupun dijadikan nilai return function
*/

// declare closure in variable
package main

import (
	"fmt"
	"strings"
)

// using alias
type isOddNum func(int) bool

func main() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...)) //mengumpulkan angka genap

	// IIFE -> immediately-invoked function expression
	//merupakan closure yg bisa langsung tereksekusi ketika pertama kali dideklarasikan
	// tdk perlu dipanggil dgn tanda kurung krn closure IIFE tereksekusi pd saat dideklarasikan
	var isPalindrome = func(str string) bool {
		var temp string = ""
		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")
	fmt.Println(isPalindrome)

	//closure as a return value
	var studentLists = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marco"}
	var find = findStudent(studentLists)
	fmt.Println(find("airell"))

	//callback
	numbers = []int{2, 5, 8, 10, 3, 99, 23}
	var find2 = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find2)

	//callback using alias
	var find3 = findOddNumbers2(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find3)
}

// menerima tipe data slice of string, mereturn closure
// closure yg direturn menerima parameter dgn tipe data string yg digunakan utk mencari data student dr parameter yg diterima func findStudent
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position)
	}
}

// callback -> closure yg dijadikan sbg parameter sebuah function
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

// untuk mempersingkat penulisan parameter callback, bisa digunakan alias
func findOddNumbers2(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
