//Strings in Depth
//Intro
//String -> kumpulan tipe data byte yang diletakkan di dalam slice atau slice of bytes
//tipe data byte merupakan alias uint8
//ketika melakukan indexing thd string, akan didapat representasi byte nya

// Looping Over String (byte-to-byte)
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte:%d\n", i, city[i]) //city[i] tdk menghasilkan char tp byte nya
	}

	//membuat string dengan menggunakan slice of bytes
	var cities []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(cities))

	//looping over string (rune-by-rune)
	//saat menggunakan len(), yg didapatkan adalah jumlah byte
	//bkn jmlh char
	str1 := "makan"
	str2 := "mâcan" //krn â accented char yg terdiri dr 2 byte
	fmt.Printf("str1 byte length -> %d\n", len(str1))
	fmt.Printf("str1 byte length -> %d\n", len(str2))

	//jika ingin jumlah char nya, perlu diubah jadi rune dulu
	//cara menggunakan utf8.RuneCount
	fmt.Printf("str1 byte length -> %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str1 byte length -> %d\n", utf8.RuneCountInString(str2))

	//cara menggunakan range loop
	for i, s := range str2 {
		fmt.Printf("undex -> %d, string -> %s\n", i, string(s))
	}
}
