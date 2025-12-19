package utils

import (
	"log"
	"os"
)

func ReadEnvVar(s string) string {

	envVar, found := os.LookupEnv(s)

	if !found {
		log.Fatal("variable not found!!!")
	}

	return envVar
}
