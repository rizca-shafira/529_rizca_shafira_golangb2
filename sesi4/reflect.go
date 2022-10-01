/*
Reflect -> untuk melakukan inspeksi variabel, mengambil informasi dr variabel tsb, atau memanipulasinya
Cakupan info yg bisa didapat lewat reflection : melihat struktur variabel, tipe, nilai pointer, dan banyak lagi
fungsi yg paling penting diketahui yaitu reflect.ValueOf() dan reflect.TypeOf()

reflect.ValueOf() mengembalikan objek dlm tipe reflect.Value yg berisi informasi berhubungan dgn nilai pd variabel yg dicari. memiliki parameter yg bisa menampung segala jenis tipe data
salah satu method reflect.Value yaitu Type() -> mengembalikan tipe data variabel yg bersangkutan dlm bentuk string
reflect.TypeOf() mengembalikan objek dlm tipe reflect.Type yg berisikan info berhubungan dgn tipe data variabel yg dicari
*/

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe variabel:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel:", reflectValue.Int)
	}

	fmt.Println("tipe variabel:", reflectValue.Type())
	fmt.Println("nilai variabel:", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	_ = nilai

	var s1 = &student{Name: "john wick", Grade: 2} //instance struct student
	fmt.Println("nama:", s1.Name)

	var reflectValues = reflect.ValueOf(s1)
	var method = reflectValues.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama:", s1.Name)
}

// method baru di struct student dgn nama SetName
func (s *student) SetName(name string) {
	s.Name = name
}

/*
Statement reflectValue.Int() menghasilkan nilai int dr var number
untuk menampilkan nilai var reflect, hrs dipastikan dulu tipe datanya
ketika tipe data int, bisa menggunakan method Int()

contoh method lain : reflectValue.String(), reflectValue.Float64(), float64
*/

/*
fungsi yg digunakan hrs sesuai dgn tipe data nilai yg ditampung variabel
kalau beda, error

contoh : pd variabel float64, tdk bisa menggunakan method String()
cara nya bisa dgn cek dulu apa tipe datanya menggunakan method Kind()
*/

/*
list konstanta tipe data dan method yg bisa dipakai di reflect go https://pkg.go.dev/reflect#Kind
*/

/*
accessing value using interface method
jika nilai hanya perlu ditampilkan ke output, bisa menggunakan .Interface()
*/

/*
accessing value using interface method
method Interface() mengembalikan nilai interface kosong atau interface{}
nilai aslinya sendiri bisa diakses dgn meng-casting interface kosong
*/

/*
identifying method information
informasi mengenai method jg bisa diakses lewat reflect, syaratnya sama seperti pengaksesan property, yaitu bermodifier public
pd contoh, method SetName akan diambil lewat reflection
*/

/*
Var s1 merupakan instace struct student
awalnya property Name pd variabel tsb berisi "john wick"
setelah itu, refleksi nilai objek tsb diambil dan refleksi method SetName jg diambil
pengambilan refleksi method dilakukan menggunakan MethodByName dgn argumen nama method yg diinginkan atau bisa jg lewat indeks methodnya (menggunakan Method(i))
stlh refleksi method yg dicari didapat, Call() dipanggil untuk eksekusi method
*/

/*
Jika eksekusi method diikuti pengisian parameter, maka parameternya hrs ditulis dlm bentuk array []reflect.Value berurutan sesuai urutan deklarasi parameternya
nilai yg dimasukkan ke array tsb harus dlm bentuk reflect.Value (gunakan reflect.ValueOf() untuk pengambilannya)
*/
