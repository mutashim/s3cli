package parser

import "strings"

func ParsePath(arg string) (protocol string, bucket string, path string) {

	// initial value
	protocol = "local"
	bucket = ""
	path = ""

	argComp := strings.Split(arg, "://")
	if len(argComp) < 2 {
		path = argComp[0]
		return
	}

	protocol = argComp[0]
	paths := strings.SplitN(argComp[1], "/", 2)

	bucket = paths[0]
	path = paths[1]
	return
}
