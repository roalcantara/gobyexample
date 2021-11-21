package main

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println(BuildHelloWorld())
}

func BuildHelloWorld() string {
	return "[BuildHelloWorld] hello world"
}

func GetHelloWorld() string {
	return "\n[GetHelloWorld] examples/hello-world -------->\n" + BuildHelloWorld()
}

func main() {
	HelloWorld()
}
