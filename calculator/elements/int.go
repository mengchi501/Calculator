package elements

type Integer struct {
	val int
}

func integer(val int) Integer {
	intn := Integer{val: val}
	return intn
}
