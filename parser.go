package main

import (
	"bufio"
	"log"
	"os"
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

		if fields[0] == "Host" {
			if hostlineIncludes(fields, p.host) {
				seen = true
				ret[p.host] = make(map[string]string)
				continue
			} else if seen {
				break
			}
		}

		if seen {
			ret[p.host][fields[0]] = strings.Join(fields[1:], " ")
		}
	}

	return ret
}

func hostlineIncludes(hostline []string, s string) bool {
	for _, x := range hostline {
		if x == s {
			return true
		}
	}

	return false
}
