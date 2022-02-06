// Package stuple implements type-safe tuples backed by struct.
package stuple

//go:generate go run ../internal/codegen/ -N 7 --outfile codegen_stuple.go
var tuplesTemplate string = `
// Packed{{ .N }} is a struct-backed tuple with {{ .N }} elements.
type Packed{{ .N }}[{{ .T1_T2___TN }} any] struct {
	{{ .RepeatOneToNWithSep "Item__N1 T__N1" "\n\t" }}
}

// Pack{{ .N }} returns an struct-backed tuple of with {{ .N }} items.
func Pack{{ .N }}[{{ .T1_T2___TN }} any](
	{{ .RepeatOneToN "e__N1 T__N1" }},
) Packed{{ .N }}[{{ .T1_T2___TN }}] {
	return Packed{{ .N }}[{{ .T1_T2___TN }}]{
		{{ .RepeatOneToN "e__N1" }},
	}
}

// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed{{ .N }}[{{ .T1_T2___TN }}]) Unpack() ({{ .T1_T2___TN }}) {
	return {{ .RepeatOneToN "t.Item__N1" }}
}

// ToArray returns a {{ .N }}-element untyped array containing the elements of
// the tuple, useful for usage with range loops.
func (t Packed{{ .N }}[{{ .T1_T2___TN }}]) ToArray() [{{ .N }}]any {
	return [{{ .N }}]any{ {{- .RepeatOneToN "t.Item__N1" -}} }
}
`
