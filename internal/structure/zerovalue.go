package structure

import (
	"fmt"
	"go/types"
	"strings"
)

// ZeroValue returns the zero value representation for a given Go type as a string.
// This is used to generate suggested fixes for missing fields in struct literals.
func ZeroValue(typ types.Type) string {
	switch t := typ.Underlying().(type) {
	case *types.Basic:
		return basicZeroValue(t)
	case *types.Pointer:
		return "nil"
	case *types.Slice:
		return "nil"
	case *types.Map:
		return "nil"
	case *types.Chan:
		return "nil"
	case *types.Interface:
		return "nil"
	case *types.Signature:
		return "nil"
	case *types.Struct:
		return structZeroValue(t)
	case *types.Array:
		return arrayZeroValue(t)
	default:
		// For unknown types, return empty braces as a fallback
		return "{}"
	}
}

func basicZeroValue(basic *types.Basic) string {
	switch basic.Kind() {
	case types.Bool:
		return "false"
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
		return "0"
	case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr:
		return "0"
	case types.Float32, types.Float64:
		return "0"
	case types.Complex64, types.Complex128:
		return "0"
	case types.String:
		return `""`
	case types.UnsafePointer:
		return "nil"
	default:
		return `""`
	}
}

func structZeroValue(strct *types.Struct) string {
	if strct.NumFields() == 0 {
		return "{}"
	}

	var fields []string
	for i := 0; i < strct.NumFields(); i++ {
		field := strct.Field(i)
		if field.Exported() {
			fieldZero := ZeroValue(field.Type())
			fields = append(fields, fmt.Sprintf("%s: %s", field.Name(), fieldZero))
		}
	}

	if len(fields) == 0 {
		return "{}"
	}

	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}

func arrayZeroValue(arr *types.Array) string {
	elemZero := ZeroValue(arr.Elem())
	if arr.Len() == 0 {
		return "{}"
	}

	// For small arrays, generate all elements
	if arr.Len() <= 5 {
		var elements []string
		for i := int64(0); i < arr.Len(); i++ {
			elements = append(elements, elemZero)
		}
		return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
	}

	// For larger arrays, just show the pattern
	return fmt.Sprintf("{%s /* ... %d elements */}", elemZero, arr.Len())
}
