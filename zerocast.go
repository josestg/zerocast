package zerocast

import (
	"reflect"
	"unsafe"
)

// StringToBytes converts a string to a byte slice without memory allocation.
// NOTE: The returned byte slice MUST NOT be modified since it shares the same backing array
// with the given string.
func StringToBytes(s string) []byte {
	// Get StringHeader from string
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	// Construct SliceHeader with capacity equal to the length
	sliceHeader := reflect.SliceHeader{Data: stringHeader.Data, Len: stringHeader.Len, Cap: stringHeader.Len}

	// Convert SliceHeader to a byte slice
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

// BytesToString converts bytes to a string without memory allocation.
// NOTE: The given bytes MUST NOT be modified since they share the same backing array
// with the returned string.
func BytesToString(b []byte) string {
	// Obtain SliceHeader from []byte.
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// Construct StringHeader from SliceHeader.
	stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}

	// Convert StringHeader to a string.
	s := *(*string)(unsafe.Pointer(&stringHeader))
	return s
}
