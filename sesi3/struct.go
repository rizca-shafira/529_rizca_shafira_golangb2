/*
Struct-> tipe data berupa kumpulan dr berbagai field dan method
*/
package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type karyawan struct {
	division string
	person   Person
}

func main() {
	var employee Employee
	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curiculum Developer"
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	//inisialisasi dan pemberian nilai struct
	//%+v untuk memformat struct menjadi string
	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employee2: %+v\n", employee2)

	//pointer to a struct
	var employee3 *Employee = &employee2
	fmt.Println("Employee2 name:", employee2.name)
	fmt.Println("Employee3 name:", employee3.name)
	employee3.name = "Rizca" //ketika employee3 merubah nilai fieldnya, maka var employee2 jg ikut terubah
	fmt.Println("Employee2 name:", employee2.name)
	fmt.Println("Employee3 name:", employee3.name)

	//embedded struct
	//struct dpt mengandung struct lain
	var karyawan1 = karyawan{}

	karyawan1.person.name = "Airelll"
	karyawan1.person.age = 23
	karyawan1.division = "Curiculum Developer"

	fmt.Printf("%+v", karyawan1)

	//Struct anonymous struct
	//struct yg tdk dideklarasikan di awal sbg tipe data struct baru, tp langsung dideklarasikan dgn pembuatan variabel
	var KAR = struct {
		person   Person
		division string
	}{}
	KAR.person = Person{name: "Airella", age: 23}
	KAR.division = "Curiculum Dev"

	var KAR2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Amanda", age: 23},
		division: "Finance",
	}
	fmt.Printf("Employee1:%+v\n", KAR)
	fmt.Printf("Employee1:%+v\n", KAR2)

	//slice of struct
	var people = []Person{
		{name: "Bram", age: 23},
		{name: "Nando", age: 23},
		{name: "Cahya", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	//slice of anonymous struct
	var emp = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Lucu", age: 24}, division: "Dev"},
		{person: Person{name: "Haha", age: 88}, division: "Chaos"},
		{person: Person{name: "Doi", age: 44}, division: "Test"},
	}

	for _, v := range emp {
		fmt.Printf("%+v\n", v)
	}
}
