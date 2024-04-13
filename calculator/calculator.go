package calculator

import "errors"

func Add(a, b int) int {
	return a + b
}

func DivideBy(a, b int) (int, error) {
	if b > a {
		return 0, errors.New("too big number")
	}

	return a / b, nil
}
