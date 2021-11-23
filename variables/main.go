package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	
)

const prompt = "then press ENTER when ready"

func main() {

	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) +2
	var secondNumber = rand.Intn(8) +2
	var thirdNumber = rand.Intn(8) +2
	var answer = firstNumber * secondNumber - thirdNumber

	Print(firstNumber,secondNumber,thirdNumber, answer)
}

func Print(firstNumber,secondNumber,thirdNumber, answer int){

	reader:= bufio.NewReader(os.Stdin)
	
	fmt.Println("Guess the number Game")
	fmt.Println("_____________________")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)

	reader.ReadString('\n')

	fmt.Println("Multipy your number by,", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("then multiply your result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide by the number you originally thought of ", prompt)
	reader.ReadString('\n')

	fmt.Println("then subtract by", thirdNumber,prompt)
	reader.ReadString('\n')

		fmt.Println("the answer is ", answer)
}