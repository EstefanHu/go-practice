package main

import (
	"fmt"
	"strings"
)

func main() {
	integerStuff()
}

func integerStuff() {
	var unsignedBoi uint8 = 200
	fmt.Println(unsignedBoi)
	var temp float64 = 12.75
	var tmpInt int = 10
	fmt.Printf("Integer to floag: %.2f\n", 10.0)
	fmt.Printf("%T\n", tmpInt)
	fmt.Printf("%T\n", temp)
}

func stringStuff() {
	str3 := "hello, world"
	str4 := "Hello, world"

	fmt.Println(strings.EqualFold(str3, str4))
	fmt.Println(strings.Replace(str4, "world", "Go", 1))
}

func greetMy() {
	fmt.Println("Hello Essentials")
	var name string = "Estefan"
	fmt.Printf("Hello %s", name)
}
