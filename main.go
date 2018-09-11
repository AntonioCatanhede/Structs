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

type person2 struct {
	firstName   string
	lastName    string
	contactInfo // nao ha dubiedade na declaracao
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	chris := person2{
		firstName: "Chris",
		lastName:  "Templeton",
		contactInfo: contactInfo{
			email:   "chris@gmail.com",
			zipCode: 95000,
		},
	}

	fmt.Printf("%+v \n", jim)
	fmt.Printf("%+v \n", chris)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex@gmail.com"
	alex.contact.zipCode = 96000

	alex.printPerson()
	alex.updateName("Alexei")

	fmt.Println(alex)

}

func (p person) updateName(newFirstName string) {

	p.firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Printf("%+v \n", p)

}
