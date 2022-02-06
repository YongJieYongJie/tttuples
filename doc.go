// Package tttuples (pronounced t-t-tuples, with a stutter), implements tuples
// for ad-hoc type-safe grouping of heterogenous elements. The two subpackages
// (atuple and stuple) implements tuples using array and struct respectively,
// with slightly different API.
//
// Refer to example for a comparison of the two implementations, alongside Go's
// builtin.
//
// Motivation
//
// Go's builtin container types (slice, array, channels) generally do not
// support heterogenous without resorting to using interface{}/any and
// sacrificing type-safety.
//
// The main builtin way of group together heterogenous data is to use struct,
// which may be anonymous.
//
// However, usage of struct (in operations such as declaration and definition,
// and multiple field access) can get unwieldy without additional helper
// methods. This package aims to provide an alternative.
//
// Goals / Design Philosophy
//
// The main design considerations (in order of priority) are as follows:
//     1. User-friendly APIs that play nice with Go langauge.
//     2. Static type safety.
//     3. Simple.
package tttuples
