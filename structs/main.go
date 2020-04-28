package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jaw",
		contact: contactInfo{
			email:   "jimjaw@gmail.com",
			zipCode: 89000,
		},
	}

	fmt.Printf("%+v", jim)
}
