package Bridge

import "fmt"

/*
	TODO: Add description later
*/

type Drawer interface {
	Draw()
}

type Point struct {
	x int
	y int
}

type Circle struct {
	Drawer   Drawer
	Position Point
	Radius   int
}

type OpenGPL struct{}
type Direct2D struct{}

func (circle Circle) DrawCircle(){
	circle.Drawer.Draw()
}

func (drawer *OpenGPL) Draw() {
	fmt.Printf("Draw by OpenGPL\n")
}

func (drawer *Direct2D) Draw(){
	fmt.Printf("Draw by Direct2D \n")
}

//usage
func ExampleUsage(){
	circle := &Circle{
		Drawer:   nil,
		Position: Point{3,4},
		Radius:   5,
	}
	circle.Drawer = &OpenGPL{}
	circle.DrawCircle()
	circle.Drawer = &Direct2D{}
	circle.DrawCircle()
}
