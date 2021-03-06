package main

type address struct {
	street     string
	department int
}

type student struct {
	firstname string
	lastname  string
	age       int
	address
}

func getStudent(firstname string) student {
	return student{
		firstname: firstname,
		lastname:  "Ahlen",
		age:       20,
		address: address{
			street:     "Any street",
			department: 32,
		},
	}
}

func (s student) toString() string {
	return s.firstname + " " + s.lastname + ", living: " + s.address.street
}

// But (pointer+type) means a pointer of that type
func (pointerToStudent *student) updateName(firstname string) {
	// Change in the original memory and not a copy of the object
	// *pointer means: Give me the value this memory address is pointing at
	(*pointerToStudent).firstname = firstname
}
