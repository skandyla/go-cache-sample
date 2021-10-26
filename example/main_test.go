package main

import (
	"testing"
	"time"

	"github.com/skandyla/go-cache-sample"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	cache := cache.New()
	cache.Set("id", 1111, time.Second)
	got, _ := getValue(cache, "id")
	want := 1111
	assert.Equal(t, want, got)
}
