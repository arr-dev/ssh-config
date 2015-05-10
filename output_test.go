package main

import (
	"encoding/json"
	"strings"
	"testing"
)

var config = map[string]string{
	"Hostname":                 "github.com",
	"PreferredAuthentications": "publickey",
	"ServerAliveInterval":      "60",
}

func TestOutputPlain(t *testing.T) {

	hostOnly := false
	format := "plain"

	o := NewOutput(config, hostOnly, format)
	got := o.Format()
	if !strings.Contains(got, "map") && !strings.Contains(got, "Hostname:github.com") {
		t.Error("Expected ", got, "to include Hostname")
	}
}

func TestOutputHostOnly(t *testing.T) {

	hostOnly := true
	format := "plain"

	o := NewOutput(config, hostOnly, format)
	got := o.Format()
	if got != config["Hostname"] {
		t.Error("Expected ", config["Hostname"], "Got ", got)
	}
}

func TestOutputJson(t *testing.T) {

	hostOnly := false
	format := "json"

	o := NewOutput(config, hostOnly, format)
	got := o.Format()
	if j, _ := json.Marshal(config); string(j) != got {
		t.Error("Expected JSON", "Got ", got)
	}
}
