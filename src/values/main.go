package main

import (
	"fmt"
	"strings"
)

func BuildValues() []string {
	return []string{
		"go" + "lang",
		"1+1=", fmt.Sprintf("%d", 1+1),
		"7.0/3.0 =", fmt.Sprintf("%f", 7.0/3.0),
		"true && false =", fmt.Sprint(true && false),
		"true || false =", fmt.Sprint(true || false),
		"!true =", fmt.Sprint(!true),
	}
}

func GetValues() string {
	return "\nexamples/values -------->\n" + strings.Join(BuildValues(), "\n")
}

func main() {
	fmt.Println(GetValues())
}