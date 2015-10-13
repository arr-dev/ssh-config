package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	p := NewParser("testdata/ssh.conf", "github.com")
	got := p.Parse()
	expected := Result{
		Host: "github.com",
		Options: map[string]string{
			"Hostname":                 "github.com",
			"PreferredAuthentications": "publickey",
			"ServerAliveInterval":      "60",
		},
		OptionKeys: []string{"Hostname", "PreferredAuthentications", "ServerAliveInterval"},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}

func TestParserHostWildCard(t *testing.T) {
	p := NewParser("testdata/ssh.conf", "example.com")
	got := p.Parse()
	expected := Result{
		Host: "example.com",
		Options: map[string]string{
			"Hostname":            "ssh.example.com",
			"ServerAliveInterval": "60",
			"Username":            "test",
		},
		OptionKeys: []string{"Hostname", "ServerAliveInterval", "Username"},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}
