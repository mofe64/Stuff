package main

import "fmt"

type contactInfo struct {
	zip   string
	state string
}
type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {

	mofe := person{
		firstname: "mofe",
		lastname:  "mofe",
		contactInfo: contactInfo{
			zip:   "zip",
			state: "state",
		},
	}

	//mofePointer := &mofe
	//mofePointer.updateName("test")
	mofe.updateName("test") //instead of creating a pointer and using it, we can pass the type directly and go converts our var to a pointer for us

	mofe.print()
}

func (p *person) updateName(newName string) {
	(*p).firstname = newName //we dereference pointer and then change the underlying value
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
