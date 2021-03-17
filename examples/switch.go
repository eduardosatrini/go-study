package main

import "fmt"

func DaysOfTheWeek(n int) string {
	switch n {
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	case 7:
		return "monday"
	default:
		return "Invalid number!"
	}
}

func main() {
	fmt.Println(DaysOfTheWeek(2))
}
