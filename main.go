package main

import "fmt"

func main() {
	/*Program Go Pertama*/
	fmt.Println("Hello Rizca")

	//variabel
	var nama string = "rizca shafira"
	var age int = 23
	var one, two, three string = "1", "2", "3"
	fmt.Println("Nama:", nama)
	fmt.Println("Umur: ", age)
	fmt.Println(one, two, three)
	fmt.Printf("Halo, namaku %s. Aku %d tahun", nama, age)
	fmt.Println("")

	//underscore variabel
	var testVariable string
	var oneName, twoName, dataUmur = "Riz", "Sha", 25
	_, _, _, _ = testVariable, oneName, twoName, dataUmur
	fmt.Printf("tes underscore variabel > %T", twoName)
}
