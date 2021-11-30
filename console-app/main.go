package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil{
		log.Fatal(err)
	}

	defer func(){
		_=keyboard.Close()
	}()
// we used the make map command to turn int to string on the console
	coffees := make(map[int]string)
	coffees[1] = "Cappachino"
	coffees[2] = "Mocha"
	coffees[3] = "Latte"
	coffees[4] = "Americano"
	coffees[5] = "Macchiato"
	coffees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("____")
	fmt.Println("1 - Cappachino")
	fmt.Println("2 - Mocha")
	fmt.Println("3 - Latte")
	fmt.Println("4 - Americano")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Expresso")
	fmt.Println("Q - Quit the menu")






	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q'{
			break
		}
		
		i, _ := strconv.Atoi(string(char))
		t:= fmt.Sprintf("you chose %s", coffees[i] )

		fmt.Println(t)
	
		
	}
	fmt.Println("program is ending.....")

}