package main
import "fmt"

func trace(s string) {
        fmt.Println("entering:", s)
    } // entering func.
func untrace(s string) {
        fmt.Println("leaving:", s)
    } // leaving func.

func a() {
    trace("a")
    defer untrace("a") // untracing via defer
    fmt.Println("in a")
}

func b() {
    trace("b")
    defer untrace("b") // untracing via defer
    fmt.Println("in b")
    a()
}

func main() {
    b()
}

/*
OUTPUT:

entering: b
in b
entering: a
in a
leaving: a
leaving: b

*/
