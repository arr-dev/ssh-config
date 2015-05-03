package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	p := NewParser("testdata/ssh.conf", "github.com")
	got := p.Parse()
	expected := map[string]map[string]string{
		"github.com": map[string]string{
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
	expected := map[string]map[string]string{
		"example.com": map[string]string{
			"ServerAliveInterval": "60",
			"Hostname":            "ssh.example.com",
			"Username":            "test",
		},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}
