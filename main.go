package main

import (
	"dataMiningGolang/crawl/googlecrawl"
	"dataMiningGolang/crawl/wikipediacrawl"
	"fmt"
)

func main() {

	fmt.Println(googlecrawl.Handle())
	fmt.Println(wikipediacrawl.Handle())

}
