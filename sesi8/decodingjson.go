//cara mendecode json ke dlm struct
/*
membuah data JSON sederhanamenggunakan tanda backtick ``.
Kemudian pada line 25,kita menggunakan function json.Unmarshal
untukmendecode data JSON kepada struct Employee.Argumen pertama
dari function json.Unmarshal menerima sebuah nilai dengan tipe
data slice of byte.Pada argument pertama itulah kita meletakkan
data JSONnya tetapi harus kita ubah terlebih dahulu menjadi tipedata
slice of byte.Kemudian pada argumen kedua, kita meletakkan pointerdari
variable result agar setelah data JSON berhasil didecode, datanya
akan disimpan kepada variable result.
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

/*
erdapat sebuah format tulisanseperti `json:”full_name”`, `json:”email”`, dan `json:”age”`.
Tulisan-tulisan tersebut disebut sebagai tag. Tag kitagunakan ketika kita
ingin mendecode data-data sepertiJSON, form data, hingga xml kemudian kita
simpan datadecode tersebut kepada field-field struct nya.
Tag yang kita buat harus kita sesuaikan dengan field padadata JSONnya.
*/

func main() {
	var jsonString = `
	{
		"full_name":"Airell Jordan",
		"email":"airell@gmail.com",
		"age":23
	}
	`
	var result Employee

	var result2 map[string]interface{} //decode json ke map

	var temp interface{} //decode json to empty interface

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var err2 = json.Unmarshal([]byte(jsonString), &result2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	var err3 = json.Unmarshal([]byte(jsonString), &temp) //to empty interface
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}

	var result3 = temp.(map[string]interface{}) //json to empty interface

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)

	fmt.Println("full_name:", result2["full_name"])
	fmt.Println("email:", result2["email"])
	fmt.Println("age:", result2["age"])

	fmt.Println("full_name:", result3["full_name"])
	fmt.Println("email:", result3["email"])
	fmt.Println("age:", result3["age"])
}

/*
Kita juga bisa men-decode data JSON kepadasebuah tipe data map.
Kita tidak perlu membuat tag seperti yang kitalakukan pada
sebuah struct.Jika kita jalankan, maka hasilnya akan seperti padagambar kedua.
*/

/*
men-decode data JSON kepada sebuah tipe dataemty interface.
Caranya seperti pada gambar di sebelah kanan.Perlu diingat
disini bahwa ketika kita ingin mengaksesfield-fieldnya, maka
harus dilakukan type assertion dari emptyinterface menjadi
tipe data map[string]interface{}
*/
