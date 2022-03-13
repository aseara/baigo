package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Visitor visitor fun
type Visitor func(shape Shape)

// Shape shape intf
type Shape interface {
	accept(Visitor)
}

// Circle circle shape
type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

// Rectangle rect shape
type Rectangle struct {
	Width, Height int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

// JSONVisitor json visitor
func JSONVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// XMLVisitor xml visitor
func XMLVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JSONVisitor)
		s.accept(XMLVisitor)
	}
}
