/*
Data string url bisa dikonversi kedalam bentuk url.URL
Dengan menggunakan tipe tersebut akan ada banyak informasiyang bisa kita manfaatkan, diantaranya adalah jenis protokolyang digunakan, path yang diakses, query, dan lainnya.
Function url.Parse digunakan untuk parsing string ke bentuk url.Mengembalikan 2 data, variabel objek bertipe url.URL dan error(jika ada)
query yang ada pada url akan otomatis diparsing juga,menjadi bentuk map[string][]string, dengan key adalah namaelemen query, dan value array string yang berisikan value elemenquery.
*/

package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://developer.com:80/hello?name=airell&age=23"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("urlL %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name:%s,age:%s\n", name, age)
}
