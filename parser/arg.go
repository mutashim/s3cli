package parser

import (
	"log"
	"os"
)

func ParseArg() (cmd string, arg1 string, arg2 string, etc string) {
	args := os.Args
	cmd = ""
	arg1 = ""
	arg2 = ""
	etc = ""

	switch len(args) {
	case 5:
		etc = args[4]
		arg2 = args[3]
		arg1 = args[2]
		cmd = args[1]
	case 4:
		arg2 = args[3]
		arg1 = args[2]
		cmd = args[1]
	case 3:
		arg1 = args[2]
		cmd = args[1]
	case 2:
		cmd = args[1]
	case 1:
		log.Fatal("args required")
	}
	return
}
