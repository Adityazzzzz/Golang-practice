package main

import "fmt"



type paymenter interface{
	pay(amount float32)
}



type payment struct{
	gateway paymenter
}
// method to payment struct
func (p payment) handlePayments(amount float32) {
/*
	//1
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)

	//2
	stripePaymentGw := stripe{}
	stripePaymentGw.pay(amount)
*/

	// instead we can create a single way
	p.gateway.pay(amount)
}



type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay",amount)
}



type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("making payment using stripe",amount)
}



func main() {
	newpayment := payment{}
	newpayment.handlePayments(100.00)
}