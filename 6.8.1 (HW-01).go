package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIntResult(num1, num2 float64, operator string) float64 { // Basic operation function
	var value float64
	switch operator {
	case "+":
		value = num1 + num2
	case "-":
		value = num1 - num2
	case "*":
		value = num1 * num2
	case "/":
		value = num1 / num2
	default:
		panic("Illegal operator")
	}
	return value
}
func main() {
	reader := bufio.NewReader(os.Stdin) //read from the console, based on OS standard input
	fmt.Println(`Enter mathematical operation in following format: "a + b, a - b, a * b, a / b"'`)
	text, _ := reader.ReadString('\n') 
	text = strings.TrimSpace(text)
	operationSlice := strings.Split(text, " ") //Making a slice, uses "spaces" as separator
	if len(operationSlice) != 3 {              //Checks if you entered correct amount of arguments. 2 numbers and 1 operator
		panic("Incorrect amount of arguments.")
	}
	num1 := operationSlice[0] //Assigning numbers and operators after slicing the text.
	num2 := operationSlice[2]
	operator := operationSlice[1]
	convInt1, _ := strconv.ParseFloat(num1, 64) 
	convInt2, _ := strconv.ParseFloat(num2, 64) 
	if convInt2 == 0 && operator == "/" {
		panic("Error division")
	}
	fmt.Println("Result: ", getIntResult(convInt1, convInt2, operator))
}
