package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Parser struct {
	file string
	host string
}

type Result struct {
	Host       string
	Options    map[string]string
	OptionKeys []string `json:"-"`
}

func NewParser(file, host string) *Parser {
	return &Parser{file: file, host: host}
}

func (p *Parser) Parse() Result {
	var seen bool
	ret := Result{}

	file, err := os.Open(p.file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := cleanLine(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		keyword := fields[0]
		arguments := fields[1:]

		if keyword == "Host" {
			if hostlineIncludes(arguments, p.host) {
				seen = true
				ret.Host = p.host
				continue
			} else {
				seen = false
			}
		}

		if seen {
			if ret.Options == nil {
				ret.Options = make(map[string]string)
			}
			ret.Options[keyword] = strings.Join(arguments, " ")
			ret.OptionKeys = append(ret.OptionKeys, keyword)
		}
	}

	sort.Strings(ret.OptionKeys)

	return ret
}

func hostlineIncludes(hostline []string, host string) bool {
	for _, field := range hostline {
		replaced := strings.Replace(field, "*", ".*", 1)
		pattern := regexp.MustCompile("^" + replaced + "$")
		if pattern.MatchString(host) {
			return true
		}
	}

	return false
}

func cleanLine(line string) string {
	i := strings.Index(strings.TrimSpace(line), "#")
	if i < 0 {
		return line
	} else {
		return line[0:i]
	}
}
