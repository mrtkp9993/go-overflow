# go-overflow

`go-overflow` is a lightweight Go library designed to handle integer arithmetic operations while checking for overflow and underflow conditions. It provides utility functions for addition, subtraction, multiplication, and division for various integer types (`int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`).

## Features

- **Overflow Detection**: Detects overflow during addition, subtraction, and multiplication.
- **Underflow Detection**: Detects underflow during subtraction.
- **Division Safety**: Handles division by zero and overflow during division.
- Supports both signed and unsigned integer types.

## Installation

To use `go-overflow` in your project, simply run:

```bash
go get github.com/mrtkp9993/go-overflow
```

## Usage

Import the library in your Go project:

```go
import "github.com/mrtkp9993/go-overflow"
```

### Examples

#### Addition with Overflow Check

```go
result, overflow := overflow.AddInt8(120, 10)
if overflow {
    fmt.Println("Overflow occurred!")
} else {
    fmt.Println("Result:", result)
}
```

#### Subtraction with Underflow Check

```go
result, underflow := overflow.SubInt16(10, 20)
if underflow {
    fmt.Println("Underflow occurred!")
} else {
    fmt.Println("Result:", result)
}
```

#### Multiplication with Overflow Check

```go
result, overflow := overflow.MulInt32(100000, 100000)
if overflow {
    fmt.Println("Overflow occurred!")
} else {
    fmt.Println("Result:", result)
}
```

#### Division with Safety Check

```go
result, errorOccurred := overflow.DivInt64(-9223372036854775808, -1)
if errorOccurred {
    fmt.Println("Error occurred during division!")
} else {
    fmt.Println("Result:", result)
}
```

## Supported Functions

### Signed Integers

- `AddInt8`, `AddInt16`, `AddInt32`, `AddInt64`
- `SubInt8`, `SubInt16`, `SubInt32`, `SubInt64`
- `MulInt8`, `MulInt16`, `MulInt32`, `MulInt64`
- `DivInt8`, `DivInt16`, `DivInt32`, `DivInt64`

### Unsigned Integers

- `AddUint8`, `AddUint16`, `AddUint32`, `AddUint64`
- `SubUint8`, `SubUint16`, `SubUint32`, `SubUint64`
- `MulUint8`, `MulUint16`, `MulUint32`, `MulUint64`

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the library.

## License

This project is licensed under the  GNU GPLv3 License. See the `LICENSE` file for details.
