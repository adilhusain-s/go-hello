package main

import (
	"fmt"
	"runtime"
)

type Name struct {
	first string
	last  string
}

func (n *Name) addFirstName(firstName string) {
	n.first = firstName
}

func (n *Name) addLastName(lastName string) {
	n.last = lastName
}

func main() {

	var MyName = Name{first: "adil", last: "shaikh"}
	fmt.Println(MyName.first, MyName.last)
	fmt.Println(runtime.GOOS, runtime.GOARCH)

}
