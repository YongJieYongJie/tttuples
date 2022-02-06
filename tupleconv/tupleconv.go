package tttuples

//go:generate go run ../internal/codegen/ -N 7 --outfile codegen_tupleconv.go --imports "github.com/YongJieYongJie/tttuples/atuple,github.com/YongJieYongJie/tttuples/stuple"
var tupleTemplate string = `
// Atos converts a array-backed tuple from tttuples/atuple to an struct-backed
// tuple from tttuples/stuple.
func Atos{{ .N }}[{{ .T1_T2___TN }} any](
	t stuple.Packed{{ .N }}[{{ .T1_T2___TN }}],
) atuple.Packed{{ .N }}[{{ .T1_T2___TN }}] {
	return atuple.Pack{{ .N }}(
		{{ .RepeatOneToN "t.Item__N1" }},
	)
}

// Stoa converts a struct-backed tuple from tttuples/stuple to an array-backed
// tuple from tttuples/atuple.
func Stoa{{ .N }}[{{ .T1_T2___TN }} any](
	t atuple.Packed{{ .N }}[{{ .T1_T2___TN }}],
) stuple.Packed{{ .N }}[{{ .T1_T2___TN }}] {
	{{ .RepeatOneToNWithSep "item__N1, _ := t[__N0].(T__N1)" "\n\t" }}
	return stuple.Packed{{ .N }}[{{ .T1_T2___TN }}]{
		{{ .RepeatOneToN "item__N1" }},
	}
}
`
