package utils

import (
	"flag"
)

func GetArgs() {
	flag.CommandLine.Arg(0)
}
