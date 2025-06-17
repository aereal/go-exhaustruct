package structure

import (
	"go/types"
	"testing"
)

func TestZeroValue(t *testing.T) {
	tests := []struct {
		name     string
		typ      types.Type
		expected string
	}{
		// Basic types
		{"bool", types.Typ[types.Bool], "false"},
		{"int", types.Typ[types.Int], "0"},
		{"int32", types.Typ[types.Int32], "0"},
		{"uint", types.Typ[types.Uint], "0"},
		{"uint64", types.Typ[types.Uint64], "0"},
		{"float32", types.Typ[types.Float32], "0"},
		{"float64", types.Typ[types.Float64], "0"},
		{"complex64", types.Typ[types.Complex64], "0"},
		{"complex128", types.Typ[types.Complex128], "0"},
		{"string", types.Typ[types.String], `""`},
		{"uintptr", types.Typ[types.Uintptr], "0"},

		// Pointer types
		{"*int", types.NewPointer(types.Typ[types.Int]), "nil"},
		{"*string", types.NewPointer(types.Typ[types.String]), "nil"},

		// Slice types
		{"[]int", types.NewSlice(types.Typ[types.Int]), "nil"},
		{"[]string", types.NewSlice(types.Typ[types.String]), "nil"},

		// Map types
		{"map[string]int", types.NewMap(types.Typ[types.String], types.Typ[types.Int]), "nil"},

		// Channel types
		{"chan int", types.NewChan(types.SendRecv, types.Typ[types.Int]), "nil"},

		// Interface types
		{"interface{}", types.NewInterfaceType(nil, nil), "nil"},

		// Function types
		{"func()", types.NewSignatureType(nil, nil, nil, nil, nil, false), "nil"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ZeroValue(tt.typ)
			if result != tt.expected {
				t.Errorf("ZeroValue(%v) = %v, want %v", tt.typ, result, tt.expected)
			}
		})
	}
}

func TestStructZeroValue(t *testing.T) {
	// Create a simple struct type: struct { A int; B string }
	fields := []*types.Var{
		types.NewField(0, nil, "A", types.Typ[types.Int], false),
		types.NewField(0, nil, "B", types.Typ[types.String], false),
	}
	structType := types.NewStruct(fields, nil)

	result := ZeroValue(structType)
	expected := `{A: 0, B: ""}`

	if result != expected {
		t.Errorf("ZeroValue(struct) = %v, want %v", result, expected)
	}
}

func TestArrayZeroValue(t *testing.T) {
	// Test small array
	smallArray := types.NewArray(types.Typ[types.Int], 3)
	result := ZeroValue(smallArray)
	expected := "{0, 0, 0}"

	if result != expected {
		t.Errorf("ZeroValue([3]int) = %v, want %v", result, expected)
	}

	// Test large array
	largeArray := types.NewArray(types.Typ[types.String], 10)
	result = ZeroValue(largeArray)
	expected = `{"" /* ... 10 elements */}`

	if result != expected {
		t.Errorf("ZeroValue([10]string) = %v, want %v", result, expected)
	}
}

func TestEmptyStruct(t *testing.T) {
	emptyStruct := types.NewStruct(nil, nil)
	result := ZeroValue(emptyStruct)
	expected := "{}"

	if result != expected {
		t.Errorf("ZeroValue(empty struct) = %v, want %v", result, expected)
	}
}
