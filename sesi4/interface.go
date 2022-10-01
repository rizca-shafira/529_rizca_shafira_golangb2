package main

import (
	"fmt"
	"math"
)

/*
Interface -> tipe data bhs go yg merupakan kumpulan dr definisi satu atau lbh method
interface dpt digunakan jika kita telah mengimplementasi method yg didefinisikan oleh interface tsb melalui tipe data lainnya
*/

// berikut interface bernama shape dgn 2 method, area dan perimeter
type shape interface {
	area() float64 //method area return data bertipe float64
	perimeter() float64
}

// struct rectangle memiliki 2 field bernama width dan height
type rectangle struct {
	width, height float64
}

// struct circle mempunyai 1 field bernama radius yg memiliki tipe data float64
type circle struct {
	radius float64
}

// struct circle mengimplementasikan method area utk menghitung luas lingkaran
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// struct rectangle mengimplementasikan method area untuk menghitung luas persegi panjang
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var c1 shape = circle{radius: 5}              //var c1 memiliki tipe data interface shape dan diberikan nilai berupa struct circle
	var r1 shape = rectangle{width: 3, height: 2} //var r1 menampung struct rectangle dan bertipe interface shape

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	fmt.Println("Circle area:", c1.area())
	fmt.Println("Circle parameter", c1.perimeter())

	fmt.Println("Rectangle area:", r1.area())
	fmt.Println("Rectangle parameter", r1.perimeter())

	print("Rectangle", r1)
	print("Circle", c1)

	//muncul error spt ini apabila volume dipanggil dlm keadaan blm didefinisikan di daftar method
	//c1.(circle).volume()

	c1.(circle).volume()

	value, ok := c1.(circle)
	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}
}

/*
Kedua variable c1 dan r1 tersebut memiliki 2 tipe data yaitu struct yang telah dijadikan sebagai nilai dan interface shape. Inilah yang disebut dengan polymorphism atau dynamic type. Dengan mengimplementasikan interface, maka suatu variable dapat mempunyai 2 tipe data.
*/

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

/*
Ketika struct circle menambahkan method selain method yg didefinisikan interface shape, make var c1 tdk dapat menggunakan method tsb
contoh ditambah func (c circle) volume() float64
untuk mengatasi, digunakan type assertion
*/

/*
type assertion -> teknik untuk mengembalikan suatu tipe data interface ke tipe data aslinya
pd kasus ini ingin dikembalikan var c1 ke tipe data aslinya yaitu tipe data struct circle
format penulisan type assertion adalah namaVariable.tipedataasli
*/

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

/*
cek type assertion berhasil atau ga : berikan 2 var opsional sbg penampung dr hasil yg dikembalikan type assertionnya. contohnya var calue akan menerima tipe data asli dan nilai yg sdh diberikan kpd field dr struct circle jika type assertion berhasil dan var ok akan menerima nilai true jika type assertion berhasil dan nilai false jika type assertion gagal
cek var ok, jika true, maka akan menampilkan nilai keseluruhan struct circle dan menampilkan hasil dr pemanggilan method volume yg dimiliki struct circle
*/
