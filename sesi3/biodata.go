//Assignment 1
//Rizca Shafira Salsabila Makasuci
//149368582100-529
//Batch2

package main

//import package
import (
	"fmt"
	"os"
)

// struct Teman
type Teman struct {
	absen     string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	//check if absen is included as argument
	if len(os.Args) < 2 {
		//error handling if absen is not included as argument
		fmt.Println("Input absen yang dicari dengan format 'go run biodata.go [absen]'")
	} else {
		input_absen := os.Args[1] //getting the absen argument
		biodata(input_absen)
	}
}

func biodata(absen string) {
	//initializing var teman including values
	var teman = []Teman{
		{absen: "1", nama: "Rizca", alamat: "Balikpapan", pekerjaan: "Mahasiswa", alasan: "Mencari pengalaman"},
		{absen: "2", nama: "Shafira", alamat: "Bandung", pekerjaan: "Helpdesk", alasan: "Meningkatkan skill"},
		{absen: "3", nama: "Salsabila", alamat: "Jakarta", pekerjaan: "Admin", alasan: "Pindah haluan karir"},
		{absen: "4", nama: "Makasuci", alamat: "Semarang", pekerjaan: "Frontend Developer", alasan: "Mempelajari bidang backend"},
		{absen: "5", nama: "Fahmi", alamat: "Bali", pekerjaan: "Pemain bola", alasan: "Mengoleksi sertifikat"},
	}
	// var found is used to search if input absen is located in var teman
	var found = false

	// search for element with matching absen as input by looping through the slice
	for i := range teman {
		if teman[i].absen == absen {
			found = true
			fmt.Println("Nama:", teman[i].nama)
			fmt.Println("Alamat:", teman[i].alamat)
			fmt.Println("Pekerjaan:", teman[i].pekerjaan)
			fmt.Println("Alasan:", teman[i].alasan)
		}
	}
	// error handling if input absen doesn't exist in var teman
	if !found {
		fmt.Println("Absen yang diinput tidak terdaftar")
	}
}
