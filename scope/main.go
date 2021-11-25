package main

import (
	"fmt"
	"chuksapp/packageone"
)

func main() {
	var one = "one of the best"

	fmt.Println(one)

	hone()

	newVar := packageone.PublicVar	

	fmt.Println("have a good day and", newVar)
}

func hone(){
	var one = "lets party"

	fmt.Println(one)
}