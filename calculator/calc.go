package calculator

import (
	"fmt"
	"strconv"

	e "./elements"
)

type Calc struct {
	stack *e.Stack
}

func NewCalc() *Calc {
	calc := new(Calc)
	calc.stack = e.NewStack()
	return calc
}

func (c *Calc) ShowResult() {
	res := c.getResult()
	real := float64(res.I) + res.F
	reals := strconv.FormatFloat(real, 'f', -1, 64)
	if real != 0 {
		fmt.Printf("%s ", reals)
	}
}

func (c *Calc) getResult() e.Number {
	res := c.stack.Pop().(e.Number)
	return res
}

func (c *Calc) Enter(i interface{}) {
	if temp, ok := i.(string); ok {
		num2 := c.stack.Pop().(e.Number)
		num1 := c.stack.Pop().(e.Number)
		var res *e.Number
		if temp == "+" {
			res = num1.Add(&num2)
		} else if temp == "-" {
			res = num1.Sub(&num2)
		} else if temp == "*" {
			res = num1.Mul(&num2)
		} else {
			res = num1.Div(&num2)
		}
		c.stack.Push(*res)
	} else {
		c.stack.Push(*e.NewNumber(i))
	}

}
