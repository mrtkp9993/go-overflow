// go-overflow is a simple library to check for integer overflow in Go.
package overflow

// AddInt8 adds two int8 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddInt8(a, b int8) (int8, bool) {
	c := a + b
	if (c > a) == (b < 0) {
		return c, true
	}
	return c, false
}

// AddInt16 adds two int16 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddInt16(a, b int16) (int16, bool) {
	c := a + b
	if (c > a) == (b < 0) {
		return c, true
	}
	return c, false
}

// AddInt32 adds two int32 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddInt32(a, b int32) (int32, bool) {
	c := a + b
	if (c > a) == (b < 0) {
		return c, true
	}
	return c, false
}

// AddInt64 adds two int64 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddInt64(a, b int64) (int64, bool) {
	c := a + b
	if (c > a) == (b < 0) {
		return c, true
	}
	return c, false
}

// AddUint8 adds two uint8 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddUint8(a, b uint8) (uint8, bool) {
	c := a + b
	if c < a {
		return c, true
	}
	return c, false
}

// AddUint16 adds two uint16 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddUint16(a, b uint16) (uint16, bool) {
	c := a + b
	if c < a {
		return c, true
	}
	return c, false
}

// AddUint32 adds two uint32 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddUint32(a, b uint32) (uint32, bool) {
	c := a + b
	if c < a {
		return c, true
	}
	return c, false
}

// AddUint64 adds two uint64 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func AddUint64(a, b uint64) (uint64, bool) {
	c := a + b
	if c < a {
		return c, true
	}
	return c, false
}

// SubInt8 subtracts two int8 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubInt8(a, b int8) (int8, bool) {
	c := a - b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

// SubInt16 subtracts two int16 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubInt16(a, b int16) (int16, bool) {
	c := a - b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

// SubInt32 subtracts two int32 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubInt32(a, b int32) (int32, bool) {
	c := a - b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

// SubInt64 subtracts two int64 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubInt64(a, b int64) (int64, bool) {
	c := a - b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

// SubUint8 subtracts two uint8 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubUint8(a, b uint8) (uint8, bool) {
	c := a - b
	if c > a {
		return c, true
	}
	return c, false
}

// SubUint16 subtracts two uint16 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubUint16(a, b uint16) (uint16, bool) {
	c := a - b
	if c > a {
		return c, true
	}
	return c, false
}

// SubUint32 subtracts two uint32 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubUint32(a, b uint32) (uint32, bool) {
	c := a - b
	if c > a {
		return c, true
	}
	return c, false
}

// SubUint64 subtracts two uint64 values and checks for underflow.
// Returns the result and a boolean indicating if underflow occurred.
func SubUint64(a, b uint64) (uint64, bool) {
	c := a - b
	if c > a {
		return c, true
	}
	return c, false
}

// MulInt8 multiplies two int8 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulInt8(a, b int8) (int8, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulInt16 multiplies two int16 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulInt16(a, b int16) (int16, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulInt32 multiplies two int32 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulInt32(a, b int32) (int32, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulInt64 multiplies two int64 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulInt64(a, b int64) (int64, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulUint8 multiplies two uint8 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulUint8(a, b uint8) (uint8, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulUint16 multiplies two uint16 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulUint16(a, b uint16) (uint16, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulUint32 multiplies two uint32 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulUint32(a, b uint32) (uint32, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// MulUint64 multiplies two uint64 values and checks for overflow.
// Returns the result and a boolean indicating if overflow occurred.
func MulUint64(a, b uint64) (uint64, bool) {
	c := a * b
	if a != 0 && c/a != b {
		return c, true
	}
	return c, false
}

// DivInt8 divides two int8 values and checks for division by zero or overflow.
// Returns the result and a boolean indicating if an error occurred.
// Integer division overflows in one specific case: dividing the smallest negative value for the data type (see Maximum and Minimum Values) by -1. That’s because the correct result, which is the corresponding positive number, does not fit (see Integer Overflow) in the same number of bits.
// Source: https://www.gnu.org/software/c-intro-and-ref/manual/html_node/Division-and-Remainder.html
func DivInt8(a, b int8) (int8, bool) {
	if b == 0 {
		return 0, true
	}

	// overflow
	if a == -128 && b == -1 {
		return 0, true
	}

	return a / b, false
}

// DivInt16 divides two int16 values and checks for division by zero or overflow.
// Returns the result and a boolean indicating if an error occurred.
// Integer division overflows in one specific case: dividing the smallest negative value for the data type (see Maximum and Minimum Values) by -1. That’s because the correct result, which is the corresponding positive number, does not fit (see Integer Overflow) in the same number of bits.
// Source: https://www.gnu.org/software/c-intro-and-ref/manual/html_node/Division-and-Remainder.html
func DivInt16(a, b int16) (int16, bool) {
	if b == 0 {
		return 0, true
	}

	// overflow
	if a == -32768 && b == -1 {
		return 0, true
	}

	return a / b, false
}

// DivInt32 divides two int32 values and checks for division by zero or overflow.
// Returns the result and a boolean indicating if an error occurred.
// Integer division overflows in one specific case: dividing the smallest negative value for the data type (see Maximum and Minimum Values) by -1. That’s because the correct result, which is the corresponding positive number, does not fit (see Integer Overflow) in the same number of bits.
// Source: https://www.gnu.org/software/c-intro-and-ref/manual/html_node/Division-and-Remainder.html
func DivInt32(a, b int32) (int32, bool) {
	if b == 0 {
		return 0, true
	}

	// overflow
	if a == -2147483648 && b == -1 {
		return 0, true
	}

	return a / b, false
}

// DivInt64 divides two int64 values and checks for division by zero or overflow.
// Returns the result and a boolean indicating if an error occurred.
// Integer division overflows in one specific case: dividing the smallest negative value for the data type (see Maximum and Minimum Values) by -1. That’s because the correct result, which is the corresponding positive number, does not fit (see Integer Overflow) in the same number of bits.
// Source: https://www.gnu.org/software/c-intro-and-ref/manual/html_node/Division-and-Remainder.html
func DivInt64(a, b int64) (int64, bool) {
	if b == 0 {
		return 0, true
	}

	// overflow
	if a == -9223372036854775808 && b == -1 {
		return 0, true
	}

	return a / b, false
}
