package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

var (
	itemCache *lru.ARCCache
)

const (
	cacheSize = 10000
)

func init() {
	var err error

	// Initiate item LRU cache
	if itemCache, err = lru.NewARC(cacheSize); err != nil {
		panic(fmt.Sprintf("Panic: Unable to create itemCache: %v", err))
	}
}
