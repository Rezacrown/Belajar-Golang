package main

import (
	"fmt"
)

var name string = "Reza"

func main() {

	if name == "Reza" {
		fmt.Println("name: reza")
	} else if name == "ary" {
		fmt.Println("name: ary")
	} else {
		fmt.Println("nama: default else")
	}

	fmt.Println(switching("switchign: "))

}

func switching(name string) string {

	// var test = TestAja()

	return name
}
