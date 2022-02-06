// Code generated by codegen.go DO NOT EDIT.
// Version: 0.0.0
// Checksum: 867318174519798771

package stuple

// Packed2 is a struct-backed tuple with 2 elements.
type Packed2[T1, T2 any] struct {
	Item1 T1
	Item2 T2
}

// Pack2 returns an struct-backed tuple of with 2 items.
func Pack2[T1, T2 any](
	e1 T1, e2 T2,
) Packed2[T1, T2] {
	return Packed2[T1, T2]{
		e1, e2,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed2[T1, T2]) Unpack() (T1, T2) {
	return t.Item1, t.Item2
}

// ToArray returns a 2-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed2[T1, T2]) ToArray() [2]any {
	return [2]any{t.Item1, t.Item2}
}

// Packed3 is a struct-backed tuple with 3 elements.
type Packed3[T1, T2, T3 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
}

// Pack3 returns an struct-backed tuple of with 3 items.
func Pack3[T1, T2, T3 any](
	e1 T1, e2 T2, e3 T3,
) Packed3[T1, T2, T3] {
	return Packed3[T1, T2, T3]{
		e1, e2, e3,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed3[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return t.Item1, t.Item2, t.Item3
}

// ToArray returns a 3-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed3[T1, T2, T3]) ToArray() [3]any {
	return [3]any{t.Item1, t.Item2, t.Item3}
}

// Packed4 is a struct-backed tuple with 4 elements.
type Packed4[T1, T2, T3, T4 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
}

// Pack4 returns an struct-backed tuple of with 4 items.
func Pack4[T1, T2, T3, T4 any](
	e1 T1, e2 T2, e3 T3, e4 T4,
) Packed4[T1, T2, T3, T4] {
	return Packed4[T1, T2, T3, T4]{
		e1, e2, e3, e4,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed4[T1, T2, T3, T4]) Unpack() (T1, T2, T3, T4) {
	return t.Item1, t.Item2, t.Item3, t.Item4
}

// ToArray returns a 4-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed4[T1, T2, T3, T4]) ToArray() [4]any {
	return [4]any{t.Item1, t.Item2, t.Item3, t.Item4}
}

// Packed5 is a struct-backed tuple with 5 elements.
type Packed5[T1, T2, T3, T4, T5 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
}

// Pack5 returns an struct-backed tuple of with 5 items.
func Pack5[T1, T2, T3, T4, T5 any](
	e1 T1, e2 T2, e3 T3, e4 T4, e5 T5,
) Packed5[T1, T2, T3, T4, T5] {
	return Packed5[T1, T2, T3, T4, T5]{
		e1, e2, e3, e4, e5,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5) {
	return t.Item1, t.Item2, t.Item3, t.Item4, t.Item5
}

// ToArray returns a 5-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed5[T1, T2, T3, T4, T5]) ToArray() [5]any {
	return [5]any{t.Item1, t.Item2, t.Item3, t.Item4, t.Item5}
}

// Packed6 is a struct-backed tuple with 6 elements.
type Packed6[T1, T2, T3, T4, T5, T6 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
	Item6 T6
}

// Pack6 returns an struct-backed tuple of with 6 items.
func Pack6[T1, T2, T3, T4, T5, T6 any](
	e1 T1, e2 T2, e3 T3, e4 T4, e5 T5, e6 T6,
) Packed6[T1, T2, T3, T4, T5, T6] {
	return Packed6[T1, T2, T3, T4, T5, T6]{
		e1, e2, e3, e4, e5, e6,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed6[T1, T2, T3, T4, T5, T6]) Unpack() (T1, T2, T3, T4, T5, T6) {
	return t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6
}

// ToArray returns a 6-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed6[T1, T2, T3, T4, T5, T6]) ToArray() [6]any {
	return [6]any{t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6}
}

// Packed7 is a struct-backed tuple with 7 elements.
type Packed7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
	Item5 T5
	Item6 T6
	Item7 T7
}

// Pack7 returns an struct-backed tuple of with 7 items.
func Pack7[T1, T2, T3, T4, T5, T6, T7 any](
	e1 T1, e2 T2, e3 T3, e4 T4, e5 T5, e6 T6, e7 T7,
) Packed7[T1, T2, T3, T4, T5, T6, T7] {
	return Packed7[T1, T2, T3, T4, T5, T6, T7]{
		e1, e2, e3, e4, e5, e6, e7,
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed7[T1, T2, T3, T4, T5, T6, T7]) Unpack() (T1, T2, T3, T4, T5, T6, T7) {
	return t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6, t.Item7
}

// ToArray returns a 7-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed7[T1, T2, T3, T4, T5, T6, T7]) ToArray() [7]any {
	return [7]any{t.Item1, t.Item2, t.Item3, t.Item4, t.Item5, t.Item6, t.Item7}
}