package main

import (
	"fmt"

	"github.com/mjgoerlitz-new/go-demo-1/functions"
	"github.com/mjgoerlitz-new/go-demo-1/helloworld"
	"github.com/mjgoerlitz-new/go-demo-1/variables"
)

func main() {
	fmt.Println(helloworld.HelloWorld())
	variables.SetVariables()
	fmt.Println("\n", functions.HelloWorld())
	fmt.Println(functions.Subtract(2, 1))
	fmt.Println(functions.Addition(1, 2, "The sum is "))
	functions.TestIncrements()
}
