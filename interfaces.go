/*
 Can't assign/pass a struct by value to an interface if it implements one (or more) of it's functions with a pointer receiver, it must be passed or assigned as a pointer.
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
