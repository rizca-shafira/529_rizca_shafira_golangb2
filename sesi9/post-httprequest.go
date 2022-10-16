// post request
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//siapkan data request body yg akan dikirimkan
	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "jordan",
		"userId": 1,
	}
	requestJson, err := json.Marshal(data) //ubah menjadi json

	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}
	//function new request menerima 3 parameter method dari request yg ingin dibuat, url dr requestnya, dan data yg ingin dikirimkan
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	//bytes new buffer mengubah tipe data yg ingin dikirim menjadi interface io.reader karena parameter ketiga function http.NewRequest menerima nilai dgn tipe data interface io.Reader dan function bytes.NewBuffer mereturn sebuah nilai tipe data pointer dari struct bytes.Buffer yg telah mengimplementasikan salah satu method dari interface io.Reader
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
