package main

import (
	"fmt"

	cache "github.com/skandyla/go-cache-sample"
)

func main() {
	cache := cache.NewCache()

	cache.Set("userId", 42)
	cache.Set("otherId", 77)
	fmt.Printf("%+v\n", cache)

	fmt.Println(cache.Get("userId"))
	fmt.Println(cache.Get("otherId"))

	cache.Delete("userId")
	fmt.Println(cache.Get("userId"))

	fmt.Printf("%+v\n", cache)
}
