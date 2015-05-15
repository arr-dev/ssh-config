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
	case "plain":
		return fmt.Sprintf("%v", o.result)
	case "json":
		ret, err := json.Marshal(o.result)
		if err == nil {
			return string(ret)
		}
	}

	return ""
}
