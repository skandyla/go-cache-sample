package cache

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

//TestGet - ensure Get method is working
func TestGet(t *testing.T) {
	c := New()
	c.items["id"] = Item{1, time.Second * 2}
	c.items["nnn"] = Item{"blabla", time.Second * 3}
	c.items["bool"] = Item{false, time.Second * 4}

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
			got, _ := c.Get(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

//TestGetAssert - ensure Get method is working
func TestGetAssert(t *testing.T) {
	c := New()
	c.items["id"] = Item{1, time.Second * 2}
	c.items["nnn"] = Item{"blabla", time.Second * 3}
	c.items["bool"] = Item{false, time.Second * 4}

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
		got, _ := c.Get(test.input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(test.want, got)
		})
	}
}

//TestSet - ensure Set method is working
func TestSet(t *testing.T) {
	c := New()
	assert := assert.New(t)
	var tests = map[string]struct {
		input string
		want  interface{}
		ttl   time.Duration
	}{
		"getInt":    {input: "id", want: 1, ttl: time.Second * 2},
		"getString": {input: "nnn", want: "blabla", ttl: time.Second * 3},
	}

	//fill cache with our testing function.
	for _, tc := range tests {
		//c.items[tc.input] = Item{tc.want, tc.ttl}
		c.Set(tc.input, tc.want, tc.ttl) //the same as above but via method we test
	}

	//verify
	for name, tc := range tests {
		got, _ := c.Get(tc.input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(tc.want, got)
		})
	}
}

//TestDelete - ensure Delete method is working
func TestDelete(t *testing.T) {
	c := New()
	assert := assert.New(t)
	var tests = map[string]struct {
		input     string        //Item.name
		value     interface{}   //Item.value
		ttl       time.Duration //Item.ttl
		want      interface{}   //what we want in test
		isDeleted bool          //is item deleted ?
	}{
		"Int":    {input: "id", value: 1, ttl: time.Second * 2, want: nil, isDeleted: true},
		"String": {input: "nnn", value: "blabla", ttl: time.Second * 3, want: "blabla", isDeleted: false},
	}

	//fill
	for _, tc := range tests {
		c.Set(tc.input, tc.value, tc.ttl)
	}

	//delete
	for _, tc := range tests {
		if tc.isDeleted {
			c.Delete(tc.input)
		}
	}

	//verify
	//TBD: handle error responses
	for name, tc := range tests {
		got, _ := c.Get(tc.input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(tc.want, got)
		})
	}
}
