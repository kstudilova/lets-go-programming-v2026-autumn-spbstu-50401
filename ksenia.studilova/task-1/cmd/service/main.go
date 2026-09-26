package main

import "fmt"

func main() {
	var firstOperand int
	_, err1 := fmt.Scan(&firstOperand)

	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var secondOperand int
	_, err2 := fmt.Scan(&secondOperand)

	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	_, err3 := fmt.Scan(&operation)

	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}
}
