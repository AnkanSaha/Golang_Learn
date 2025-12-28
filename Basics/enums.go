// Enums in Golang
package basics

import "fmt"

// custom type
type myType int;

const (
	Create myType = iota
	In_Progress
	Shipped
	Devivered
)

// this is String enum
const (
	RECHIVED string = "RECHIVED"
	In_Progresss string = "InProgress"
)


func CreateOrder (status myType){
	fmt.Println("order status", status)
}
func CreateStringOrder (status string){
	fmt.Println("order status", status)
}

func CreateEnum (){
	CreateStringOrder(In_Progresss)
	CreateOrder(In_Progress)
}