package main

import (
	"fmt"
	"lru/internal/lrucache"
)

func main() {
	lru := lrucache.NewLRUCache(2)
	fmt.Println("Add key1")
	ok := lru.Add("key1", "val1")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Get key1")
	val, ok := lru.Get("key1")
	fmt.Println(ok)
	fmt.Println(val)
	fmt.Println("============")

	fmt.Println("Add key1 with another val")
	ok = lru.Add("key1", "val123")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Add key2")
	ok = lru.Add("key2", "val2")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Get key3")
	val, ok = lru.Get("key3")
	fmt.Println(ok)
	fmt.Println(val)
	fmt.Println("============")

	fmt.Println("Get key1")
	val, ok = lru.Get("key1")
	fmt.Println(ok)
	fmt.Println(val)
	fmt.Println("============")

	fmt.Println("Add key5")
	ok = lru.Add("key5", "val5")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Add key6")
	ok = lru.Add("key6", "val6")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Remove key1")
	ok = lru.Remove("key1")
	fmt.Println(ok)
	fmt.Println("============")

	fmt.Println("Remove key5")
	ok = lru.Remove("key5")
	fmt.Println(ok)
	fmt.Println("============")

}
