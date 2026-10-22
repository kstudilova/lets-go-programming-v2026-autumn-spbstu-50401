package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Input reading error:", err)
		return
	}

	for i := 0; i < departmentCount; i++ {
		minTemp := 15
		maxTemp := 30

		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Input reading error:", err)
			return
		}

		for j := 0; j < employeeCount; j++ {
			var operation string
			var temp int

			_, err := fmt.Scan(&operation, &temp)
			if err != nil {
				fmt.Println("Input reading error:", err)
				return
			}

			switch operation {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			default:
				fmt.Println("Unknown operation:", operation)
				return
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
