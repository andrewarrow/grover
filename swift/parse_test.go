package swift

import "testing"

func TestParse1(t *testing.T) {
	Parse("example.swift")
}

/*
func TestParse1(t *testing.T) {
	item := Parse("example.swift")
	if len(item.Classes) != 2 {
		t.Fatal("wrong length")
	}
	name := item.Classes[0].Name
	if name != "AnotherClass" {
		t.Fatal("wrong name", name)
	}
	name = item.Classes[1].Name
	if name != "RootViewController" {
		t.Fatal("wrong name", name)
	}
}

func TestParse2(t *testing.T) {
	item := Parse("example.swift")
	if len(item.Classes[0].Functions) != 1 {
		t.Fatal("wrong length")
	}
	if len(item.Classes[1].Functions) != 4 {
		t.Fatal("wrong length")
	}
}
*/
