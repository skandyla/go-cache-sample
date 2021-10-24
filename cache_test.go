package cache

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

//var c *Cache

//func init() {
//	c = NewCache()
//	c.items["id"] = Item{1, time.Second * 2}
//	c.items["nnn"] = Item{"blabla", time.Second * 3}
//	c.items["bool"] = Item{false, time.Second * 4}
//}

//func TestGetSimple(t *testing.T) {
//	got := c.Get("id")
//	want := 1
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}
//
//func TestGetSimpleReflect(t *testing.T) {
//	got := c.Get("id")
//	want := 1
//	if !reflect.DeepEqual(want, got) {
//		t.Fatalf("expected: %v, got: %v", want, got)
//	}
//}

//func TestGetSimpleTable(t *testing.T) {
//	var cases = []struct {
//		input string
//		want  interface{}
//	}{
//		{"id", 1},
//		{"nnn", "blabla"},
//		{"bool", false},
//	}
//	for _, test := range cases {
//		got := c.Get(test.input)
//		want := test.want
//		if got != want {
//			t.Errorf("got %q want %q", got, want)
//		}
//	}
//}

//func TestGetSimpleTableSubTests(t *testing.T) {
//	var cases = []struct {
//		input string
//		want  interface{}
//	}{
//		{"id", 1},
//		{"nnn", "blabla"},
//		{"bool", false},
//	}
//	for _, test := range cases {
//		testname := fmt.Sprintf("%s,%v", test.input, test.want)
//		t.Run(testname, func(t *testing.T) {
//			got := c.Get(test.input)
//			want := test.want
//			if got != want {
//				t.Errorf("got %q want %q", got, want)
//			}
//		})
//	}
//}

func TestGetSimpleTableSubTestsNames(t *testing.T) {
	c := NewCache()
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
	c := NewCache()
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

/// tests for Set
//type setInput struct {
//	key   Item
//	value interface{}
//}
//
//func TestSet(t *testing.T) {
//	var tests = map[string]struct {
//		input *Item
//		want  interface{} //want nil
//	}{
//		"setInt":    {input: &Item{name: "id", ttl: time.Second * 3}, want: nil},
//		"setString": {input: &Item{name: "idstr", ttl: time.Second * 4}, want: nil},
//	}
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			got := c.Set(tc.input.name, tc.input.ttl)
//			diff := cmp.Diff(tc.want, got)
//			if diff != "" {
//				t.Fatalf(diff)
//			}
//		})
//	}
//}
//
//// Delete
//func TestDelete(t *testing.T) {
//	c = NewCache()
//	c.items["id"] = 111
//
//	var tests = map[string]struct {
//		input *setInput
//		want  interface{} // want nil
//	}{
//		"setInt": {input: &setInput{key: "id", value: 111}, want: nil},
//	}
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			got := c.Delete(tc.input.key)
//			diff := cmp.Diff(tc.want, got)
//			if diff != "" {
//				t.Fatalf(diff)
//			}
//		})
//	}
//}
//
