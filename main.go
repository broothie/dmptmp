package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	//go:embed usage.txt
	usage string

	//go:embed VERSION
	version string

	logger = log.New(os.Stdout, "[dmptmp] ", 0)

	dir     = flag.String("dir", "", "Directory in which to create temp file")
	pattern = flag.String("pattern", "", "Pattern for constructing filename")
)

func init() {
	flagUsage := flag.Usage
	flag.Usage = func() {
		fmt.Printf("dmptmp %s\n\n", strings.TrimSpace(version))
		fmt.Println(usage)
		flagUsage()
	}

	flag.Parse()
}

func main() {
	file, err := os.CreateTemp(*dir, *pattern)
	if err != nil {
		logger.Fatalln("failed to create temp file", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Fatalln("failed to close temp file", err)
		}
	}()

	fmt.Println(file.Name())
	if _, err := io.Copy(file, os.Stdin); err != nil {
		logger.Fatalln("failed to write temp file", err)
	}
}
