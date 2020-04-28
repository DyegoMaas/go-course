package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	guilherme := person{"Guilherme", "Maas"} // positional and ugly
	dyego := person{firstName: "Dyego", lastName: "Maas"}
	var tinti person
	tinti.firstName = "Angelo"
	tinti.lastName = "Tinti"

	fmt.Println(guilherme)
	fmt.Println(dyego)
	fmt.Printf("%+v", tinti)
}
