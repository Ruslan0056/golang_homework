package VM

import (
	"OOP/Products"
	"fmt"
)

type VendingMachine struct {
	vmName       string
	capacity     int
	productsList []Products.Product
	beverageList []Products.Beverage
}

func NewVendingMachine(capacity int) *VendingMachine {

	return &VendingMachine{capacity: capacity, productsList: make([]Products.Product, 0)}
}

func (vm *VendingMachine) AddProduct(product *Products.Product) {
	vm.productsList = append(vm.productsList, *product)
}

func (vm *VendingMachine) AddBeverage(beverage *Products.Beverage) {
	vm.beverageList = append(vm.beverageList, *beverage)
}

func (vm *VendingMachine) getProducts() []Products.Product {
	return vm.productsList
}

func (vm *VendingMachine) String() string {
	return fmt.Sprintf("Вендинг машина %v вместимостью %v включает в себя следующие продукты %v", vm.vmName,
		vm.capacity, vm.productsList)
}
