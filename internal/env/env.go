package env

import (
	"os"
	"strconv"
)

func GetEnvVar[T int | string](key string, fallback T) T {
	envVarValue, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	switch any(fallback).(type) {
	case int:
		intVal, err := strconv.Atoi(envVarValue)
		if err != nil {
			return fallback
		}
		return any(intVal).(T)
	case string:
		return any(envVarValue).(T)
	}

	return fallback
}
