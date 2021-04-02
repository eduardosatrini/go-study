package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("MY OS.")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X", os)
	case "linux":
		fmt.Println("GNU/LINUX", os)
	default:
		fmt.Println("IDK")
		fmt.Printf("%s", os)
	}

}
