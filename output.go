package main

import (
	"encoding/json"
	"fmt"
)

type Output struct {
	config   map[string]string
	hostOnly bool
	format   string
}

func NewOutput(config map[string]string, hostOnly bool, format string) *Output {
	return &Output{config: config, hostOnly: hostOnly, format: format}
}

func (o *Output) Format() string {

	if o.hostOnly {
		return o.config["Hostname"]
	}

	switch o.format {
	case "plain":
		return fmt.Sprintf("%v", o.config)
	case "json":
		ret, err := json.Marshal(o.config)
		if err == nil {
			return string(ret)
		}
	}

	return ""
}
