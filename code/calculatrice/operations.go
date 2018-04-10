package main

type calc float64

func (c calc) addFive() calc {
	return c + 5
}

func (c calc) multipleByFive() calc {
	return c * 5
}
