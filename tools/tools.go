package tools

func Ternary(expr bool, first interface{}, second interface{}) interface{} {
	if expr {
		return first
	}

	return second
}
