package basics;

import "fmt"

func ForLoop() {

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration for loop:", i)
	}

	i := 1

	// While loop using for
	for i <= 5 {
		fmt.Println("Iteration while loop:", i)
		i++
	}


	// Infinite loop with break
	j := 1
	for {
		if j > 5 {
			break
		} else if j == 3 {
			j++
			continue
		}
		fmt.Println("Iteration infinite loop:", j)
		j++
	}

	// Range loop over slice like forEach in JavaScript
	numbers := [5]int {3,5,4,5,5}

	for index, value := range numbers {
		fmt.Println(index, value)

	}
}