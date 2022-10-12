/*
menggunakan package database/sql denganmengintegrasikannya dengan sebuah sql driver.
Karena sekarang kita akan menggunakan database Postgresql, makadari itu kita perlu
terlebih dahulu menginstall driver dari Postgresql
go get -u github.com/lib/pq
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	//tanda underscore saat mengimpor krn tidak akan menggunakan syntax apapun dari driver Postgresql.
	//Yang kita perlu lakukan adalah hanya meregristrasikan nya dengan cara
	//mengimportnya saja agar package databases/sqlmengetahui tentang driver jenis apa yang kita pakai
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "db-go-sql"
)

//var untuk menyimpan seluruh informasidari database Postgresql pada sistem kita

var (
	db *sql.DB
	//Variable db memiliki tipe data pointer dari struct sql.DB
	//yangnantinya akan kita reassign dengan object dari sql.DB
	//pada saatkita membangun koneksi pada function main.
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       string
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//kita menggabungkan seluruh info dari Postgresqlyang
	//telah kita buat pada line diatasnya menjadi sebuah
	//stringpanjang yang kita simpan pada variable psqlInfo
	db, err = sql.Open("postgres", psqlInfo)
	//function Open berasaldari package database/sql.
	//Function Open menerima 2parameter yaitu nama dari
	//driver yang kita pakai, dan sebuahstring berupa
	//informasi tentang bagaiman cara database/sqlmembangn
	//koneksi nya kepada database kita sesuai dengandriver yang kita gunakan.

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	/*
		Kita dapatmenggunakan method Ping karena setelah penggunaan
		functionOpen berhasil tereksekusi, maka function Open akanmengembalikan
		nilai berupa pointer dari struct sql.DB yangdimana struct
		tersebut memiliki method Ping.Pemanggilan method Ping
		merupakan hal yang sangat pentingkarena function Open
		tidak akan membangun koneksi kedatabase, melainkan hanya
		berfungsi untuk memvalidasiarugmen-argumen yang diberikan.
		Dengan memanggil methodPing, maka kita dapat membangun koneksi
		ke databasesekaligus memeriksa jika string panjang berupa info
		yang kitaberikan pada function Open merupakan info yang 100% valid.
	*/
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected to database")

	//CreateEmployee() //dikomen untuk mencoba func getemployee
	//GetEmployee() //dikomen karena mau coba func update elmpoyee
	//UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}
	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1,$2,$3,$4)
	Returning *
	`
	/*
		statement untuk melakukan create datapada database.
		Lalu tanda $1, $2 dst.. merupakanrepresentasi dari nilai-nilai
		yang akan kita masukkan nantinya.Lalu terdapat penulisan
		Returning * yang memiliki arti bahwakita ingin mendapat
		nilai-nilai dari seluruh field yang berasaldari data yang baru dibuat.
	*/

	err = db.QueryRow(sqlStatement, "Airell Johnson", "airell@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	/*
		method QueryRow.Method in digunakan untuk mengeksekusi
		sebuah query sqltergantung dari statement sql yang kita
		berikan. Karenastatement sql yang kita berikan bertujuan
		unutk create data,maka method QueryRow ini akan berfungsi
		untuk membuatdata baru dengan nilai yang kita berikan pada
		parameter keduadari method QueryRow. Parameter kedua dari
		QueryRowmerupakan sebuah parameter variadic dan nilai-nilai
		yang kitaberikan adalah nilai-nilai yang akan me-replace
		statement yangkita buat pada line 55 VALUES($1, $2, $3, $4).

		method QueryRow kita chaning denganmethod Scan agar kita dapat mendapatkan
		nilai-nilaibalikan dari statement yang telah kita buat. Karenastatement
		kita menggunakan Returning *, maka kitaakan mendapatkan seluruh seluruh
		nilai dari field-fieldyang berasal dari data yang baru dibuat.
		Dannilai-nilai balikan tersebut akan kita simpan kedalamfield-field dari variable employee
	*/

	if err != nil {
		panic(err)
	}

	fmt.Printf("New employee data: %+v\n", employee)
}

func GetEmployee() {
	var results = []Employee{}
	sqlStatement := `SELECT * from employees`
	rows, err := db.Query(sqlStatement)
	/*
		method Query biasa digunakanuntuk mendapatkan
		banyak data dari suatu tablepada database dikarenakan
		method ini dapatmengembalikan satu atau lebih rows
		dari suatutable pada database. Dan kita juga harus
		menutuprows tersebut dengan method Close yang
		kitapanggil denang keyword defer.
	*/

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		/*
			melakukan perulangan/loopingsebanyak data yang kita
			dapatkan dengan acuan rows. Next. Method rows. Next
			akan menghasilkannilai true selama data nya masih ada, namun
			jikasudah tidak ada maka dia akan me-return falsedan proses looping akan terhenti
		*/
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
		/*
			melakukan proses scanning
		*/
		if err != nil {
			panic(err)
		}

		results = append(results, employee)
		/*
			setelah proses scaning selesaiuntuk satu row,
			maka kita akan memasukkan datadari row tersebut kedalam variable results
		*/

	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airelhidayat@gmail.com", "CurDevs", 24)
	/*
		menggunakanmethod Exec untuk mengupdate data.
		Sebetulnya kita juga bisa mengupdate sebuah
		data denganmenggunakan method QueryRow, namun
		dianjurkan untuk menggunakan method Exec untuk
		melakukanproses insert, update, dan delete.
		Dengan menggunakan method Exec, kita tidak dapat
		mendapatkan data yangbaru saja diupdate maupun
		data yang baru kita buat, namun kita dapat menggunakan
		method RowsAffecteduntuk mengetahui berapa jumlah
		row atau data yang baru diupdate. Method RowsAffected
		didapatkan dariinterface sql.Result yang merupakan
		nilai kembalian dari method Exec.
	*/
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}
