// belajar bagaimana cara membuat web server beserta routingnya pada bahasa Go
package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080" //nilai  port locahost yang mengarah pada localhost:8080.

func main() {
	http.HandleFunc("/", greet) //function HandleFunc yang berasal dari package http. Function HandleFunc digunakan untuk keperluan routing
	//HandleFunc menerima 2 parameter, pertama digunakan untuk mendefinisikan pathroutingnya, sedangkan parameter kedua menerima sebuah function dengan 2 parameter yaitu http.ResponseWriter dan pointer http.Request
	http.ListenAndServe(PORT, nil)
	//FunctionListenAndServe menerima 2 parameter yaitu keterangan portyang kita pakai, dan http.Handler yang merupakan sebuahinterface. Namun karena kita tidak menggunakan http.Handler,
	//maka kita cukup memberikan zero value dari tipe data interface yaitu nil.
}

func greet(w http.ResponseWriter, r *http.Request) {
	//http.ResponseWriter adalah sebuah interface dengan berbagai method yang digunakan untuk mengirim response bali kepada client yang mengirimkan request.
	//Kemudian http.Request adalah sebuah struct yang digunakan untuk mendapat berbagai macam data request seperti form value,headers, url parameter dan lain-lain.
	msg := "Hello World"
	fmt.Fprint(w, msg)
	//mengirim response berupa tulisan“Hello World” kepada client. Lalu untuk mengirim responsenya kita menggunakan function fmt.Fprint.
}
