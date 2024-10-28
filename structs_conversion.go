/**
When we have a struct type and define an alias type for it, both types have the same underlying type and can be converted into one another. 
But defining a new type based on a struct, can't be converted from one to another
*/

type number struct {
	f	float32
}

type nrAlias = number // alias type - we can convert number types to and from nrAlias
type nr number   // new distinct type - we can't convert number types to or from nr types

func main() {
	a := number{5.0}
	b := nr{5.0}
	c := nrAlias{5.0}
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var d number = b    // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	var d number = number(b)
	// b = c  // compile-error: cannot use c (variable of type number) as type nr in assignment
	// an alias doesn't need conversion:
	a = c
	c = a
	fmt.Println(a, b, c, d)
}
