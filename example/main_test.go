package main

import (
	"testing"

	"github.com/skandyla/go-cache-sample"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	cache := cache.NewCache()
	cache.Set("id", 1111)
	got := getValue(cache, "id")
	want := 1111
	assert.Equal(t, want, got)
}
