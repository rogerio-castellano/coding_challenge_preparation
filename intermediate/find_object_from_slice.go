package main

import (
	"fmt"
	"slices"
)

type Config struct {
	Key   string
	Value string
}

func main() {
	var myconfig = []Config{
		{Key: "foo", Value: "bar"},
		{Key: "key1", Value: "test"},
		{Key: "web/key1", Value: "test2"},
	}

	idx := slices.IndexFunc(myconfig, func(c Config) bool { return c.Key == "key1" })

	fmt.Println(idx)
}