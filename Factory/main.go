package Factory

import "fmt"

/*
	- With Factory Pattern, we can create an object but don't need to figure out what their concrete type
*/

const (
	CashMethodType = "Cash"
	OnlineMethodType = "Online"
	//OtherMethodType = "Others"
)
type PaymentMethod interface {
	Pay(amount float32) string
}

type CashMethod struct {
}

type OnlineMethod struct {
}

func (m CashMethod) Pay(amount float32) string {
	return fmt.Sprintf("Pay with Cash, amount: %f\n", amount)
}

func (m OnlineMethod) Pay(amount float32) string {
	return fmt.Sprintf("Pay with E-banking, amount: %f\n", amount)
}

func GetPaymentMethod(method string) PaymentMethod {
	switch method {
	case CashMethodType:
		return new(CashMethod)
	case OnlineMethodType:
		return new(OnlineMethod)
	default:
		return nil
	}
}

//usage
func UsageExample(){
	payment := GetPaymentMethod(CashMethodType)
	fmt.Print(payment.Pay(124))
	payment = GetPaymentMethod(OnlineMethodType)
	fmt.Print(payment.Pay(124))
}