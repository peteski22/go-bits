package niable

import (
	"fmt"
	"time"
)

type NilableType interface {
	int | string | time.Duration
}

type Nilable[V NilableType] interface {
	value() func() (V, error)
	hasValue() func() bool
}

type Nil[V NilableType] struct {
	ptrToType *V
}

func (n *Nil[V]) value() (V, error) {
	if n.hasValue() {
		return *(n.ptrToType), nil
	}

	var v V
	return v, fmt.Errorf("invalid operation on nilable type, value not set")
}

func (n *Nil[V]) hasValue() bool {
	return n.ptrToType != nil
}

func (n *Nil[V]) set(v *V) {
	n.ptrToType = v
}

func NewNil[V NilableType]() *Nil[V] {
	return &Nil[V]{}
}

func doNilableStuff() {
	var p *int

	n := NewNil[int]()

	n.set(p)

	fmt.Println(n.hasValue())
	fmt.Println(n.value())

	p = new(int)
	*p = 12345
	n.set(p)

	fmt.Println(n.hasValue())
	fmt.Println(n.value())
}
