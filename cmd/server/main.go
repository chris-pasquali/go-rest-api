package main

import "fmt"

/*
responsible for instatiation and
startup of go application
*/
func Run() error {
	fmt.Println("starting up of our application")
	return nil
}

func main() {
	fmt.Println("Go REST API entry point")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
