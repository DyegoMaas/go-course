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

func (p *person /*pointer to person*/) updateFistName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateFistName2(newFirstName string) {
	(*p).firstName = newFirstName // gets the value from the pointer
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
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

	jim.updateFistName2("little jimmy")
	jim.print()

	jimPointer := &jim
	fmt.Println("value:", *&*&*&*&*&*&jim)
	fmt.Println(jimPointer)

	jimPointer.updateFistName2("big jimmy")
	jim.print()

	jim.updateFistName("just jimmy")
	jim.print()
}
