package rop_test

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/sivchari/rop"
)

func Example() {
	result := Then(
		Then(
			Then(
				Pipe("42"),
				parseint,
			),
			divideByTwo,
		),
		stringify,
	)

	val, err := result.Match()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
	// Output:
	// 21
}

func parseint(in string) Result[int] {
	if in == "" {
		return Err[int](errors.New("empty input"))
	}
	i, err := strconv.Atoi(in)
	if err != nil {
		return Err[int](err)
	}
	return OK(i)
}

func divideByTwo(in int) Result[int] {
	if in == 0 {
		return Err[int](errors.New("cannot divide by zero"))
	}
	return OK(in / 2)
}

func stringify(in int) Result[string] {
	return OK(strconv.Itoa(in))
}
