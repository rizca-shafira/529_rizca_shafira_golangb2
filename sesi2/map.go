/*
Map -> menyimpan data dalam key:value pairs, mirip seperti object di javascript
semua key value punya tipe data yg static
setiap key hrs unik
termasuk reference type (dicopy, mengganti map lain)
zero value map = nil
*/

/*
ketika ingin membuat map, harus segera diinisialisasi
*/

package main

import "fmt"

func main() {
	var person map[string]string //deklarasi. tipe data key dan value string
	person = map[string]string{} //inisialisasi
	person["name"] = "Airell"
	person["age"] = "23"
	person["address"] = "Jalan Sudirman"
	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	//bisa juga diberi value dengan key nya secara langsung
	var manusia = map[string]string{
		"name": "rizca",
		"age":  "23",
	}
	fmt.Println("name:", manusia["name"])
	fmt.Println("age:", manusia["age"])

	//looping with map dgn range loop
	for key, value := range manusia {
		fmt.Println(key, ":", value)
	}

	//delete value. argumen 1 : map yg ingin dihapus atau variabel tempat map disimpan. argumen 2: key yg menyimpan value
	delete(manusia, "age")
	fmt.Println("manusia stlh dihapus:", manusia)

	//detect a value
	value, exist := person["age"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("No value")
	}

	//combining slice with map
	var org = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "22"},
		{"name": "Mailo", "age": "15"},
	}
	for i, org := range org {
		fmt.Printf("Index:%d, name:%s, age:%s\n", i, org["name"], org["age"])
	}
}
