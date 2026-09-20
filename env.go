package env

import (
	"fmt"
	"os"
)

type parseFunc[T any] func(s string) (T, error)

func Require[T any](key string, parseFn parseFunc[T]) T {
	strValue, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("environment variable %q must be set", key))
	}

	value, err := parseFn(strValue)
	if err != nil {
		panic(fmt.Errorf("could not parse %q: %w", key, err))
	}

	return value
}

func Fallback[T any](key string, fallback T, parseFn parseFunc[T]) (T, error) {
	strValue, ok := os.LookupEnv(key)
	if !ok {
		return fallback, nil
	}

	return parseFn(strValue)
}
