package elements

import "math"

type Number struct {
	I int
	F float64
}

func NewNumber(val interface{}) *Number {
	number := new(Number)
	switch temp := val.(type) {
	case int:
		number.I = temp
		number.F = 0.0
	case float64:
		it, ft := math.Modf(temp)
		number.I = int(it)
		number.F = ft
	case complex128:
		it, ft := math.Modf(float64(real(temp)))
		number.I = int(it)
		number.F = ft
	}
	return number
}

func (n1 *Number) Add(n2 *Number) *Number {
	tempi := n1.I + n2.I
	tempf := n1.F + n2.F
	res := Number{I: tempi, F: tempf}
	return &res
}

func (n1 *Number) Sub(n2 *Number) *Number {
	tempi := n1.I - n2.I
	tempf := n1.F - n2.F
	res := Number{I: tempi, F: tempf}
	return &res
}

func (n1 *Number) Mul(n2 *Number) *Number {
	s1 := float64(n1.I) + n1.F
	s2 := float64(n2.I) + n2.F
	res := NewNumber(s1 * s2)
	return res
}

func (n1 *Number) Div(n2 *Number) *Number {
	s1 := float64(n1.I) + n1.F
	s2 := float64(n2.I) + n2.F
	res := NewNumber(s1 / s2)
	return res
}
