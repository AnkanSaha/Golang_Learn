package basics

import (
	"fmt"
	"time"
)

type Details struct {
	name      string
	age       float64
	createdAt time.Time
}

func Makestruct(age float64) Details {
	mydetisl := Details{
		name: "Ankan",
		age:  age,
	}

	var newStruct = &mydetisl
	newStruct.name = "Ankan Saha"

	fmt.Println("new", newStruct, mydetisl)

	return mydetisl
}

// struct with oops

type Oops struct {
	name  string
	age   int32
	class string
}

// rechiver type ( like method in js opps taking Oops pointer)

func (modifyer *Oops) ChnageStruct(name string, age int16, class string) {
	modifyer.age = int32(age)
	modifyer.name = string(name)
	modifyer.class = string(class)
}

func Makestruct2() {
	var myStruct = Oops {
		name: "Ankan",
	}

	fmt.Println("newStruct2", myStruct)

	myStruct.ChnageStruct("Ankan Saha", 22, "Six Semester")
	fmt.Println("newStruct2", myStruct)
}

// OPPS in Go

type NewOpps struct {
	name string
	age  int
}

// methods
func (O *NewOpps) ChnageAge(age int) {
	O.age = age
}

func NewOppss(name string, age int) *NewOpps {
	var newOpesUser = NewOpps{
		name: name,
		age:  age,
	}
	return &newOpesUser
}


// Unnamed Struct

func ShowunnammedStruct (){
	language := struct {
		name string
	} {"Ankan"}

	fmt.Println("Struct Unnamed", language)
}


// struct embedding

type balance struct {
	avl int
}
type user struct {
	name string;
	age int;
	balance
}


func MakeStructure () {
	userDetails := user {
		name: "Ankan",
		age: 22,
		balance: balance {
			avl: 222,
		},
	}

	fmt.Println("struct embedding", userDetails)
}