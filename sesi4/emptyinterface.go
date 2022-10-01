/*
empty interface -> tipe data yg menerima segala data pd bhs go. maka, dpt diberikan nilai dgn tipe data yg berbeda
var randomValues memiliki tipe data empty interface. kemudian diberi nilai string, int, bool, dan slice of int. menunjukkan bisa diberi nilai apapun
*/
package main

func main() {
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}

	var v interface{}
	v = 20
	//v = v * 9
	//akan error krn  kita hanya bisa melakukan perkalian thd tipe data int yg asli, bukan empty interface yg diberi tipe data int
	//untuk menanggulangi bisa melakukan type assertion dgn cara yg sama dgn type assertion pd tipe data interface

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	//ketika map dideklarasikan dan value valuenya diberikan tipe data empty interface, maka map tsb dpt memiliki value dgn tipe data yg berbeda2
	//sama halnya dgn slice atau array

	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}
	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
