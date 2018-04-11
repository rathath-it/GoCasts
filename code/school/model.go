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

func (s student) updateName(firstname string) {
	s.firstname = firstname
}
