package main

import (
	"fmt"
	"log"

	"github.com/mutashim/s3cli/app"
	"github.com/mutashim/s3cli/parser"
	"github.com/mutashim/s3go"
)

const (
	ACCESS_KEY = ""
)

func main() {

	// load config
	cfg := LoadConfig()

	cmd, arg1, arg2, etc := parser.ParseArg()

	s3client, err := s3go.New(&s3go.Config{
		AccessKey: cfg.AccessKey,
		SecretKey: cfg.SecretKey,
		Endpoint:  cfg.Endpoint,
		Region:    cfg.Region,
	})

	if err != nil {
		log.Fatalf("Cannot create s3 client: %s", err.Error())
	}
	defer fmt.Println()

	if err := app.Run(s3client, cmd, arg1, arg2, etc); err != nil {
		fmt.Println("ERR:", err)
	}

}
