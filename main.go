package main

import (
	"fmt"

	"github.com/FilterX/logic"
)

func init_SearchTest() *logic.Search {
	search := logic.NewSearch()
	list := []string{"f\u00a0a\u00a0l\u00a0u\u00a0n"}
	search.SetKeywords(list)
	return search
}

func main() {
	serrch:=init_SearchTest()

	serrch.GetStringFindFirst("f\u00a0a\u00a0l\u00a0u\u00a0n")
	fmt.Println(len("f\u00a0a\u00a0l\u00a0u\u00a0n"))
}