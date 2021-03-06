package main

import (
	"fmt"

	handlers "github.com/niroopreddym/caymanislandsquerybuilder/Handlers"
)

//Driver drives the biz layer
type Driver struct {
	Handler handlers.IBizHandler
}

func main() {
	inputFileName := "query2.json"
	driver := Driver{
		Handler: handlers.NewHandler(inputFileName),
	}

	// output1 := driver.Handler.Assignment1()
	// fmt.Println(output1)

	output2 := driver.Handler.Assignment2()
	fmt.Println(output2)
	// ----------------------------------------------------------
}
