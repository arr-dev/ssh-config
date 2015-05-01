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
			"Hostname":                 "github.com",
			"PreferredAuthentications": "publickey",
		},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected ", expected, "Got ", got)
	}
}
