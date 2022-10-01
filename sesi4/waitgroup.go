/*
waitgroup -> struct dr package sync yg digunakan untuk melakukan sinkronisasi thd go routine
cara menggunakan sleep utk goroutine bkn cara yg baik
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "mango", "durian", "rambutan"} //var slice buah2an
	var wg sync.WaitGroup                                      //var bernama wg dgn tipe data sync.Waitgroup / struct dr package sync
	for index, fruit := range fruits {
		wg.Add(1)                        //method add digunakan untuk menambah counter dari waitgroup. counter berguna untuk memberitahu waitgroup ttg jumlah goroutine yg hrs ditunggu. krn looping 4 kali, berarti ada 4 goroutine yg hrs ditunggu sblm function main menghentikan eksekusinya
		go printFruit(index, fruit, &wg) //func printFruit dijalankan sbg goroutine dgn awalan command "go"
	}
	wg.Wait() //untuk menunggu seluruh goroutine menyelesaikan prosesnya dan menahan function main hingga seluruh proses go routine selesai
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) { //func menerima 3 parameter
	fmt.Printf("index -> %d, fruit -> %s\n", index, fruit)
	wg.Done() //method untuk memberitahu waitgroup ttg goroutine yg telah menyelesaikan proses. waitgroup pd printFruit adalah sebuah pointer, agar waitgroup pd function main dan printFruit mengandung memori yg sama
}
