package main

import "math"

type MyString string

const (
	StringA MyString = "339"
	StringB MyString = "248"
)

func SomeString(arg MyString) {}

type MyInt int

const (
	IntA MyInt = 1238
	IntB MyInt = 8884432
)

func SomeInt(arg MyInt) {}

type MyFloat float32

const (
	FloatA MyFloat = 2.138
	FloatB MyFloat = math.Pi
)

func SomeFloat(arg MyFloat) {}

func main() {
	// string
	SomeString("339")   // want "^raw literal \\(STRING\\) passed to type alias \\(main\\.MyString\\), use a constant instead$"
	SomeString(StringA) // OK
	SomeString(StringB) // OK

	// int
	SomeInt(1238) // want "^raw literal \\(INT\\) passed to type alias \\(main\\.MyInt\\), use a constant instead$"
	SomeInt(IntA) // OK
	SomeInt(IntB) // OK

	// float
	SomeFloat(2.138)  // want "^raw literal \\(FLOAT\\) passed to type alias \\(main\\.MyFloat\\), use a constant instead$"
	SomeFloat(FloatA) // OK
	SomeFloat(FloatB) // OK
}
