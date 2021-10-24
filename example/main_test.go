package main

import (
	"log"
	"testing"
	"time"

	"github.com/skandyla/go-cache-sample"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	cache := cache.NewCache()
	err := cache.Set("id", 1111, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	got, _ := getValue(cache, "id")
	want := 1111
	assert.Equal(t, want, got)
}
