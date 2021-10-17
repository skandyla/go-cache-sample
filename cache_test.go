package cache

import (
	"testing"

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

func TestGetSimpleTable(t *testing.T) {
	var cases = []struct {
		input    string
		expected interface{}
	}{
		{"id", 1},
		{"nnn", "blabla"},
		{"bool", false},
	}
	for _, test := range cases {
		got := c.Get(test.input)
		want := test.expected
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestGetSimpleTableAssert(t *testing.T) {
	assert := assert.New(t)
	var cases = []struct {
		input    string
		expected interface{}
	}{
		{"id", 1},
		{"nnn", "blabla"},
		{"bool", false},
	}
	for _, test := range cases {
		assert.Equal(c.Get(test.input), test.expected)
	}
}
