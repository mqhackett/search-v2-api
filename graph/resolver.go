package graph

import (
	"fmt"

	"github.com/open-cluster-management/search-v2-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	// searchResult []*model.SearchResult

	searchResultRelated []*model.SearchRelatedResult
}

// type SearchResultItems struct {
// 	Items []map[string]interface{}
// }

// type SearchResultCount struct {
// 	Count int
// }

type SearchResult struct { //car
	Count int
	Items []map[string]interface{}
	// Related [SearchRelatedResult]
}

type SearchResultRelatedItems struct { //person
	Items []map[string]interface{}
}

func (sr *SearchResult) Related() *SearchResultRelatedItems {
	related := &SearchResultRelatedItems{}

	fmt.Println(related)
	return related
}

// type SearchResultRelatedItems { //person
//   items: [Map]
//   }

// type SeachResultRelatedCount {
//   count: Int
// }

// type SearchRelatedResult {
//     kind: String!
//     count: Int
//     items: [Map]
//   }
