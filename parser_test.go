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
			"ServerAliveInterval":      "60",
			"Hostname":                 "github.com",
			"PreferredAuthentications": "publickey",
		},
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
			"ServerAliveInterval": "60",
			"Hostname":            "ssh.example.com",
			"Username":            "test",
		},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}

func TestParserHostIgnoresWildCardAnywhere(t *testing.T) {
	p := NewParser("testdata/ssh.conf", "another-example")
	got := p.Parse()
	expected := Result{
		Host: "another-example",
		Options: map[string]string{
			"ServerAliveInterval": "60",
			"Username":            "other",
		},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}
