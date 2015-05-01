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
	var host string
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
			if fields[1] == p.host {
				seen = true
				host = strings.Join(fields[1:], " ")
				ret[host] = make(map[string]string)
				continue
			} else if seen {
				break
			}
		}

		if seen {
			ret[host][fields[0]] = strings.Join(fields[1:], " ")
		}
	}

	return ret
}
