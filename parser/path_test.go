package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mutashim/s3cli/parser"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	pro, buck, path := parser.ParsePath("s3://bucket/dir/file.ext")
	assert.Equal(t, "s3", pro)
	assert.Equal(t, "bucket", buck)
	assert.Equal(t, "dir/file.ext", path)
}

func TestSpl(t *testing.T) {
	fmt.Println(strings.Split("google", "://"))
}
