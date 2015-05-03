package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	file string
	host string
}

func NewParser(file, host string) *Parser {
	return &Parser{file: file, host: host}
}

func (p *Parser) Parse() map[string]map[string]string {
	var seen bool
	ret := make(map[string]map[string]string)

	file, err := os.Open(p.file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		keyword := fields[0]
		arguments := fields[1:]

		if keyword == "Host" {
			if hostlineIncludes(arguments, p.host) {
				seen = true
				continue
			} else {
				seen = false
			}
		}

		if seen {
			if ret[p.host] == nil {
				ret[p.host] = make(map[string]string)
			}
			ret[p.host][keyword] = strings.Join(arguments, " ")
		}
	}

	return ret
}

func hostlineIncludes(hostline []string, host string) bool {
	for _, field := range hostline {
		replaced := strings.Replace(field, "*", ".*", 1)
		pattern := regexp.MustCompile(replaced)
		if pattern.MatchString(host) {
			return true
		}
	}

	return false
}
