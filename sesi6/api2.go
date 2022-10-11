/*
func getEmployees diganti dari api.go karena mau ditampilkan di html
*/
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")
		//function template.ParseFilesyang berasal dari package html/template yang digunakanuntuk mem-parsing file html kita.
		//Karena file html yang barukita buat bernama template.html, maka dari itu kita meletakkannama file html kita sebagai argumen functiontemplate.ParseFiles
		//function tersebut akanmengembaikan suatu tipe data struct template
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		//Template yang mempunyai sebuah method bernama Executeyang digunakan untuk memberikan response kepada clientberupa template html. Parameter kedua dari method Executedigunakan untuk mengirimkan data yang ingin kita tampilkanpada file html kita.
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
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
