package Adapter

import "fmt"

/*
	Adapter pattern provide adapter to translate a interface to a compatible interface
	It allows classes to work together that normally could not because of incompatible interface
*/

type Animal interface {
	Move()
}

type Dog struct {
}

type Crocodile struct {
}

func (d *Dog) Move() {
	fmt.Printf("Dog can move\n")
}

func (c *Crocodile) Slither() {
	fmt.Printf("Crocodile move by slithering\n")
}

type CrocodileAdapter struct {
	*Crocodile
}

func (ca *CrocodileAdapter) Move() {
	ca.Slither()
}

func NewDog() *Dog {
	return &Dog{}
}

func NewCrocodile() *CrocodileAdapter{
	return &CrocodileAdapter{&Crocodile{}}
}
// usage
func ExampleUsage(){
	dog := NewDog()
	crocodile := NewCrocodile()
	crocodile.Move()
	dog.Move()
}