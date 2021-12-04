package test

import (
	"fmt"
	"os"
	"strconv"
)

func SetTesting() {
	os.Setenv(getTestEnvName(), "true")
}

func IsTest() bool {
	envName := getTestEnvName()
	value, ok := os.LookupEnv(envName)
	if ok {
		result, err := strconv.ParseBool(value)
		if err != nil {
			return false
		}

		return result
	}

	return false
}

func getTestEnvName() string {
	envPrefix := "PROJECT"
	if value, ok := os.LookupEnv("PREFIX"); ok {
		envPrefix = value
	}
	return fmt.Sprintf("%s_TEST", envPrefix)
}
