package main

import "fmt"

func main() {
	var departmentCount int
	fmt.Scan(&departmentCount)

	for i := 0; i < departmentCount; i++ {
		var employeeCount int
		fmt.Scan(&employeeCount)

		for j := 0; j < employeeCount; j++ {
			var operation string
			var temp int
			fmt.Scan(&operation, &temp)
		}
	}
}
