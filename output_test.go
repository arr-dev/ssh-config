package main

import (
	"testing"
)

var result = Result{
	Host: "github.com",
	Options: map[string]string{
		"Hostname":                 "github.com",
		"PreferredAuthentications": "publickey",
		"ServerAliveInterval":      "60",
	},
	OptionKeys: []string{"Hostname", "PreferredAuthentications", "ServerAliveInterval"},
}

func TestOutputPlain(t *testing.T) {

	hostOnly := false
	format := "plain"

	o := NewOutput(result, hostOnly, format)
	got := o.Format()
	expected :=
		`Host github.com
    Hostname github.com
    PreferredAuthentications publickey
    ServerAliveInterval 60
`
	if got != expected {
		t.Error("Expected ", got, "to match ", expected)
	}
}

func TestOutputPretty(t *testing.T) {

	hostOnly := false
	format := "pretty"

	o := NewOutput(result, hostOnly, format)
	got := o.Format()
	expected :=
		`Host                         github.com
    Hostname                 github.com
    PreferredAuthentications publickey
    ServerAliveInterval      60
`
	if got != expected {
		t.Error("Expected ", got, "to match ", expected)
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
	expected := `{"Host":"github.com","Options":{"Hostname":"github.com","PreferredAuthentications":"publickey","ServerAliveInterval":"60"}}`
	if got != expected {
		t.Error("Expected ", expected, "Got ", got)
	}
}
