package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

//  to read user input 
	reader := bufio.NewReader(os.Stdin)

//	var whatToSay string can also be expressed as shown below
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

// ReadString helps to capture whatever the user types until he/she presses enter ('\n')  
// in order to creat a loop we use the FOR commmand 

for{
	fmt.Print("=> ")

	userInput, _ := reader.ReadString('\n')

// since the quit command already has a function in the program
// we need to write a specific command in our code to end the program
//(\r\n) represents the return key for windows WHILE (\n) is for mac and other OS
// -1 is the number of times to repeat the command, -1 mean repteat whenever you see this command.
	userInput = strings.Replace(userInput, "\r\n","",-1)

	if userInput == "quit" {	
		break
		}else{
			respose:= doctor.Response(userInput)
	
	fmt.Println(respose)
		}
	}

	
}


