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
