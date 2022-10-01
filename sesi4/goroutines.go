//Concurrency & Go Routines
//concurrency is about dealing with lots of things at once
//parallelism is a bout doing lots of things at once

/*
Concurrency -> mengeksekusi sebuah proses scr independen atau berhadapan dgn lbh dari satu tugas dlm waktu yg sama
bedanya parallelism pekerjaan diselesaikan bersamaan dari awal hingga akhir. sdngkan pd concurency, tdk tahu pekerjaan mana yg akan diselesaikan terlebih dahulu
*/

/*
goroutines -> thread ringan pd bhs go utk concurency. ukurannya lbh ringan dibanding thread lain
1 goroutine hny butuh 2kb memori, 1 thread biasa bisa hbs 1-2 mb memori
*/

/*
untuk membuat goroutine, hrs terlebih dahulu membuat sebuah function. function dipanggil dgn keyword go sblm memanggil fungsi. function bernama goroutine dipanggil dgn cara menuliskan go dulu. scr otomatis function goroutine menjadi sebuah Goroutine
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go goroutine()

	fmt.Println("main execution started")
	secondProcess(8)
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("main execution ended")
}

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

/*
goroutine (asynchronous process #1)
terdapat 2 function : firstProcess dan secondProcess. kedua function digunakan utk menampilkan angka dari 1 hingga bilangan yg ditentukan dr parameter yg diterima dgn melakukan looping
firstProcess dijadikan goroutine krn dipanggil dgn menggunakan keyword go.
func NumGoroutine dari package runtime utk mengetahui jumlah Goroutine yg sedang berjalan
*/

/*
jika diperhatikan, function firstProcess tdk menampilkan hasilnya. krn goroutin bekerja scr asynchronous dan satu goroutine tdk saling tunggu dgn goroutine lainnya
kemudian ada 2 goroutine yg berjalan pdhl hanya menjalankan 1 function yg dijadikan sbg sebuah goroutine. krn function main jg merupakan goroutine shg function main tdk akan menunggu goroutine lainnya selesai tereksekusi.
hal ini yg menyebabkan function firstProcess tdk menampilkan hasil walaupun sebetulnya function tsb telah tereksekusi
*/

/*
goroutine butuh waktu yg sedikit lbh lama untuk mulai dibandingkan function biasa. function Sleep digunakan untuk menahan function main agar tdk langsung menyelesaikan eksekusinya.
function sleep berasal dr package time
implementasi ada di file goroutines-sleep
*/
