/*
You are designing a billing system for a coffee shop.
The shop sells a basic coffee for $5.
Customers can customize their coffee with add-ons such as milk (+$1), sugar (+$0.5), and whipped cream (+$1.5).
The system should allow customers to choose any combination of add-ons, including multiple quantities of the same add-on (e.g., double sugar).
New add-ons should be easy to introduce in the future without modifying the core coffee class.
*/

package main

import (
	"log"
)

var (
	BASE_COFFEE_USD         = 5.0
	MILK_ADDON_USD          = 1.0
	SUGAR_ADDON_USD         = 0.5
	WHIPPED_CREAM_ADDON_USD = 1.5
)

/*
Coffee interface
*/
type Coffee interface {
	Cost() float64
	Description() string
}

/*
Base Coffee implementing Coffee interface
*/
type BaseCoffee struct {
	Desc string
}

func (d *BaseCoffee) Cost() float64 {
	return BASE_COFFEE_USD
}

func (d *BaseCoffee) Description() string {
	return d.Desc
}

/*
Base Decorator which does composition as well as implement the interface
*/

type AddOns struct {
	Coffee Coffee
	Desc   string
	CostUp float64
}

func (d *AddOns) Cost() float64 {
	return d.Coffee.Cost() + d.CostUp
}

func (d *AddOns) Description() string {
	return d.Coffee.Description() + ", " + d.Desc
}

/*
Concrete decorators
Extend this to add new add ons
*/
type Milk struct {
	AddOns
}

type Sugar struct {
	AddOns
}

type WhippedCream struct {
	AddOns
}

func main() {
	baseCoffee := &BaseCoffee{
		Desc: "Coffee",
	}

	milk := &Milk{
		AddOns{
			Coffee: baseCoffee,
			Desc:   "milk",
			CostUp: MILK_ADDON_USD},
	}
	log.Println("Desc: [", milk.Description(), "] Cost: [", milk.Cost(), "]")

	whippedCream := &WhippedCream{
		AddOns{
			Coffee: milk,
			Desc:   "whipped cream",
			CostUp: WHIPPED_CREAM_ADDON_USD},
	}

	log.Println("Desc: [", whippedCream.Description(), "] Cost: [", whippedCream.Cost(), "]")

	sugar := &Sugar{
		AddOns{
			Coffee: whippedCream,
			Desc:   "sugar",
			CostUp: SUGAR_ADDON_USD},
	}

	log.Println("Desc: [", sugar.Description(), "] Cost: [", sugar.Cost(), "]")

	sugar = &Sugar{
		AddOns{
			Coffee: sugar,
			Desc:   "sugar",
			CostUp: SUGAR_ADDON_USD},
	}

	log.Println("Desc: [", sugar.Description(), "] Cost: [", sugar.Cost(), "]")

}
