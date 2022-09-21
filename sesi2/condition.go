package main

import "fmt"

//Condition
func main() {
	var age = 23

	if age > 18 {
		fmt.Println("Sudah punya KTP")
	} else {
		fmt.Println("Belum punya KTP")
	}
	//else if condition

	if age > 18 {
		fmt.Println("Umur ", age, "> 18")
	} else if age == 18 {
		fmt.Println("Umur ", age, "= 18")
	} else {
		fmt.Println("Umur ", age, "< 18")
	}
	//switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//switch with relational operator & fallthrough
	fmt.Print("Write ", i, " as ")
	switch {
	case (i > 5):
		fmt.Println("lebih dari 5")
	case (i == 5):
		fmt.Println("sama dengan 5")
	default:
		fmt.Println("kurang dari 5")
	}
	/*bisa juga nested conditional, jadi di dalam codition ada condition lain*/
}
