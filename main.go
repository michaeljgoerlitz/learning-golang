package main

import (
	"fmt"

	"github.com/mjgoerlitz-new/go-demo-1/helloworld"
	"github.com/mjgoerlitz-new/go-demo-1/variables"
)

func main() {
	fmt.Println(helloworld.HelloWorld())
	variables.SetVariables()
}
