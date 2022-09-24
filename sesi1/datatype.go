/*
Tipe data bahasa Go dibagi menjadi 4 kategori :
1. Basic type = number, string, boolean
2. Aggregate type = array, struct
3. Reference type = slice, pointer, map, function, channel
4. Interface type = interface
*/

/*
Number : terbagi jadi integer dan float
integer -> non desimal. dibagi menjadi uint (unsigned) dan int (bilangan positif)
berdasarkan besaran nilainya, dibagi lg menjadi int32 atau int64
sebagiknya tipe data disesuaikan dgn kebutuhan utk menghindari boros memory
float jg terbagi menjadi float32 dan float64 brdsrkn cakupan nilai desimal yg bisa ditampung
*/

/*
Nil dan Zero Value
Nil bkn tipe data, tetapi sebuah nilai variabel kosong
Zero value string = ""
Zero value bool = false
Zero value int = 0
Zero value float = 0.0

Zero value berbeda dgn nilai nill. Nil kosong. Tipe data yg bisa diset dgn nilai nill =
-pointer, tipe data function, slice, map, channel, interface()
*/
package main

import "fmt"

func main() {
	var desimalNumber float32 = 3.63
	fmt.Printf("desimalNumber: %f\n", desimalNumber)
	fmt.Printf("desimalNumber3: %.3f\n", desimalNumber)

	var boolVar bool = true
	fmt.Printf("is it permitted? %t\n", boolVar)

	var message string = "Haloo!\n"
	fmt.Printf(message)

	var message2 string = "Selamat Datang di 'Balikpapan'"
	fmt.Printf(message2)
}
