package main

import (
	"fmt"
)

var version = "dev"

func TellHello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(version)
	fmt.Println(TellHello("dhana"))
}
