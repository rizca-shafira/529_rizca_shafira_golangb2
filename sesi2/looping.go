/*
Looping -> proses berulang dimana proses hanya berhenti jika memenuhi suatu kondisi
di Go, looping hanya menggunakan for loop
*/

package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	var j = 0
	for j < 3 {
		fmt.Println("Angka", j)
		j++
	}

	var k = 0
	for {
		fmt.Println("Angka", k)
		k++
		if k == 3 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue //jika var i ganjil, looping berlanjut dan mengabaikan if kedua dan print
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//nested looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	//label -> break dan continue hanya berulang di bagian label
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke- ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 { //jika i = 2, looping pertama atau terluar akan dihentikan
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
