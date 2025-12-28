package basics

import (
	"fmt"
	"time"
)

type Details struct {
	name string;
	age float64;
	createdAt time.Time
}

func Makestruct (age float64) Details {
	mydetisl := Details {
		name: "Ankan",
		age: age,
	}

	var newStruct = &mydetisl
	newStruct.name = "Ankan Saha"
	
	fmt.Println("new", newStruct, mydetisl)

	return  mydetisl
}