package main

import "fmt"

// start with monday because I want.
func DaysOfTheWeek(n int) string {
	switch {
	case n == 1:
		return "monday"
	case n == 2:
		return "tuesday"
	case n == 3:
		return "wednesday"
	case n == 4:
		return "thuesday"
	case n == 5:
		return "friday"
	case n == 6:
		return "saturday"
	case n == 7:
		return "sunday"
	default:
		return "invalid input"
	}
}

func main() {
	fmt.Println(DaysOfTheWeek(3))
}
