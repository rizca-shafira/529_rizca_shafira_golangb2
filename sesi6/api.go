/*
membuat sebuah API yang menggunakan 2 method berupa GET dan POST
dimana akan kita gunakan untuk mendapatkan data employee dan membuat data employee baru
*/

/*
data-data employee dlm tipe data slice yang berisikan tipe data structEmployee.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Caisar", Age: 22, Division: "Finance"},
	{ID: 3, Name: "Fahmi", Age: 24, Division: "IT"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)
	//Sebelumnya kita mencoba untuk mengirimkan requestmenggunakan Postman, kita perlu meregistrasikan functioncreateEmployee kedalam routingan kita
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	//function getEmployees kita gunakan untuk mendapatkandata-data empoyee dan pada umumnya untuk mendapatkandata dari server menggunakan method GET.
	w.Header().Set("Content-Type", "application/json") //method Header dari interfacehttp.ResponseWriter yang kemudian di chaining denganmethod Set
	//untuk menentukan bentuk dari dataresponse yang ingin kita kirimkan kepada client. Karena saat inikita ingin mengirim response data dalam bentuk JSON, makakita dapat mengatur Content-Type nya menjadiapplication/json dalam method Set.

	if r.Method == "GET" { //melakukan pengecekan method.
		json.NewEncoder(w).Encode(employees)
		//mengkonversi data employees menjadidata berbentuk JSON untuk dikirimkan kepada client denganmenggunakan method NewEncoder yang berasal dari packagejson, yang kemudian kita chaining dengan method Encodeuntuk mengkonversi datanya menjadi bentuk JSON.
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
	//jika method yang dikirimkan oleh client bukanmethod GET, maka kita akan mengirimkan response errordengan menggunakan function Error dari package http. Lalu http.StatusBadRequest merupakan sebuah konstanta daripackage http.StatusBadRequest yang merepresentasikanstatus400.
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" { //menggunakan method POST untuk menambah data employee baru
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")
		//mendapatkan nilai inputform dari client dengan menggunakan method FormValueyang berasal dari struct http.Request

		convertAge, err := strconv.Atoi(age)
		//mengkonversi input age dari tipedata string menjadi tipe data int karena seluruh nilai input yangkita dapat dari client akan memiliki tipe data string, sedangkanfield age pada struct Employee memiliki tipe data int.

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		//variable newEmployee yangakan menampung nilai-nilai input dari client yang kemudianakan kita append atau masukkan kedalam slice employees.
		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
