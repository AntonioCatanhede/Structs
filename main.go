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

	alexPointer := &alex // variable that stores the memory address of alex
	alexPointer.updateName("Alexei")

	fmt.Println(alex)

	matt := person{
		firstName: "Matt",
		lastName:  "Smeets",
		contact: contactInfo{
			email:   "matt@gmail.com",
			zipCode: 98000,
		},
	}
	mattPointer := &matt
	mattValue := *mattPointer
	newMattPointer := &mattValue

	fmt.Println(mattPointer)

	fmt.Println(mattValue)

	fmt.Println(newMattPointer)

	Name := &matt.firstName

	contato := &matt.contact

	fmt.Println(Name)
	fmt.Println(contato)

	a := 1

	fmt.Println("pointer: ", &a)
	fmt.Println("pointer: ", *(&a))
	fmt.Println("pointer: ", &(*(&a)))

}

func (pointerToPerson *person) updateName(newFirstName string) {

	(*pointerToPerson).firstName = newFirstName // *' poderia ser omitido aqui

}

func (p person) printPerson() {

	fmt.Printf("%+v \n", p)

}
