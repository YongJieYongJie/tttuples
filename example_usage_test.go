package tttuples_test

import (
	"fmt"

	"github.com/YongJieYongJie/tttuples/atuple"
)

func Example_usage() {
	// import "github.com/YongJieYongJie/tttuples/atuple"

	// Save as otherwise noted below, this example applies to stuple as well.

	fmt.Println("Basic declaration and initialization:")

	arrayTuple := atuple.Pack3(123, "abc", "def")
	fmt.Printf("%v\n", arrayTuple)
	fmt.Printf("%#v\n", arrayTuple)

	fmt.Println("\nUnpacking tuple items:")

	item1, item2, item3 := arrayTuple.Unpack()
	fmt.Println(item1)
	fmt.Println(item2)
	fmt.Println(item3)

	fmt.Println("\nUnpacking tuple items into function arguments:")

	MultiArg := func(a1 int, a2 string, a3 string) { fmt.Printf("MultiArg() called with %v, %v, and %v\n", a1, a2, a3) }
	Variadic := func(a1 ...any) { fmt.Printf("Variadic() called with %v\n", a1) }
	Mixed := func(a1 int, a2 ...string) { fmt.Printf("Mixed() called with %v, and %v\n", a1, a2) }

	MultiArg(arrayTuple.Unpack())
	Variadic(arrayTuple.Unpack())
	Mixed(arrayTuple.Unpack())

	fmt.Println("\nRange loop:")

	for _, item := range arrayTuple { // Note: for stuple, we need to call .ToArray() on the tuple object first.
		fmt.Println(item)
	}

	fmt.Println("\nDeclaring and initializing slice of tuples:")

	arrayTuples := []atuple.Packed3[int, string, float64]{
		{1, "b", 3.3},
		{4, "e", 6.6},
	}
	fmt.Printf("%v\n", arrayTuples)
	fmt.Printf("%#v\n", arrayTuples)

	// Output:
	// Basic declaration and initialization:
	// [123 abc def]
	// atuple.Packed3[int,string,string]{123, "abc", "def"}
	//
	// Unpacking tuple items:
	// 123
	// abc
	// def
	//
	// Unpacking tuple items into function arguments:
	// MultiArg() called with 123, abc, and def
	// Variadic() called with [123 abc def]
	// Mixed() called with 123, and [abc def]
	//
	// Range loop:
	// 123
	// abc
	// def
	//
	// Declaring and initializing slice of tuples:
	// [[1 b 3.3] [4 e 6.6]]
	// []atuple.Packed3[int,string,float64]{atuple.Packed3[int,string,float64]{1, "b", 3.3}, atuple.Packed3[int,string,float64]{4, "e", 6.6}}
}
