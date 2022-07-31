package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mutashim/s3cli/parser"
	"github.com/mutashim/s3go"
)

func LoadKeys() (accessKey string, secretKey string) {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("cannot get home directory:", err.Error())
	}

	data, err := os.ReadFile(filepath.Join(dir, ".s3/conf"))
	if err != nil {
		if err != nil {
			log.Fatal("cannot read config:", err.Error())
		}
	}
	conf := strings.Split(string(data[:]), ":")

	accessKey = conf[0]
	secretKey = conf[1]
	return
}

func Run(s3client s3go.S3, cmd string, arg1 string, arg2 string, opt string) error {

	a := &app{client: s3client}

	switch cmd {
	case "cp":
		return a.cp(arg1, arg2, opt)
	case "mv":
		return errors.New("feature not available")
	case "rm":
		return a.rm(arg1)
	case "ls":
		return a.ls(arg1)
	case "chmod":
		return a.chmod(arg1, arg2)
	case "link":
		if arg2 == "" {
			return errors.New("additional param required")
		}

		d, err := strconv.Atoi(arg2)
		if err != nil {
			return err
		}
		return a.link(arg1, d)
	case "mkbuck":
		return a.mkbuck(arg1)
	case "lsbuck":
		return a.lsbucket()
	default:
		return errors.New("command not recognized")
	}
}

type App interface {
	chmod(origin string, option string) error
	cp(path1, path2, option string) error
	ls(origin string) error
	lsbucket() error
	rm(input string) error
	rmbucket(name string) error
	mkbuck(name string) error
	link(origin string, duration int) error
}

type app struct {
	client s3go.S3
}

// Copy
func (a *app) cp(pathA string, pathB string, option string) error {
	protA, bucketA, absPathA := parser.ParsePath(pathA)
	protB, bucketB, absPathB := parser.ParsePath(pathB)

	if protA == protB {
		return errors.New("origin path and target path are from same drive")
	}

	// action: download
	// origin: s3
	// target: local
	if protA == "s3" {
		a.client.SetBucket(bucketA)
		err := a.client.Download(absPathA, absPathB)
		if err != nil {
			return err
		}

		fmt.Println("file has been downloaded to ", absPathB)
		return nil
	}

	// action: upload
	// origin: local
	// target: s3
	if protA == "local" {
		a.client.SetBucket(bucketB)
		acl := s3go.PRIVATE
		if option == "public" {
			acl = s3go.PUBLIC
		}
		err := a.client.Upload(absPathA, absPathB, acl)
		if err != nil {
			return err
		}

		fmt.Printf("file has been uploaded to s3://%s/%s\n", bucketB, absPathB)
		return nil
	}

	return errors.New("no action")
}

// List bucket
func (a *app) lsbucket() error {
	result, space, err := a.client.ListBucket()
	if err != nil {
		return err
	}

	for _, v := range *result {
		fmt.Printf("%s %s %s \n", v.Name, strings.Repeat(" ", space-len(v.Name)), v.CreatedAt)
	}

	return nil
}

// List file
func (a *app) ls(path string) error {
	prot, bucket, pathParsed := parser.ParsePath(path)
	if prot != "s3" {
		return errors.New("given path not s3")
	}

	// fetch
	a.client.SetBucket(bucket)
	result, space, err := a.client.ListFile(pathParsed)
	if err != nil {
		return err
	}

	// show
	for _, v := range *result {
		fmt.Printf("%s %s %s\n", v.Name, strings.Repeat(" ", space-len(v.Name)), v.LastMod)
	}

	return nil
}

// Remove file
func (a *app) rm(origin string) error {
	prot, bucket, path := parser.ParsePath(origin)
	if prot != "s3" {
		return errors.New("origin path not from s3")
	}

	a.client.SetBucket(bucket)
	err := a.client.Delete(path)
	if err != nil {
		return err
	}

	fmt.Println("file has been deleted")
	return nil
}

// change modifier
func (a *app) chmod(origin string, option string) error {
	prot, bucket, path := parser.ParsePath(origin)
	if prot != "s3" {
		return errors.New("origin path not from s3")
	}

	a.client.SetBucket(bucket)
	acl := s3go.PRIVATE
	if option == "public" {
		acl = s3go.PUBLIC
	}

	err := a.client.SetAcl(path, acl)
	if err != nil {
		return err
	}

	fmt.Println("file has been set to:", option)
	return nil
}

// make bucket
func (a *app) mkbuck(name string) error {
	return a.client.MakeBucket(name)
}

// remove bucket
func (a *app) rmbucket(name string) error {
	err := a.client.RemoveBucket(name)
	if err != nil {
		return err
	}

	fmt.Println("success deleted: " + name)
	return nil
}

// get link
func (a *app) link(origin string, duration int) error {

	// validation
	if origin == "" {
		return errors.New("origin path required")
	}
	if duration == 0 {
		return errors.New("duration parameter required")
	}

	prot, bucket, path := parser.ParsePath(origin)
	if prot != "s3" {
		return errors.New("origin path not from s3")
	}

	a.client.SetBucket(bucket)
	link, err := a.client.Share(path, int64(duration))
	if err != nil {
		return err
	}

	// print url
	fmt.Println("link of file has been generated:")
	fmt.Printf("File:     %s\n", path)
	fmt.Printf("Bucket:   %s\n", bucket)
	fmt.Printf("duration: %s minute(s)\n", strconv.Itoa(duration))
	fmt.Printf("Link: \n%s\n", link)
	return nil
}
