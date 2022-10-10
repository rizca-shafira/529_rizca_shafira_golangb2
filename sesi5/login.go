package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer Defer()

	var username string
	fmt.Println("Username:")
	fmt.Scanln(&username)
	fmt.Println("Pass:")
	pass, _ := gopass.GetPasswdMasked()

	if valid, err := validate(username, pass); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func Defer() {
	if r := recover(); r != nil {
		fmt.Println("Error occured.", r)
	} else {
		fmt.Println("No error occured.")
	}
}

func validate(username string, password []byte) (string, error) {
	p := string(password[:])
	if username != "rizca" || p != "h4h4h4" {
		return "", errors.New("Wrong username/password!")
	}

	return "You're logged in!", nil
}
