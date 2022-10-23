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
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)
	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening to port 3000")
	err := http.ListenAndServe(":3000", mux)

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
		"Hobi: " + strHobi + "\n"
	fmt.Fprint(w, msg)
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		msg := "Hello middleware" + "\n"
		fmt.Fprint(w, msg)
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		nama := biodatas.Nama
		kodePeserta := biodatas.KodePeserta
		strHobi := ""
		for _, v := range biodatas.Hobi {
			strHobi += v + " "
		}
		msg := "Nama: " + nama + "\n" +
			"Kode Peserta: " + kodePeserta + "\n" +
			"Hobi: " + strHobi + "\n"
		fmt.Fprint(w, msg)
		next.ServeHTTP(w, r)
	})
}
