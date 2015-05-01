package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	var host string
	var file string

	user, _ := user.Current()
	defaultFile := fmt.Sprintf("%s/.ssh/config", user.HomeDir)

	flag.StringVar(&file, "f", defaultFile, "SSH config file to parse")
	flag.StringVar(&file, "file", defaultFile, "SSH config file to parse")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "%s [options] HOSTNAME\n\n", filepath.Base(os.Args[0]))
		fmt.Fprint(os.Stderr, "Options:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	host = flag.Arg(0)

	p := NewParser(file, host)
	config := p.Parse()

	fmt.Println(config)
}
