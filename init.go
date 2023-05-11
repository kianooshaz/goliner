package goliner

var Epsilon float64

func init() {
	e := 1.0
	for 1.0+e/2 != 1.0 {
		e /= 2
	}
	Epsilon = e
}
