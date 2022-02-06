// Package atuple implements type-safe tuples backed by array.
package atuple

//go:generate go run ../internal/codegen/ -N 7 --outfile codegen_atuple.go
var tupleTemplate string = `
// Packed{{ .N }} is an array-backed tuple with {{ .N }} elements.
type Packed{{ .N }}[{{ .T1_T2___TN }} any] [{{ .N }}]any

// Pack{{ .N }} returns an array-backed tuple of with {{ .N }} items.
func Pack{{ .N }}[{{ .T1_T2___TN }} any](
	{{ .RepeatOneToN "e__N1 T__N1" }},
) Packed{{ .N }}[{{ .T1_T2___TN }}] {
	return [{{ .N }}]any{
		{{ .RepeatOneToN "e__N1" }},
	}
}
{{- "\n" -}}
{{ range $index, $_ := .OneToN }}
// Item{{ . }} returns the element at index {{ . }} (1-based), with the original
// type.
func (t Packed{{ $.N }}[{{ $.T1_T2___TN }}]) Item{{ . }}() T{{ . }} {
	out, _ := t[{{ $index }}].(T{{ . }})
	return out
}
{{ end }}
{{- "\n" -}}
// Unpack returns each element of the tuple as separate return values and with
// their original type, useful for assigning to multiple variables at once or
// passing into other functions.
func (t Packed{{ .N }}[{{ .T1_T2___TN }}]) Unpack() ({{ .T1_T2___TN }}) {
	return {{ .RepeatOneToN "t[__N0].(T__N1)" }}
}
`
