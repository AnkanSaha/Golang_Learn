package main

import (
	"fmt"
	"learn_go/Basics"
)

func main() {

	fmt.Println("Hello, World!")
	basics.Variables()
	basics.Constants()
	basics.ForLoop()
	basics.Arrays()
	basics.Maps()
	basics.ExampleRange()
	fmt.Print("A",basics.Makestruct(22));
	basics.FuncWithanyType("Ankan", "Saha")
	basics.FunctionWithVariadicParams(1, 2, 3, "Ankan")
	value1, value2 := basics.FunctionWithMultiReturn(22, 55)

	if value2 != nil {
		fmt.Println(*value2, value1)
	}

	basics.Makestruct2()

	users := basics.NewOppss("Rohut", 22)

	if users != nil {
		users.ChangeAge(2)
	}

	basics.ShowunnammedStruct()
	basics.MakeStructure()
	basics.Payment()
}