package env

import "strconv"

type parseFunc[T any] func(s string) (T, error)

func StringParser() parseFunc[string] {
	return func(s string) (string, error) {
		return s, nil
	}
}

func IntParser() parseFunc[int] {
	return func(s string) (int, error) {
		return strconv.Atoi(s)
	}
}

func BoolParser() parseFunc[bool] {
	return func(s string) (bool, error) {
		return strconv.ParseBool(s)
	}
}
