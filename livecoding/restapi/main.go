package main

import (
	"fmt"
	"log"
	"net/http"
)

type Biodata struct {
	Nama        string
	KodePeserta string
	Hobi        []string
}

var biodatas = Biodata{
	Nama:        "Rizca Shafira",
	KodePeserta: "529",
	Hobi:        []string{"Nonton", "Nulis"},
}

var PORT = ":3000"

func main() {
	http.HandleFunc("/rizca", greet)
	fmt.Println("Application is listening on port", PORT)
	err := http.ListenAndServe(PORT, nil)
	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	nama := biodatas.Nama
	kodePeserta := biodatas.KodePeserta
	strHobi := ""
	for _, v := range biodatas.Hobi {
		strHobi += v + " "
	}
	msg := "Nama: " + nama + "\n" +
		"Kode Peserta: " + kodePeserta + "\n" +
		"Hobi: " + strHobi
	fmt.Fprint(w, msg)
}
