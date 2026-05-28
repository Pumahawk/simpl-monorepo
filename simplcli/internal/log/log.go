package log

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var DebugEnabled = false

func init() {
	if v, ok := os.LookupEnv("SIMPL_CLI_DEBUG"); ok {
		if v, err := strconv.ParseBool(v); err != nil {
			fmt.Fprintf(os.Stderr, "invalid debug value configuration, set log debug to false: %s\n", err)
		} else {
			DebugEnabled = v
		}
	}
}

func Debug(fmts string, v ...any) {
	if DebugEnabled {
		log.Printf(fmts, v...)
	}
}
