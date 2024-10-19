package main

import "fmt"

type Product struct {
	name  string
	price float32
	qty   uint
}

// constructor
func newProduct(name string, price float32, qty uint) *Product {
	return &Product{name: name, price: price, qty: qty}
}

// member function
func (p *Product) setPrice(price float32) {
	p.price = price
}

// member function
func (p *Product) setQty(qty uint) {
	p.qty = qty
}

// method function getter
func (p Product) getPrice() float32 {
	return p.price
}

func (p Product) getQty() uint {
	return p.qty
}

func main() {

	var p1 *Product = newProduct("apple", 100, 30)

	fmt.Println("old price :", p1.getPrice())
	p1.setPrice(120.0)
	fmt.Println("new price :", p1.getPrice())

	var p2 = Product{"Orange", 60.0, 10}
	fmt.Println("Orange old price :", p2.getPrice())
	p2.setPrice(45.0)
	p2.setQty(15)
	fmt.Printf("orange new price %f : qty %d \n", p2.getPrice(), p2.getQty())

}
