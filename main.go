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

	fmt.Printf("%+v", jim)
	fmt.Printf("%+v", chris)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex@gmail.com"
	alex.contact.zipCode = 96000

	fmt.Printf("%+v", alex)

	fmt.Println(alex)

}
