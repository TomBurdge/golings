// structs2
// Make me compile!
package main

import "fmt"

type Person struct {
	// don't just create the phone field here. embed a new struct
	name  string
	age   int
	phone ContactDetails
}

type ContactDetails struct {
	number string
}

func main() {
	contactDetails := ContactDetails{"+44568"}
	person := Person{name: "John", age: 32, phone: contactDetails}
	fmt.Printf("%s is %d kyears old and his phone is %s\n", person.name, person.age, person.phone)
}
