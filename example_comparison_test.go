package tttuples_test

import (
	"fmt"
	"reflect"

	"github.com/YongJieYongJie/tttuples/atuple"
	"github.com/YongJieYongJie/tttuples/stuple"
)

// The following example compares the API of
//     (1) Go's builtin elements,
//     (2) array-backed tuple from atuple, and
//     (3) struct-backed tuple from stuple,
// in relation to
//     (a) declaration and initialization,
//     (b) single element access,
//     (c) multiple element access (via range loops), and
//     (d) multiple element access (passing as arguments to functions),
// for the purpose of evaluating
//     (A) user-friendliness,
//     (B) type-safety.
func Example_comparison() {

	// (a) declaration and initialization

	// (a)(1) declaration and initialization - Go's builtin elements

	var slice_ = []int{1, 2, 3} // slice

	var array = [3]int{1, 2, 3} // array

	type structDef struct{ field1, field2, field3 int } // named struct
	var structNamed = structDef{1, 2, 3}

	var structAnon = struct{ field1, field2, field3 int }{1, 2, 3} // annonymous struct

	var structHetero = struct {
		field1         int
		field2, field3 string
	}{1, "b", "c"} // heterogenous struct

	var chan_ = func() <-chan int { // channel
		// Note: a func is used because a chan seeded with particuler elements
		// cannot be declared and initialized in a single line.
		c := make(chan int, 3)
		defer close(c)
		c <- 1
		c <- 2
		c <- 3
		return c
	}

	// General note: Go's builtin container types does not support heterogenous
	// data (not without resorting to interface{} / any, and losing type safety).

	// (a)(2) declaration and initialization - array-backed tuple

	var atupleHomo = atuple.Pack3(1, 2, 3)
	var atupleHetero = atuple.Pack3(1, "b", "c")

	// (a)(3) declaration and initialization - struct-backed tuple

	var stupleHomo = stuple.Pack3(1, 2, 3)
	var stupleHetero = stuple.Pack3(1, "b", "c")

	// (b) single element access

	// (b)(1) single element access - Go's builtin elements

	fmt.Printf("slice_[0]: %v\n", slice_[0])
	fmt.Printf("array[0]: %v\n", array[0])
	fmt.Printf("structNamed.field1: %v\n", structNamed.field1)
	fmt.Printf("structAnon.field1: %v\n", structAnon.field1)
	fmt.Printf("<- chan: %v\n", <-chan_())

	// (b)(2) single element access - array-backed tuple

	fmt.Printf("atupleHomo[0]: %v\n", atupleHomo[0])
	fmt.Printf("type of element 0 in atupleHomo: %v\n", reflect.TypeOf(atupleHomo).Elem())
	fmt.Printf("atupleHomo.Item1(): %v\n", atupleHomo.Item1())
	fmt.Printf("return type of atupleHomo.Item1(): %v\n", reflect.TypeOf(atupleHomo.Item1).Out(0))

	// General note: atuple supports two way of accessing single element. The
	// first way is by using the square bracket index operator, which looks like
	// normal Go code, but does not offer type safety as the underlying
	// interface{} is returned. The second way is by using the ItemN() method,
	// which preserves type information, but (depending on the person) may look
	// less like normal Go builtin.

	// (b)(3) single element access - struct-backed tuple

	fmt.Printf("stupleHomo.Item1: %v\n", stupleHomo.Item1)
	fmt.Printf("type of field Item1 in stupleHomo: %v\n", reflect.TypeOf(stupleHomo).Field(0).Type)

	// General note: stuple supports a single way of accessing a single element
	// via the dot operator. Type information is preserved.

	// (c) multiple element access (via range loop)

	// (c)(1) multiple element access (via range loop) - Go's builtin elements

	for i, elem := range slice_ {
		fmt.Printf("range loop over slice_ at index %d. Element value: %v\n", i, elem)
	}
	for i, elem := range array {
		fmt.Printf("range loop over array at index %d. Element value: %v\n", i, elem)
	}
	for i := 0; i < reflect.TypeOf(structNamed).NumField(); i++ {
		fmt.Printf(
			"using reflection to loop over fields of structNamed. Field %s's value: %v\n",
			reflect.TypeOf(structNamed).Field(i).Name,
			reflect.ValueOf(structNamed).Field(i),
		)
	}
	for i := 0; i < reflect.TypeOf(structAnon).NumField(); i++ {
		fmt.Printf(
			"using reflection to loop over fields of structAnon. Field %s's value: %v\n",
			reflect.TypeOf(structAnon).Field(i).Name,
			reflect.ValueOf(structAnon).Field(i),
		)
	}
	for elem := range chan_() {
		fmt.Printf("range loop over channel. Element value: %v\n", elem)
	}

	// (c)(2) multiple element access (via range loop) - array-backed tuple

	for i, elem := range atupleHomo {
		fmt.Printf("range loop over atupleHomo at index %d. Element value: %v\n", i, elem)
	}

	// (c)(3) multiple element access (via range loop) - struct-backed tuple

	for i, elem := range stupleHomo.ToArray() {
		fmt.Printf("range loop over stupleHomo at index %d. Element value: %v\n", i, elem)
	}

	// (d) multiple element access (passing as arguments to functions)

	// For the purpose of this test (d), we are interested in the following types
	// of function signatures:

	multiArg := func(int, int, int) {}
	variadic := func(...int) {}
	multiArgVariadic := func(int, ...int) {}

	multiArgHetero := func(int, string, string) {}
	variadicHetero := func(int, ...string) {}

	// (d)(1) multiple element access (passing as arguments to functions) - Go's builtin elements

	// slice_

	multiArg(slice_[0], slice_[1], slice_[2])
	variadic(slice_...)
	multiArgVariadic(slice_[0], slice_[1:]...)

	// General note: interfacing between a builtin slice and a function with
	// multiple arguments can be awkward and requires indexing operation.
	//
	// However, with additional insights into the type and length of the slice
	// elements, it is possible to write wrapper/adapter the greatly improves the
	// interfacing experience:
	//
	// type wrapped []int
	// func (s wrapped) Unpack() (int, int, int) {
	// 	return s[0], s[1], s[2]
	// }
	multiArg(wrapped(slice_).Unpack())
	variadic(wrapped(slice_).Unpack())
	multiArgVariadic(wrapped(slice_).Unpack())

	// array

	multiArg(array[0], array[1], array[2])
	variadic(array[:]...) // note the additional conversion to slice required
	multiArgVariadic(array[0], array[1:]...)

	// General note: as with slice, a similar wrapper/adapter can be written to
	// simplify the code.

	// structNamed (similar for structAnon)

	multiArg(structNamed.field1, structNamed.field2, structNamed.field3)
	variadic(structNamed.field1, structNamed.field2, structNamed.field3)
	multiArgVariadic(structNamed.field1, structNamed.field2, structNamed.field3)

	multiArgHetero(structHetero.field1, structHetero.field2, structHetero.field3)
	variadicHetero(structHetero.field1, structHetero.field2, structHetero.field3)

	// General note: as with slice, a similar wrapper/adapter can be written to
	// simplify the code. Alternatively, additional method can be added directly
	// to the struct.

	// General note 2: Here we can also see that a struct is Go's builtin way of
	// grouping heterogenous data.

	c1 := chan_()
	multiArg(<-c1, <-c1, <-c1)
	c2 := chan_()
	variadic(<-c2, <-c2, <-c2)
	c3 := chan_()
	multiArgVariadic(<-c3, <-c3, <-c3)

	// (d)(2) multiple element access (passing as arguments to functions) - array-backed tuple

	multiArg(atupleHomo.Unpack())
	variadic(atupleHomo.Unpack())
	multiArgVariadic(atupleHomo.Unpack())

	multiArgHetero(atupleHetero.Unpack())
	variadicHetero(atupleHetero.Unpack())

	// (d)(2) multiple element access (passing as arguments to functions) - struct-backed tuple

	multiArg(stupleHomo.Unpack())
	variadic(stupleHomo.Unpack())
	multiArgVariadic(stupleHomo.Unpack())

	multiArgHetero(stupleHetero.Unpack())
	variadicHetero(stupleHetero.Unpack())
}

type wrapped []int

func (s wrapped) Unpack() (int, int, int) {
	return s[0], s[1], s[2]
}
