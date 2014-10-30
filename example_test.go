// Copyright (c) 2014 Datacratic. All rights reserved.

package set_test

import (
	"github.com/datacratic/goset"

	"fmt"
)

func Example_Basics() {

	x := set.NewString("a", "b", "c")

	fmt.Println("len(x):", len(x))
	fmt.Println("x.Test(a):", x.Test("a"))
	fmt.Println("x.Test(d):", x.Test("d"))

	y := x.Copy()
	y.Put("d")
	y.Del("a")

	fmt.Print("x: { ")
	for v := range x {
		fmt.Printf("%s ", v)
	}
	fmt.Println("}")

	fmt.Println("y:", y)

	// Output:
	// len(x): 3
	// x.Test(a): true
	// x.Test(d): false
	// x: { a b c }
	// y: { b c d }
}

func Example_Operands() {

	x := set.NewString("a", "b", "c")
	y := set.NewString("b", "c", "d")
	
	fmt.Println("x.Union(y):", x.Union(y))
	fmt.Println("x.Intersect(y):", x.Intersect(y))
	fmt.Println("x.Difference(y):", x.Difference(y))

	// Output:
	// x.Union(y): { a b c d }
	// x.Intersect(y): { b c }
	// x.Difference(y): { a }
}
