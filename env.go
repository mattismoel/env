package env

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ErrInvalidKey   = errors.New("invalid env key")
	ErrInvalidValue = errors.New("invalid value")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no environment variable file(s) was found")
	}
}

// Returns an environment varible of string type.
// If the key does not exist, the fallback value is returned.
func Str(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

// Returns an environment variable of int type.
// If the key does not exist, the fallback value is returned.
func Int(key string, fallback int) int {
	val := Str(key, "")
	if val == "" {
		return fallback
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

// Returns an environment variable of boolean type.
// If the key does not exist,  or the value is not "true" or "false",
// the fallback value is returned.
//
// The function is case-insensitive, meaning that "True" and "False" are valid.
func Bool(key string, fallback bool) bool {
	val := Str(key, "")
	if val == "" {
		return fallback
	}

	switch strings.ToLower(val) {
	case "true":
		return true
	case "false":
		return false
	}

	return fallback
}

// Returns an environment variable of floating point type.
// If the key does not exist, the fallback value is returned.
func Float32(key string, fallback float32) float32 {
	val := Str(key, "")
	if val == "" {
		return fallback
	}

	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return fallback
	}

	return float32(f)
}

// Sets an environment variable of string type.
// If not able to set, an error is returned.
//
// Internally, only the os.Setenv is called.
func SetStr(key string, value string) error {
	// Validation fo the key and value.
	switch {
	case strings.ContainsRune(key, '"'):
		return ErrInvalidKey
	case strings.ContainsRune(value, '"'):
		return ErrInvalidValue
	case key == "":
		return ErrInvalidKey
	}

	return os.Setenv(key, value)
}

// Sets an environment variable of integer type.
// If not able to set, an error is returned.
func SetInt(key string, value int) error {
	i := strconv.Itoa(value)
	return SetStr(key, i)
}

// Sets an environment vaiable of boolean type.
// If not able to set, an error is returned.
func SetBool(key string, value bool) error {
	bStr := ""
	switch value {
	case true:
		bStr = "true"
	case false:
		bStr = "false"
	}

	return SetStr(key, bStr)
}

// Sets an environment variable of floating point type.
// The prec specifies the amount of digits after the decimal point.
//
// If not able to set, an error is returned.
func SetFloat32(key string, value float32, prec ...int) error {
	p := 2
	switch {
	case len(prec) > 1:
		p = prec[0]
	}

	fStr := strconv.FormatFloat(float64(value), byte('f'), p, 32)

	return SetStr(key, fStr)
}
