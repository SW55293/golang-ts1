package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("/Users/steph/projects/golang-projects/tests"); err == nil {

		if os.IsNotExist(err) {
			fmt.Println("didnt work", err)
		} else {
			fmt.Println("did work")
		}
	}

	// fmt.Println("it ran")
}
