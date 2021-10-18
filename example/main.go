package main

import (
	"fmt"
	"log"

	cache "github.com/skandyla/go-cache-sample"
)

func main() {
	cache := cache.NewCache()

	err := cache.Set("userId", 42)
	if err != nil {
		log.Fatalln(err)
	}
	err = cache.Set("otherId", 77)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", cache)

	//It is good to separate your "domain" code from the outside world (side-effects). The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
	fmt.Println(getValue(cache, "userId"))
	//fmt.Println(cache.Get("userId"))

	err = cache.Delete("userId")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(getValue(cache, "userId"))

	fmt.Printf("%+v\n", cache)
}

func getValue(cache *cache.Cache, key string) interface{} {
	return cache.Get(key)
}
