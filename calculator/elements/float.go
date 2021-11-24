package elements

type Float struct {
	val float64
}

func float(val float64) Float {
	fn := Float{val: val}
	return fn
}
