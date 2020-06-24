package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Starting Directory: ", wd) // /Users/renishb/Code/Ren/smallgohers/golang_systems

	if err = os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}

	wd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wd) // /
}
