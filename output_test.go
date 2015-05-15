package main

import (
	"encoding/json"
	"strings"
	"testing"
)

var result = Result{
	Host: "github.com",
	Options: map[string]string{
		"Hostname":                 "github.com",
		"PreferredAuthentications": "publickey",
		"ServerAliveInterval":      "60",
	}}

func TestOutputPlain(t *testing.T) {

	hostOnly := false
	format := "plain"

	o := NewOutput(result, hostOnly, format)
	got := o.Format()
	if !strings.Contains(got, "map") && !strings.Contains(got, "Hostname:github.com") {
		t.Error("Expected ", got, "to include Hostname")
	}
}

func TestOutputHostOnly(t *testing.T) {

	hostOnly := true
	format := "plain"

	o := NewOutput(result, hostOnly, format)
	got := o.Format()
	if got != result.Options["Hostname"] {
		t.Error("Expected ", result.Options["Hostname"], "Got ", got)
	}
}

func TestOutputJson(t *testing.T) {

	hostOnly := false
	format := "json"

	o := NewOutput(result, hostOnly, format)
	got := o.Format()
	if j, _ := json.Marshal(result); string(j) != got {
		t.Error("Expected JSON", "Got ", got)
	}
}
