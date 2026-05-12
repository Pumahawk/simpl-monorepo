package utils

import "os"

func EnvOrDef(key, def string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	} else {
		return def
	}
}
