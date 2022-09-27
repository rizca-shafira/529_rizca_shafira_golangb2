package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Airell", 23)
	greet2("Airell", "Jl Digital")

	var names = []string{"Airell", "Jordan"}
	var printMsg = greet3("Heii", names)
	fmt.Println(printMsg)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	area, circumference = calculate2(diameter)
	fmt.Println("Area2:", area)
	fmt.Println("Circumference2:", circumference)

	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")
	fmt.Printf("%v", studentLists)
	fmt.Println(" ")

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("Result:", result)

	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

// jika semua parameter tipe datanya sama, ga perlu dituliskan satu persatu
func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

// jika function mereturn sebuah nilai, perlu ditulis tipe data hasil return nya
// function join untuk menggabungkan nilai dgn tipe data string yg ada di slice atau array
func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

// returning multiple values
func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

// predefined return value
// var area dan circumference di dalam function bukan var baru, tapi var yg digunakan sbg nilai return
func calculate2(d float64) (area float64, circumference float64) { //bedanya var sdh didefinisikan di main
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return //bedanya return
}

// variadic function 1 -> dpt menerima argumen yg jumlahnya tak terbatas
func print(names ...string) []map[string]string {
	var result []map[string]string
	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// variadic function 2 -> menggunakan tipe data slice sbg parameter dari sebuah function variadic
// perlu menggunakan tanda ellipsis / titik tiga
func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

// variadic function 3
// menggabungkan parameter biasa dgn parameter variadic pd sebuah function variadic
// parameter variadic perlu diletakkan di posisi akhir parameter
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")
	fmt.Println("Hello there! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
