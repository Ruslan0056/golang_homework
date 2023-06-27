package main

import (
	"OOP/Products"
	"OOP/VM"
	"fmt"
)

func main() {
	pr1, _ := Products.NewProduct(1, "Apple", "fruit", 50)
	pr2, _ := Products.NewProduct(2, "Water", "Water", 100)
	pr3, _ := Products.NewProduct(3, "Orange", "fruit", 80)
	pr4, _ := Products.NewBeverage(4, "Cola", "Beverage", 80, 500)
	fmt.Println(pr1, pr2, pr3, pr4)

	vm1 := VM.NewVendingMachine(2)
	vm2 := VM.NewVendingMachine(2)

	vm1.AddProduct(pr1)
	vm1.AddProduct(pr2)
	vm2.AddProduct(pr3)
	vm2.AddBeverage(pr4)

	fmt.Println(vm1, vm2)

}
