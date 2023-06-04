package main

import "fmt"

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up out application")
	return nil
}

func main() {
	fmt.Println("Go Rest App")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
