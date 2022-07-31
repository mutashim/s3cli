package parser_test

import (
	"os"
	"testing"

	"github.com/mutashim/s3cli/parser"
	"github.com/stretchr/testify/assert"
)

func TestArg(t *testing.T) {
	os.Args = append(os.Args, "arg1 arg2 arg3 arg4")
	ar1, ar2, ar3, ar4 := parser.ParseArg()
	assert.Equal(t, "arg1", ar1)
	assert.Equal(t, "arg2", ar2)
	assert.Equal(t, "arg3", ar3)
	assert.Equal(t, "arg4", ar4)
}
