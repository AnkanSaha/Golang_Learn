package basics

import (
	"fmt"
	"slices"
)

func Arrays() {
	var simpleArray = [6]int64{1, 6, 5, 8}

	fmt.Println(simpleArray)

	// Slices
	var sliceArray = []int{}
	var sliceWithMake = make([]int, 0, 5)

	// Append Value with loop
	for i := range 100 {
		// fmt.Println("loop", cap(sliceArray), sliceArray)
		fmt.Println(sliceWithMake)
		sliceArray = append(sliceArray, i)
	}

	// copy slice to destination slice
	copy(sliceWithMake, sliceArray)

	fmt.Println(slices.Equal(sliceArray, sliceWithMake), sliceArray, sliceWithMake)

	// fmt.Println(cap(sliceArray), len(sliceArray), sliceArray, cap(sliceWithMake), len(sliceWithMake), sliceWithMake)
}
