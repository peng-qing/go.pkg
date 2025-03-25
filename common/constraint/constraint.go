package constraint

// None is an empty struct.
type None struct{}

// Pointer is a constraint that permits any pointer type.
type Pointer[T any] interface {
	~*T
}

// Integer is a constraint that permits any integer type.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex-number type.
type Complex interface {
	~complex64 | ~complex128
}
