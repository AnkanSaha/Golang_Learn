package basics

import "fmt"

func FunctionSExample() {
}

func FunctionWithParams(a int, b int) int {
	return a + b
}

func FunctionWithSameType(a, b int) int {
	return a + b
}

func FunctionWithMultipleReturn(a int, b int) (string, string) {
	return "Ankan", "Saha"
}

func FuncWithanyType(a any, b any) {
	fmt.Println(a, b)
}

func FunctionWithVariadicParams(nums ...any) {
	for _, value := range nums {
		fmt.Println(value)
	}
}

func FunctionWithMultiReturn (anyNumber ...int64) (int64, *int64) {
	if len(anyNumber) < 2 {
		return 0, nil
	}
	return anyNumber[0], &anyNumber[1]
}
