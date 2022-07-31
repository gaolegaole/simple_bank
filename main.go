package main

import "fmt"

type Amount struct {
	Username string
	Password string
}

func main() {
	amount := new(Amount)
	amount.Password = "Password"
	amount.Username = "Username"
	fmt.Printf("%#v", amount)
	fmt.Printf("%v", amount)
}
