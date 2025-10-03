package main

import "log"

/*
Food delivery app interface
*/
type FoodDeliveryApp interface {
	Which()
	Delivery()
	Payment()
}

/*
Delivery strategy
*/
type Delivery struct {
	Time TimeStrategy
	Mode ModeStrategy
}

/*
TimeStrategy interface
*/
type TimeStrategy interface {
	ETA()
}

/*
Concrete struct implementing Time strategy
*/
type TenMinuteDelivery struct{}

func (d *TenMinuteDelivery) ETA() {
	log.Println("This is a 10 minute delivery")
}

type SixtyMinuteDelivery struct{}

func (d *SixtyMinuteDelivery) ETA() {
	log.Println("This is a 60 minute delivery")
}

/*
ModeStrategy interface
*/
type ModeStrategy interface {
	Mode()
}

/*
Concrete struct implementing mode strategy
*/
type Drone struct{}

func (d *Drone) Mode() {
	log.Println("This is a drone delivery")
}

type Bike struct{}

func (d *Bike) Mode() {
	log.Println("This is a bike delivery")
}

/*
Payment strategy interface
*/
type PaymentStrategy interface {
	PaymentMode()
}

/*
Concrete structs implementing payment strategy
*/
type Card struct{}

func (d *Card) PaymentMode() {
	log.Println("Payment mode is card")
}

type UPI struct{}

func (d *UPI) PaymentMode() {
	log.Println("Payment mode is UPI")
}

type Cash struct{}

func (d *Cash) PaymentMode() {
	log.Println("Payment mode is Cash")
}

/*
Zomato logic
*/
type Zomato struct {
	DeliveryStrategies []Delivery // Zomato **has a** delivery strategy
	PaymentStrategy    []PaymentStrategy
}

func (d *Zomato) Which() {
	log.Println("This is Zomato")
}

func (d *Zomato) Delivery() {
	for _, v := range d.DeliveryStrategies {
		v.Time.ETA()
		v.Mode.Mode()
	}
}

func (d *Zomato) Payment() {
	for _, v := range d.PaymentStrategy {
		v.PaymentMode()
	}
}

/*
End of Zomato logic
*/

/*
Swiggy logic
*/
type Swiggy struct {
	DeliveryStrategies []Delivery
	PaymentStrategy    []PaymentStrategy
}

func (d *Swiggy) Which() {
	log.Println("This is Swiggy")
}

func (d *Swiggy) Delivery() {
	for _, v := range d.DeliveryStrategies {
		v.Time.ETA()
		v.Mode.Mode()
	}
}

func (d *Swiggy) Payment() {
	for _, v := range d.PaymentStrategy {
		v.PaymentMode()
	}
}

/*
End of Swiggy logic
*/

/*
Magicpin logic
*/
type Magicpin struct {
	DeliveryStrategies []Delivery
	PaymentStrategy    []PaymentStrategy
}

func (d *Magicpin) Which() {
	log.Println("This is Magicpin")
}

func (d *Magicpin) Delivery() {
	for _, v := range d.DeliveryStrategies {
		v.Time.ETA()
		v.Mode.Mode()
	}
}

func (d *Magicpin) Payment() {
	for _, v := range d.PaymentStrategy {
		v.PaymentMode()
	}
}

/*
End of Magicpin logic
*/

/*
Eatsure logic
*/
type Eatsure struct {
	DeliveryStrategies []Delivery
	PaymentStrategy    []PaymentStrategy
}

func (d *Eatsure) Which() {
	log.Println("This is Eatsure")
}

func (d *Eatsure) Delivery() {
	for _, v := range d.DeliveryStrategies {
		v.Time.ETA()
		v.Mode.Mode()
	}
}

func (d *Eatsure) Payment() {
	for _, v := range d.PaymentStrategy {
		v.PaymentMode()
	}
}

/*
End of Eatsure logic
*/

/*
Driver
*/
func main() {
	var foodDeliveryApp FoodDeliveryApp

	foodDeliveryApp = &Zomato{[]Delivery{{Time: &TenMinuteDelivery{}, Mode: &Drone{}}, {Time: &SixtyMinuteDelivery{}, Mode: &Bike{}}}, []PaymentStrategy{&Card{}, &UPI{}}}
	foodDeliveryApp.Which()
	foodDeliveryApp.Delivery()
	foodDeliveryApp.Payment()

	foodDeliveryApp = &Magicpin{[]Delivery{{Time: &SixtyMinuteDelivery{}, Mode: &Bike{}}}, []PaymentStrategy{&UPI{}}}
	foodDeliveryApp.Which()
	foodDeliveryApp.Delivery()
	foodDeliveryApp.Payment()
}
