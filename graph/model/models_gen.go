// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID          string  `json:"id"`
	Kind        *string `json:"kind"`
	Description *string `json:"description"`
}

type SeachRelatedCount struct {
	Count *int `json:"count"`
}

type SearchFilter struct {
	Property string    `json:"property"`
	Values   []*string `json:"values"`
}

type SearchInput struct {
	Keywords     []*string       `json:"keywords"`
	Filters      []*SearchFilter `json:"filters"`
	Limit        *int            `json:"limit"`
	RelatedKinds []*string       `json:"relatedKinds"`
}

type SearchRelatedResult struct {
	Kind  string                   `json:"kind"`
	Count *int                     `json:"count"`
	Items []map[string]interface{} `json:"items"`
}

type SearchResultCount struct {
	Count *int `json:"count"`
}

type SearchResultItems struct {
	Items []map[string]interface{} `json:"items"`
}

type UserSearch struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	SearchText  *string `json:"searchText"`
}
