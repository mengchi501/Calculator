package main

import (
	c "./calculator"
)

func main() {
	calc := c.NewCalc()

	// 3.2 * (2 + 5) - (1.25 + 1.25) * 8
	calc.Enter(2)
	calc.Enter(5)
	calc.Enter("+")
	calc.Enter(3.2)
	calc.Enter("*")
	calc.Enter(1.25)
	calc.Enter(1.25)
	calc.Enter("+")
	calc.Enter(8)
	calc.Enter("*")
	calc.Enter("-")
	calc.ShowResult()
}
