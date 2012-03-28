package main

import (
	"testing"
)

var expectedTags = []Tag{
	Tag{Name: "TestPackage", File: "tests/input.go", Address: "1", Type: "p", Fields: map[string]string{"line": "1"}},
	Tag{Name: "go/ast", File: "tests/input.go", Address: "4", Type: "i", Fields: map[string]string{"line":"4"}},
	Tag{Name: "variable", File: "tests/input.go", Address: "7", Type: "v", Fields: map[string]string{"line": "7"}},
	Tag{Name: "constant", File: "tests/input.go", Address: "9", Type: "c", Fields: map[string]string{"line": "9"}},
	Tag{Name: "Function1", File: "tests/input.go", Address: "11", Type: "f", Fields: map[string]string{"line": "11", "access": "public", "signature": "()"}},
	Tag{Name: "function2", File: "tests/input.go", Address: "14", Type: "f", Fields: map[string]string{"line": "14", "access": "private", "signature": "(p1, p2 int, p3 *string)"}},
	Tag{Name: "Struct", File: "tests/input.go", Address: "17", Type: "t", Fields: map[string]string{"line": "17"}},
	Tag{Name: "MyInt", File: "tests/input.go", Address: "23", Type: "t", Fields: map[string]string{"line": "23"}},
	Tag{Name: "F1", File: "tests/input.go", Address: "25", Type: "f", Fields: map[string]string{"line": "25", "access": "public", "signature": "()", "type": "MyInt"}},
}

func TestParse(t *testing.T) {
	tags, err := Parse("tests/input.go")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(tags) != len(expectedTags) {
		t.Fatalf("len(tags) == %d, want %d", len(tags), len(expectedTags))
	}

	for i, exp := range expectedTags {
		if tags[i].String() != exp.String() {
			t.Errorf("tag(%d)\n  is:%s\nwant:%s", i, tags[i].String(), exp.String())
		}
	}
}