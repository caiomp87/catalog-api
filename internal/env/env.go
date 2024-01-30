package env

import (
	"os"
	"strconv"
	"strings"
)

type Env string

func GetEnv(key string) Env {
	value := os.Getenv(key)
	return Env(value)
}

func (e Env) String() string {
	return string(e)
}

func (e Env) Bool() bool {
	return strings.ToLower(e.String()) == "true"
}

func (e Env) FallbackString(fallback string) string {
	if e == "" {
		return fallback
	}
	return e.String()
}

func (e Env) FallbackInt(fallback int) int {
	strFallback := strconv.Itoa(fallback)
	parsedValue, err := strconv.Atoi(e.FallbackString(strFallback))
	if err != nil {
		return fallback
	}
	return parsedValue
}

func (e Env) FallbackInt32(fallback int32) int32 {
	parsedValue, err := strconv.ParseInt(e.FallbackString(strconv.Itoa(int(fallback))), 10, 32)
	if err != nil {
		return fallback
	}
	return int32(parsedValue)
}

func (e Env) FallbackFloat64(fallback float64) float64 {
	strFallback := strconv.FormatFloat(fallback, 'f', -1, 64)
	parsedValue, err := strconv.ParseFloat(e.FallbackString(strFallback), 64)
	if err != nil {
		return fallback
	}
	return parsedValue
}
