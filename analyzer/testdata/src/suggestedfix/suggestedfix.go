//nolint:all
package suggestedfix

import "suggestedfix/d"

type SimpleStruct struct {
	A int
	B string
	C bool
}

type MixedStruct struct {
	ID       int
	Name     string
	Exported bool
	hidden   string
}

type NestedStruct struct {
	Simple SimpleStruct
	Count  int
}

type Nested2 struct {
	Simple d.SimpleStruct
}

func testEmptyStruct() {
	_ = SimpleStruct{} // want "suggestedfix.SimpleStruct is missing fields A, B, C"
}

func testPartialStruct() {
	_ = SimpleStruct{ // want "suggestedfix.SimpleStruct is missing fields B, C"
		A: 42,
	}
}

func testMixedFields() {
	_ = MixedStruct{} // want "suggestedfix.MixedStruct is missing fields ID, Name, Exported, hidden"
}

func testNestedStruct() {
	_ = NestedStruct{} // want "suggestedfix.NestedStruct is missing fields Simple, Count"
}

func testNestedStruct_composite() {
	_ = Nested2{Simple: d.SimpleStruct{}} // want "d.SimpleStruct is missing fields A, B, C"
}

func testNestedStruct_composite_2() {
	_ = Nested2{} // want "suggestedfix.Nested2 is missing field Simple"
}
