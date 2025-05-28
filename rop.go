package rop

// Result is a generic type that represents a computation that can either succeed with a value of type T or fail with an error.
// T: the type of the successful result
// Err: the type of the error that can occur
type Result[T any] struct {
	Value T
	Err   error
}

// Unwrap returns the value if the Result is successful, or panics if there is an error.
func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Value
}

// Match returns the value if the Result is successful, or the error if there is an error.
func (r Result[T]) Match() (value T, err error) {
	if r.Err != nil {
		return value, r.Err
	}
	return r.Value, nil
}

// OK returns a Result with a successful value.
func OK[T any](v T) Result[T] {
	return Result[T]{Value: v}
}

// Err returns a Result with an error.
func Err[T any](e error) Result[T] {
	return Result[T]{Err: e}
}

// Then is a function that takes a Result and a function that operates on the successful value.
func Then[T any, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.Err != nil {
		return Err[U](r.Err)
	}
	return f(r.Value)
}

// Pipe is a function that wraps a Result at first argument and a function that operates on the successful value as the second argument.
func Pipe[T any](in T) Result[T] {
	return OK(in)
}
