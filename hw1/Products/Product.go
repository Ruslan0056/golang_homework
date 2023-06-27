package Products

import "fmt"

type Product struct {
	productId       int
	productName     string
	productCategory string
	price           float64
}

func NewProduct(productId int, productName string, productCategory string, price float64) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("Цена указана некорректно!")
	}
	return &Product{productId, productName, productCategory, price}, nil
}

func (p *Product) getName() string {
	return p.productName
}

func (p *Product) getCategory() string {
	return p.productCategory
}

func (p *Product) getId() int {
	return p.productId
}

func (p *Product) getPrice() float64 {
	return p.price
}

func (p *Product) String() string {
	return fmt.Sprintf("Продукт %s из категории %s ценой %.0f под номером %v.",
		p.productName, p.productCategory, p.price, p.productId)
}
