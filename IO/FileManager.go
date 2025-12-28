package io

import (
	"fmt"
	"os"
)

func ReadFile() {
	// os.Create("Ankan.txt")

	file, err := os.Create("Ankan.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes := []byte("Ankan Saha")

	n, err:= file.Write(bytes)

		if err != nil {
		panic(err)
	}

	fmt.Println(n)
}