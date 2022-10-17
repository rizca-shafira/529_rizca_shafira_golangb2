go cross platform artinya bisa digunakan di OS yg beda2
open source artinya tersedia scr gratis dan bisa digunakan siapa saja
go itu cepat. cepat dikompile nya.
statically typed artinya tipe variabelnya ditentukan di awal sebelum di compile
punya concurency (bbrp task diproses bersama tanpa mempengaruhi hsl akhir) dan punya memory management / garbage collector.
didesain utk komputer multicore
sintaks mirip dgn c++ 
tdk mendukung class dan object

setauku, skrg goroot nya di C:/Users/USER. jd aku buat workspace disitu, trus di folder aku go mod init [nama_folder]
go mod adalah manajemen dependensi

go workspace terdiri dr 3 folder : src utk file go, pk utk hasil compile, bin utuk exe hasil build

cara run file go : go run [filename].go
kalo dlm 1 folder bnyk yg package main dan mau di run semua tinggal go run [filename1].go filename2.go
kalo menggunakan go run, file di compile. kalo menggunakan go build, file dibuild jd exe

utk mendownload package ke gopath, gunakan command go get

dlm 1 project minimal ada 1 file yg package main

fmt.Println("text1", "text2")
output : text1 text2\n

fmt.Print("x", "y")
output : xy

//komentar
/*
komentar juga
*/
