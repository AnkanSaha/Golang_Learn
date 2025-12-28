package basics

import "fmt"

type razorpay struct{}

func (razorpay razorpay) pay(amput int32) {
	fmt.Println("Paying ", amput, "through razorpay")
}

type stripe struct {}

func (stripe stripe) pay(amput int32) {
	fmt.Println("Paying ", amput, "through stripe")
}

type paypal struct {}

func (paypal paypal) pay (amount int32) {
	fmt.Println("Paying ", amount, "through paypal")

}


// interface
type gateway interface {
	pay (amount int32)
}

type PaymentPG struct {
	gateway gateway
}

func Payment (){
	// stripe:= stripe{}
	// razorpay:= razorpay{}
	paypal:= paypal {}
	paymentPG := PaymentPG {
		gateway: paypal,
	}

	paymentPG.gateway.pay(22)
}