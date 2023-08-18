package option

type Option[T any] struct {
	value  T
	isNone bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value, false}
}

func None[T any]() Option[T] {
	return Option[T]{isNone: true}
}

func (option Option[T]) IsSome() bool {
	return !option.isNone
}

func (option Option[T]) IsNone() bool {
	return option.isNone
}

func (option Option[T]) Unwrap() T {
	return option.value
}

func (option Option[T]) UnwrapOr(callback func() T) T {
	if option.isNone {
		return callback()
	}

	return option.value
}
