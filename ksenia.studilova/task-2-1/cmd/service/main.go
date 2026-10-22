package main

import "fmt"

func main() {
	var departmentCount int
	fmt.Scan(&departmentCount)

	for i := 0; i < departmentCount; i++ {
		minTemp := 15
		maxTemp := 30

		var employeeCount int
		fmt.Scan(&employeeCount)

		for j := 0; j < employeeCount; j++ {
			var operation string
			var temp int
			fmt.Scan(&operation, &temp)

			if operation == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			} else if operation == "<=" {
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
