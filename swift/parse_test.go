package swift

import (
	"testing"
)

func TestParse(t *testing.T) {
	item := Parse("example.swift")
	if len(item.Classes) != 2 {
		t.Fatal("wrong length")
	}
	name := item.Classes[0].Name
	if name != "AnotherClass" {
		t.Fatal("wrong name", name)
	}
	name = item.Classes[0].Name
	if name != "RootViewController" {
		t.Fatal("wrong name", name)
	}

}
