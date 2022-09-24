/*
Konstata merupakan jenis variable yg nilainya tetap
ketika variabel dibuat harus langsung diberi nilai
*/

/*
Operator di Go ada 3 jenis : aritmatika, logika,
*/

/*
Operator aritmatika => + - * / % ++ --
Operator perbandingan/relasional => == != > < >= <=
Operator logical => && || !
*/

package main

import "fmt"

func main() {
	const g float32 = 9.8
	fmt.Printf("Konstanta gravitasi = %f\n", g)

	var aritmatika = (2 + 2) * 3
	fmt.Println(aritmatika)

	var relasional bool = "riz" == "Riz"
	fmt.Println(relasional)

	var wrong = false //logical
	fmt.Println(!wrong)
}
