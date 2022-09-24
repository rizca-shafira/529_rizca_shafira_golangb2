/*
Naming Convention Variable
camelCase -> currentAge
KAPITALISASI AKRONIM -> HTPP, URL
*/

package main

import "fmt"

func main() {
	// Variable with data type
	var name string = "rizca"
	var age int = 23

	fmt.Println("Nama saya: ", name)
	fmt.Println("Umur saya: ", age)

	//Variable tanpa assignment
	var namaLengkap string
	namaLengkap = "Rizca Shafira"
	fmt.Println("Nama lengkap: ", namaLengkap)

	//Variable without data type -> type interference
	var shoeSize = 41
	fmt.Printf("%T, %d", shoeSize, shoeSize)
	fmt.Println("")

	//Short Declaration -> tdk perlu menggunakan var lg, tapi harus langsung diberi nilai
	minumanFav := "CocaCola"
	fmt.Printf("%T, %s\n", minumanFav, minumanFav)

	//Multiple Variable Declarations
	var one, two, three int
	one, two, three = 1, 2, 3
	fmt.Println(one, two, three)

	var namaJalan, noRumah = "Digital", 12
	fmt.Println(namaJalan, noRumah)

	//Underscore variable -> digunakan untuk mencegah error krn var yg tdk terpakai.
	var tdkTerpakai string
	_ = tdkTerpakai
	var tdkTerpakai1, tdkTerpakai2 string
	_, _ = tdkTerpakai1, tdkTerpakai2
}
