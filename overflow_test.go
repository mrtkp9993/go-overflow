package overflow

import (
	"testing"
)

func TestAddInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name  string
		args  args
		want  int8
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 127, b: 1}, want: -128, want1: true},
		{name: "Underflow", args: args{a: -128, b: -1}, want: 127, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddInt8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddInt8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddInt8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 32767, b: 1}, want: -32768, want1: true},
		{name: "Underflow", args: args{a: -32768, b: -1}, want: 32767, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddInt16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddInt16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 2147483647, b: 1}, want: -2147483648, want1: true},
		{name: "Underflow", args: args{a: -2147483648, b: -1}, want: 2147483647, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddInt32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 9223372036854775807, b: 1}, want: -9223372036854775808, want1: true},
		{name: "Underflow", args: args{a: -9223372036854775808, b: -1}, want: 9223372036854775807, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddInt64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddUint8(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name  string
		args  args
		want  uint8
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 255, b: 1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddUint8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddUint8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddUint16(t *testing.T) {
	type args struct {
		a uint16
		b uint16
	}
	tests := []struct {
		name  string
		args  args
		want  uint16
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 65535, b: 1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddUint16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddUint16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddUint32(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 4294967295, b: 1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddUint32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddUint32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 20}, want: 30, want1: false},
		{name: "Overflow", args: args{a: 18446744073709551615, b: 1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("AddUint64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddUint64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name  string
		args  args
		want  int8
		want1 bool
	}{
		{name: "No underflow", args: args{a: 10, b: 20}, want: -10, want1: false},
		{name: "Underflow", args: args{a: -128, b: 1}, want: 127, want1: true},
		{name: "Overflow", args: args{a: 127, b: -1}, want: -128, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubInt8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubInt8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubInt8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		{name: "No underflow", args: args{a: 10, b: 20}, want: -10, want1: false},
		{name: "Underflow", args: args{a: -32768, b: 1}, want: 32767, want1: true},
		{name: "Overflow", args: args{a: 32767, b: -1}, want: -32768, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubInt16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubInt16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		{name: "No underflow", args: args{a: 10, b: 20}, want: -10, want1: false},
		{name: "Underflow", args: args{a: -2147483648, b: 1}, want: 2147483647, want1: true},
		{name: "Overflow", args: args{a: 2147483647, b: -1}, want: -2147483648, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubInt32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{name: "No underflow", args: args{a: 10, b: 20}, want: -10, want1: false},
		{name: "Underflow", args: args{a: -9223372036854775808, b: 1}, want: 9223372036854775807, want1: true},
		{name: "Overflow", args: args{a: 9223372036854775807, b: -1}, want: -9223372036854775808, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubInt64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubUint8(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name  string
		args  args
		want  uint8
		want1 bool
	}{
		{name: "No underflow", args: args{a: 20, b: 10}, want: 10, want1: false},
		{name: "Underflow", args: args{a: 0, b: 1}, want: 255, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubUint8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubUint8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubUint8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubUint16(t *testing.T) {
	type args struct {
		a uint16
		b uint16
	}
	tests := []struct {
		name  string
		args  args
		want  uint16
		want1 bool
	}{
		{name: "No underflow", args: args{a: 20, b: 10}, want: 10, want1: false},
		{name: "Underflow", args: args{a: 0, b: 1}, want: 65535, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubUint16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubUint16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubUint16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubUint32(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{name: "No underflow", args: args{a: 20, b: 10}, want: 10, want1: false},
		{name: "Underflow", args: args{a: 0, b: 1}, want: 4294967295, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubUint32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubUint32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubUint32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		{name: "No underflow", args: args{a: 20, b: 10}, want: 10, want1: false},
		{name: "Underflow", args: args{a: 0, b: 1}, want: 18446744073709551615, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SubUint64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SubUint64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SubUint64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name  string
		args  args
		want  int8
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 127, b: 2}, want: -2, want1: true},
		{name: "Underflow", args: args{a: -128, b: 2}, want: 0, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 5}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulInt8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulInt8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulInt8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 32767, b: 2}, want: -2, want1: true},
		{name: "Underflow", args: args{a: -32768, b: 2}, want: 0, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 5}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulInt16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulInt16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 2147483647, b: 2}, want: -2, want1: true},
		{name: "Underflow", args: args{a: -2147483648, b: 2}, want: 0, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 5}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulInt32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 9223372036854775807, b: 2}, want: -2, want1: true},
		{name: "Underflow", args: args{a: -9223372036854775808, b: 2}, want: 0, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 5}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulInt64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulUint8(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name  string
		args  args
		want  uint8
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 255, b: 2}, want: 254, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 2}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulUint8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulUint8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulUint8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulUint16(t *testing.T) {
	type args struct {
		a uint16
		b uint16
	}
	tests := []struct {
		name  string
		args  args
		want  uint16
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 65535, b: 2}, want: 65534, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 2}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulUint16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulUint16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulUint16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulUint32(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 4294967295, b: 2}, want: 4294967294, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 2}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulUint32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulUint32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulUint32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMulUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 5}, want: 50, want1: false},
		{name: "Overflow", args: args{a: 18446744073709551615, b: 2}, want: 18446744073709551614, want1: true},
		{name: "Zero multiplication", args: args{a: 0, b: 2}, want: 0, want1: false},
		{name: "One multiplication", args: args{a: 1, b: 5}, want: 5, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MulUint64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("MulUint64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MulUint64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDivInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name  string
		args  args
		want  int8
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 2}, want: 5, want1: false},
		{name: "Division by zero", args: args{a: 10, b: 0}, want: 0, want1: true},
		{name: "Overflow", args: args{a: -128, b: -1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DivInt8(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("DivInt8() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DivInt8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDivInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 2}, want: 5, want1: false},
		{name: "Division by zero", args: args{a: 10, b: 0}, want: 0, want1: true},
		{name: "Overflow", args: args{a: -32768, b: -1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DivInt16(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("DivInt16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DivInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDivInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 2}, want: 5, want1: false},
		{name: "Division by zero", args: args{a: 10, b: 0}, want: 0, want1: true},
		{name: "Overflow", args: args{a: -2147483648, b: -1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DivInt32(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("DivInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DivInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDivInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{name: "No overflow", args: args{a: 10, b: 2}, want: 5, want1: false},
		{name: "Division by zero", args: args{a: 10, b: 0}, want: 0, want1: true},
		{name: "Overflow", args: args{a: -9223372036854775808, b: -1}, want: 0, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DivInt64(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("DivInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DivInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
