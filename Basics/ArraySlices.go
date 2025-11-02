package basics

import "fmt"

func Arrays() {
	var simpleArray = [6]int64{1, 6, 5, 8}

	fmt.Println(simpleArray)

	// Slices
	var sliceArray = []int{2, 5, 5}
	var sliceWithMake = make([]int, 0, 5)

	// Append Value with loop

	for i := range 100 {
		fmt.Println("loop", cap(sliceArray), sliceArray)
		fmt.Println(sliceWithMake)
		sliceArray = append(sliceArray, i)
	}

	fmt.Println(cap(sliceArray), len(sliceArray), sliceArray, cap(sliceWithMake), len(sliceWithMake), sliceWithMake)
}