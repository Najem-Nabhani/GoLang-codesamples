/*
 1. Can't assign/pass a struct by value to an interface if it implements one (or more) of it's functions with a pointer receiver, it must be passed or assigned as a pointer.
*/

type Getter interface {
  Get() int
}

type Box struct {
  number int
}

func (b *Box) Get() int{
  return b.number
}

func main() {
  b := Box{}
  //getter := Getter(b) //ERROR
  getter := Getter(&b) // Valid!
}

/*
 2. Casting an interface to another interface
*/

type EnclosedShape interface {
	CalculateArea() int
}

type Shaper interface {
	EnclosedShape
	GetWidth() int
	GetHeight() int
}

type Rectangle struct {
	height int
	width  int
}

func NewRectangle(w, h int) Rectangle {
	return Rectangle{width: w, height: h}
}

func (r Rectangle) GetWidth() int {
	return r.width
}

func (r Rectangle) CalculateArea() int {
	return r.height * r.width
}

func (r Rectangle) GetHeight() int {
	return r.height
}

func main() {
rectangle := interfaces.NewRectangle(3, 4)
	shapes := []interfaces.Shaper{rectangle}
	for _, shape := range shapes {
		fmt.Println(shape.CalculateArea())
	}

	if shape, ok := shapes[0].(interfaces.EnclosedShape); ok {
		fmt.Printf("shape is of type %T", shape) // rectangle
		shape.CalculateArea() // valid
  //shape.GetWidth() // Error! shape was casted from Shaper to EnclosedShape which doesn't implement GetWidth
	}
}

/*
 3. You can use the type-switch syntax to determine the type of the interface.
*/

package main
import (
  "fmt"
  "math"
)

type Square struct {
  side float32
}

type Circle struct {
  radius float32
}

type Shaper interface {
  Area() float32
}

func main() {
  var areaIntf Shaper
  sq1 := new(Square)
  sq1.side = 5
  areaIntf = sq1

  switch t := areaIntf.(type) {
    case *Square:
      fmt.Printf("Type Square %T with value %v\n", t, t)

    case *Circle:
      fmt.Printf("Type Circle %T with value %v\n", t, t)

    default:
      fmt.Printf("Unexpected type %T", t)
    }
}

func (sq *Square) Area() float32 {
  return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
  return ci.radius * ci.radius * math.Pi
}
