package main

import (
	"fmt"
	"time"

	cache "github.com/skandyla/go-cache-sample"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	fmt.Println(getValue(cache, "userId")) //42 <nil>

	cache.Set("otherId", 77, time.Second*7)
	fmt.Println(getValue(cache, "otherId")) //77 <nil>

	//verify results after 6 seconds
	time.Sleep(time.Second * 6)
	fmt.Println(getValue(cache, "userId"))  //<nil> item not exist
	fmt.Println(getValue(cache, "otherId")) //77 <nil>

	cache.Delete("otherId")
	cache.Delete("otherId") //test again

	fmt.Println(getValue(cache, "otherId")) //<nil> item not exist

	//cache closing case
	cache.Set("userId", 33, time.Second*2)
	fmt.Println(getValue(cache, "userId")) //33 <nil>
	cache.Close()                          // immidiately cause all cache expired
	fmt.Println(getValue(cache, "userId")) //<nil> item not exist
}

//It is good to separate your "domain" code from the outside world (side-effects). The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
func getValue(cache *cache.Cache, key string) (interface{}, error) {
	v, err := cache.Get(key)
	if err != nil {
		//fmt.Printf("Got err:%v\n", err) //side effect
		return nil, err
	}
	return v, nil
}
