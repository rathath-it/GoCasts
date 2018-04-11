package main

import (
	"strconv"
)

type mynumber int

func (myn mynumber) isEven() bool {
	return (myn%2 == 0)
}

func (myn mynumber) isOdd() bool {
	return !myn.isEven()
}

func (myn mynumber) feedback() string {
	message := strconv.Itoa(int(myn)) + " is odd"
	if myn.isEven() {
		message = strconv.Itoa(int(myn)) + " is even"
	}
	return message

}
