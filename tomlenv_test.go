package tomlenv

import (
	"testing"
)

type Toml struct {
	Top string
	Foo Foo
	Bar []Bar
	Env map[string]Env
}

type Foo struct {
	V1 string
	V2 int
}

type Bar struct {
	V string
}

type Env struct {
	V string
}

func TestDecodeEnvCommon(t *testing.T) {
	nm := "TestDecodeEnvCommon"
	cs := "common => dev"
	// common => dev
	tml := Toml{}
	if _, err := DecodeEnv("dev", "./test", &tml); err != nil {
		t.Fatalf("%s %s failed. %v", nm, cs, err)
	}
	expected := Toml{
		Top: "dev_value",
		Foo: Foo{
			V1: "common_value",
			V2: 1,
		},
		Bar: []Bar{
			{V: "common_value1"},
			{V: "common_value2"},
		},
		Env: map[string]Env{
			"dev": {V: "dev_value"},
		},
	}

	if tml.Top != expected.Top {
		t.Fatalf("%s %s Top failed. expected: %s, actual: %s", nm, cs, expected.Top, tml.Top)
	}
	if tml.Foo != expected.Foo {
		t.Fatalf("%s %s Foo failed. expected: %v, actual: %v", nm, cs, expected.Foo, tml.Foo)
	}
	if len(tml.Bar) != len(expected.Bar) ||
		len(tml.Bar) != 2 ||
		tml.Bar[0] != expected.Bar[0] ||
		tml.Bar[1] != expected.Bar[1] {

		t.Fatalf("%s %s Bar failed. expected: %v, actual: %v", nm, cs, expected.Bar, tml.Bar)
	}
	if len(tml.Env) != len(expected.Env) ||
		len(tml.Env) != 1 ||
		tml.Env["dev"] != expected.Env["dev"] {

		t.Fatalf("%s %s Env failed. expected: %v, actual: %v", nm, cs, expected.Env, tml.Env)
	}
}

func TestDecodeEnvFooBar(t *testing.T) {
	nm := "TestDecodeEnvFooBar"
	cs := "foo => bar => dev"
	// common => dev
	tml := Toml{}
	if _, err := DecodeEnv("dev", "./test", &tml, "foo", "bar"); err != nil {
		t.Fatalf("%s %s failed. %v", nm, cs, err)
	}
	expected := Toml{
		Top: "dev_value",
		Foo: Foo{
			V1: "foo_value",
			V2: 1,
		},
		Bar: []Bar{
			{V: "bar_value1"},
			{V: "bar_value2"},
		},
		Env: map[string]Env{
			"dev": {V: "dev_value"},
		},
	}

	if tml.Top != expected.Top {
		t.Fatalf("%s %s Top failed. expected: %s, actual: %s", nm, cs, expected.Top, tml.Top)
	}
	if tml.Foo != expected.Foo {
		t.Fatalf("%s %s Foo failed. expected: %v, actual: %v", nm, cs, expected.Foo, tml.Foo)
	}
	if len(tml.Bar) != len(expected.Bar) ||
		len(tml.Bar) != 2 ||
		tml.Bar[0] != expected.Bar[0] ||
		tml.Bar[1] != expected.Bar[1] {

		t.Fatalf("%s %s Bar failed. expected: %v, actual: %v", nm, cs, expected.Bar, tml.Bar)
	}
	if len(tml.Env) != len(expected.Env) ||
		len(tml.Env) != 1 ||
		tml.Env["dev"] != expected.Env["dev"] {

		t.Fatalf("%s %s Env failed. expected: %v, actual: %v", nm, cs, expected.Env, tml.Env)
	}
}
