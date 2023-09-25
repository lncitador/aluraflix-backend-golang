package value_objects

type ValueObjectContract[Value any] interface {
	Equals(value ValueObjectContract[Value]) bool
	GetValue() Value
	ToString() string
}

type ValueObject[Value any] struct {
	value Value
}
