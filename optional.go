package optional

type Optional[T any] struct {
	Value   T
	IsSet   bool
	Default T
}

func NilValue[T any]() Optional[T] {
	return Optional[T]{IsSet: false}
}

// Function to create an optional value
func SetValue[T any](v T) Optional[T] {
	return Optional[T]{Value: v, IsSet: true}
}

// Function to create an optional value with a default value
func SetValueWithDefault[T any](v T, defaultValue T) Optional[T] {
	return Optional[T]{Value: v, IsSet: true, Default: defaultValue}
}

// Function to get the value of an optional value
func (o Optional[T]) GetValue() T {
	if !o.IsSet {
		return o.Default
	}
	return o.Value
}

func (o Optional[T]) ValueOrDefault(defaultValue T) T {
	if o.IsSet {
		return o.Value
	}
	return defaultValue
}

func (o Optional[T]) IsEmpty() bool {
	return IsEmpty(o.Value)
}

func (o Optional[T]) IsZero() bool {
	return IsZero(o.Value)
}
