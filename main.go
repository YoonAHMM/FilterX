package main

import (
	"fmt"

	"github.com/FilterX/logic"
)

func init_SearchTest() *logic.Search {
	search := logic.NewSearch()
	list := []string{"15768"}
	search.SetKeywords(list)
	return search
}

func main() {
	serrch:=init_SearchTest()

	s:=serrch.GetStringFindFirst("15768hmm love ljx")
	fmt.Println(s)
}