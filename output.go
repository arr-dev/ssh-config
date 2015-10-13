package main

import (
	"encoding/json"
	"fmt"
)

type Output struct {
	result   Result
	hostOnly bool
	format   string
}

func NewOutput(result Result, hostOnly bool, format string) *Output {
	return &Output{result: result, hostOnly: hostOnly, format: format}
}

func (o *Output) Format() string {

	if o.hostOnly {
		return o.result.Options["Hostname"]
	}

	switch o.format {
	case "plain", "p":
		return plainFormat(o.result)
	case "pretty":
		return prettyFormat(o.result)
	case "json", "j":
		ret, err := json.Marshal(o.result)
		if err == nil {
			return string(ret)
		}
	}

	return ""
}

func plainFormat(r Result) string {
	ret := ""
	ret += fmt.Sprintf("Host %s\n", r.Host)
	for _, k := range r.OptionKeys {
		ret += fmt.Sprintf("    %s %v\n", k, r.Options[k])
	}
	return ret
}

func prettyFormat(r Result) string {
	maxLength := 0
	for k, _ := range r.Options {
		if l := len(k); l > maxLength {
			maxLength = l
		}
	}
	headerTemplate := fmt.Sprintf("%%-%ds %%s\n", maxLength+4)
	template := fmt.Sprintf("    %%-%ds %%s\n", maxLength)

	ret := fmt.Sprintf(headerTemplate, "Host", r.Host)

	for _, k := range r.OptionKeys {
		ret += fmt.Sprintf(template, k, r.Options[k])
	}
	return ret
}
