/*
channel -> mekanisme goroutine saling berkomunikasi dgn goroutine lainnya
berkomunikasi dgn pengiriman dan pertukaran data antara goroutine satu dgn goroutine lainnya
utk membuat perlu built-in function make
var c mempunyai tipe data channel of string/channel yg memiliki tipe data string
chan merupakan keyword utk membuat channel
*/

/*
butuh operator dr channel agar channel dpt dijadikan sbg alat komunikasi antara goroutine dgn yg lain
tanda <- merupakan operator channel utk proses pengiriman data dr goroutine satu dgn lainnya
c <- mengirimkan data ke channel c
result := <-c adalah menerima data dr channel c dan menyimpan di var result
pengiriman dan penerimaan data dr channel bersifat synchronous
*/

/*
krn function main mrpkn goroutine, maka syntax dibawah mrpkn komunikasi antara goroutine main dgn goroutine introduce melalui channel
ketika dijalankan, hasilnya dpt berbeda2 urutan print nya
krn antara masing2 goroutine tdk saling menunggu shg bisa digunakan utk concurrency
*/

/*
ketika menggunakan channel sbg parameter function, kita dpt menentukan apakah channel hny menerima, hny mengirim, atau menerima dan mengirim data
channel direction bersifat optional dan digunakan utk kepentingan type safety
parameter syntax | detail
c chan string | parameter c dpt mengirim dan menerima data
c chan <- string | parameter c hny dpt digunakan utk mengirim data
c <- chan string | parameter c hny dpt digunakan utk menerima data
*/

/*
channel bersifat unbuffered atau tdk dibuffer di memori. channel yg kita gunakan di bwh merupakan unbuffered channel dmn proses pengiriman dan penerimaan bersifat synchronous
lain halnya dgn unbuffered channel yg bisa ditentukan kapasitas channelnya selama jmlh data yg dikirim melalui unbuffered channel tdk melebihi kapasitasnya, maka proses bersifat asynchronous
cara membuat unbuffered channel dgn kapasitas buffer :

	func main(){
		c:= make(chan int,3)
	}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)  //unbuffered channel
	go func(c chan int) { //go routine func anonymous
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)
	/*
		tulisan berupa func goroutine after sending data into the channel seharusnya ditampilkan setelah tulisan func goroutine starts sending data into the channel, tetapi faktanya tulisan func goroutine after sending data into the channel ditampilkan diakhir
		Hal ini terjadi karena syntax apapun yang berada dibawah proses pengiriman data melalui unbuffered channel akan tertahan hingga datanya diterima oleh Goroutine lainnya
		Jika kita perhatikan pada gambar di halaman sebelumnya, proses penerimaan data pada function main tertahan selama 2 detik sehingga. Karena syntax untuk menampilkan tulisan func goroutine after sending data into the channel berada dibawah proses pengiriman data, maka dari itu tulisan tersebut ditampilkan setelah datanya di terima pada function main.
	*/

	c2 := make(chan int, 3) //buffered channel
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c2)
	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c2 { //v=<-c2
		fmt.Println("main goroutine received value from channel:", v)
	}

	/*
		func goroutine #1 after sending data into the channel dan tulisan lainnya yg mengandung kata after langsung ditampilkan setelah proses pengiriman data. Hal ini krn digunakan channel dgn buffer kapasitas 3.
		Proses pengiriman data terhenti ketika pengiriman  data sdh melebihi kapasitas dari kapasitas buffer yg kita berikan sebanyak 3 buffer
		jika pengiriman data sdh melebihi kapasitas, proses penerimaan data akan dieksekusi hingga paling tdk sdh ada data yg diterima.
		kemudian setelah proses penerimaan data tlh selesai dan jika masih ada data yg dikirimkan, maka proses pengiriman data dimulai lagi secara asynchronous
	*/

	/////////////////////////////////////

	c := make(chan string)
	//mengirim data ke channel
	//c <- value
	//menerima dari dr channel
	//result := <-c

	//goroutine memanggil function introduce sbnyk 3 kali dgn keyword go
	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	msg1 := <-c //func main menerima data dr function introduce dan disimpan ke var msg1, msg2, msg3
	fmt.Println(msg1)
	msg2 := <-c
	fmt.Println(msg2)
	msg3 := <-c
	fmt.Println(msg3)

	students := []string{"Rizca", "Shafira", "Salsabila"}
	for _, v := range students {
		go func(student string) { //func introduce diganti func anonymous di dlm range loop
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s anonymous", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}
	close(c) //channel tdk digunakan utk berkomunikasi lg

	//Select -> fitur go yang memungkinkan kita utk dpt menggunakan lbh dari satu channel utk proses komunikasi antara goroutine satu dgn yg lainnya
	//contoh : dibuat channel c3 dan c4 dan 2 goroutine menggunakan function anonymous
	c3 := make(chan string)
	c4 := make(chan string)
	_, _ = c3, c4

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c4 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ { //melakukan looping sebanyak channel yg dibuat (2)
		select {
		case msg1 := <-c3: //case terpenuhi saat ada penerimaan data dr channel c3 yg kemudian ditampung oleh var msg1
			fmt.Println("Received", msg1)
		case msg2 := <-c4: //case terpenuhi saat ada penerimaan dari channel c4 yg kemudian ditampung di msg2
			fmt.Println("Received", msg2)
		}
	}
}

// func introduce membuat pesan dgn kalimat yg dibuat function fmt.Sprintf
func introduce(student string, c chan<- string) { //direction hny utk mengirim data
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result //mengirimkan nilai yg dikandung var result ke channel c
}

// func menerima data melalui channel c
func print(c <-chan string) { //direction channel hny utk menerima data
	fmt.Println(<-c) //ketika tdk ingin menampung sebuah kiriman data dr channel, operator ditulis langsung
}
