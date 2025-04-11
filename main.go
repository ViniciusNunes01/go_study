package main

import "fmt"

func main() {

	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)

	sayHello(name)

}

func sayHello(stringName string) {
	fmt.Printf("Hello %s, from function!\n", stringName)
}
