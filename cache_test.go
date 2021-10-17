package cache

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var c *Cache

func init() {
	c = NewCache()
	c.items["id"] = 1
	c.items["nnn"] = "blabla"
	c.items["bool"] = false

}

func TestGetSimple(t *testing.T) {
	got := c.Get("id")
	want := 1
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetSimpleReflect(t *testing.T) {
	got := c.Get("id")
	want := 1
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetSimpleTable(t *testing.T) {
	var cases = []struct {
		input string
		want  interface{}
	}{
		{"id", 1},
		{"nnn", "blabla"},
		{"bool", false},
	}
	for _, test := range cases {
		got := c.Get(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestGetSimpleTableSubTests(t *testing.T) {
	var cases = []struct {
		input string
		want  interface{}
	}{
		{"id", 1},
		{"nnn", "blabla"},
		{"bool", false},
	}
	for _, test := range cases {
		testname := fmt.Sprintf("%s,%v", test.input, test.want)
		t.Run(testname, func(t *testing.T) {
			got := c.Get(test.input)
			want := test.want
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestGetSimpleTableSubTestsNames(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  interface{}
	}{
		"getInt":    {input: "id", want: 1},
		"getString": {input: "nnn", want: "blabla"},
		"getBool":   {input: "bool", want: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := c.Get(tc.input)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Fatalf("expected: %v, got: %v", tc.want, got)
			//}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestGetSimpleTableAssert(t *testing.T) {
	assert := assert.New(t)
	var tests = map[string]struct {
		input string
		want  interface{}
	}{
		"getInt":    {input: "id", want: 1},
		"getString": {input: "nnn", want: "blabla"},
		"getBool":   {input: "bool", want: false},
	}
	for name, test := range tests {
		//assert.Equal(test.want, c.Get(test.input))
		//convert to subtests
		t.Run(name, func(t *testing.T) {
			assert.Equal(test.want, c.Get(test.input))
		})
	}
}
