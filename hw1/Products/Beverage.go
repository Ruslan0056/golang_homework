package Products

type Beverage struct {
	*Product
	volume int
}

func NewBeverage(productId int, productName string, productCategory string, price float64, volume int) (*Beverage, error) {
	product, err := NewProduct(productId, productName, productCategory, price)
	if err != nil {
		return nil, err
	}
	return &Beverage{product, volume}, err
}

func (b *Beverage) GetVolume() int {
	return b.volume
}
