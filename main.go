package main

import "fmt"

func main() {
	mybill := newBill("Rizzy bill")

	mybill.addItem("juuusss", 7.50)
	mybill.addItem("juuuss", 7.50)
	mybill.addItem("juuus", 7.50)
	mybill.addItem("juuu", 7.50)

	mybill.updateTip(20)

	fmt.Println(mybill.format())
}
