package camelcase

import "testing"

func TestCamelCase(t *testing.T) {
	c := "HelloWorld"
	r := CamelCase(c)

	if r != "helloWorld" {
		t.Errorf("not convert in camelcase")
	}
}

func TestCamelCaseWithUnderline(t *testing.T) {
	c := "Hello_world"
	r := CamelCase(c)

	if r != "helloWorld" {
		t.Errorf("not convert in camelcase")
	}
}

func TestSnakeCase(t *testing.T) {
	c := "HelloWorld"
	r := SnakeCase(c)

	if r != "hello_world" {
		t.Errorf("not convert in snake case")
	}
}

func TestPascalCase(t *testing.T) {
	c := "helloWorld"
	r := PascalCase(c)

	if r != "HelloWorld" {
		t.Errorf("not convert in pascal case")
	}
}
