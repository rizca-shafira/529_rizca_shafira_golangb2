/*
Slice -> tipe data yang mirip array, tp tdk fixed length
slice termasuk reference type, jika kita copy suatu slice dan kita ubah elemennya, maka slice semulanya jg ikut terubah
*/

package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "mango"}
	_ = fruits
	fmt.Printf("%#v\n", fruits)

	//bisa jg membuat slice dgn make function
	var buah = make([]string, 3)
	_ = buah
	fmt.Printf("%#v\n", buah) //slice blm berisi apapun tp diketahui isinya 3 elemen

	//append function -> untuk menambahkan element pada slice
	//fungsi append mengembalikan nilai dari slice yg ditambahkan jd harus disimpan ke variabel
	buah = append(buah, "watermelon") //jgn lupa hsl append disimpan di buah skrg
	fmt.Printf("%#v\n", buah)         //append masuknya di elemen terakhir

	//append function with ellipsis -> jika ingin memasukkan elemen array ke array lainnya, gunakan tanda elipsis
	buah = append(fruits, buah...)
	fmt.Printf("%#v\n", buah)

	//copy function -> mengcopy seluruh elemen pd sebuah slice ke slice lainnya
	//elemen di slice yg lainnya terreplace sama hasil copy ini
	var arrlama = []string{"elmn1", "elmn2"}
	var arrbaru = []string{"halo", "nice"}
	nn := copy(arrlama, arrbaru)
	fmt.Printf("Arrlama => %#v\n", arrlama)
	fmt.Printf("Arrbaru => %#v\n", arrlama)
	fmt.Printf("Copied element => %#v\n", nn)

	//slicing -> mendapat elemen dari sebuah slice dan menentukan elemen dr index yg ingin didapatkan
	//start -> awal index yg diakses, stop -> index akhirnya
	var ambil = buah[1:4] //index 1  - 3
	fmt.Printf("%#v\n", ambil)
	var ambil1 = buah[0:] //index 0 - terakhir
	fmt.Printf("%#v\n", ambil1)
	var ambil2 = buah[:3] //index 0 - 3
	fmt.Printf("%#v\n", ambil2)
	var ambil3 = buah[:] //index 0 - terakhir
	fmt.Printf("%#v\n", ambil3)

	//combining slicing and append untuk mereplace index tertentu
	var froots = []string{"apples", "banana", "mango", "durian", "pineapple"}
	froots = append(fruits[:3], "rambutan") //meng-append elemen rambutan ke array yg isinya elemen fruits[0] - fruits[2]
	fmt.Printf("froots : %#v\n", froots)

	/*
		setiap dibuat slice pd bahasa go, scr otomatis dibuat array tersembunyi bernama backing aray.
		backing array digunakan utk menyimpan elemen slice.
		bhs go mengimplementasi slice sbg struktur data yg disebut slice header
		yg terdiri dari :
		1. alamat memori/address backing array
		2. panjang slice yg bisa didapatkan dgn fungsi len
		3. kapasitas slice yg bisa didapatkan dgn fungsi cap
	*/

	/*
		ketika kita mencoba mendapat elemen dg slicing, tdk dibuat backing array baru tp berbagi backing array yg sama
		masalahnya, jika kita slicing dr fruits1 ke fruits2. trus elemen di fruits2[0] kita ganti, maka elemen fruits1 jg ikut terganti
	*/

	/*
			Cap function -> mengetahui kapasitas array/slice
			|variable|slicing|result|len()|cap()|
			|fruits1|fruits[1:]|[fruits, fruit, fruit, fruit]|4|4|
			|fruits2|fruits[0:3]|[fruit, fruit, fruit, ..]|3|4| //kapasitas ttp sama
		|fruits3|fruits[1:]|[..., fruit, fruit, fruit]|3|3| //kapasitas berubah krn elemen ke 1 menjadi elemen ke 0 slice baru
	*/

	//Creating new backing array -> menggunakan  fungsi append
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}
