package main

import "fmt"

// struct act as a blueprint for any bill
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills--takes name and returns bill
func newBill(name string) bill {
	// initial bill object
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// formating the bill using --RECEIVER Functions i.e (b bill)variable b of type bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		// -25 makes the name to occupy 25 letter spaces to the right
		fs += fmt.Sprintf("%-25v ...ksh %v \n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ....ksh %0.2f \n", "tip: ", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ....ksh %0.2f", "total:", total+b.tip)

	return fs
}

// updating tip--when calling a method where we are updting a value..we pass in a pointer
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
