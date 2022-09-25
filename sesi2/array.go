/*
Array tipe data utk menyimpan kumpulan data dgn tipe yg sama
bersifat fixed length, ukuran hrs ditentukan di awal pembuatan array
*/

/*
2 cara deklarasi array : deklarasi tanpa nilai dan langsung dgn nilai
*/

package main

import "fmt"

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Airell", "Nanda", "Mailo"}
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings) //flag v untuk memformat array agar terlihat tipe array, elemen, dan panjang

	//modify element through index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "melon"
	fmt.Printf("%#v\n", fruits)

	//loop through element
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Indes: %d, Value:%s\n", i, fruits[i])
	}

	//multidimensional array -> array di dalam array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
