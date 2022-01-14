package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stolostron/search-v2-api/graph/generated"
	"github.com/stolostron/search-v2-api/graph/model"
	"github.com/stolostron/search-v2-api/pkg/schema"
	"github.com/vektah/gqlparser/v2/ast"
	klog "k8s.io/klog/v2"
)

func (r *mutationResolver) DeleteSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SaveSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func fieldSelections(ctx context.Context) map[string]bool {
	reqCtx := graphql.GetOperationContext(ctx)
	// var sels []string
	selMap := map[string]bool{}
	fieldSelections := graphql.GetFieldContext(ctx).Field.Selections

	for _, sel := range fieldSelections {
		switch sel := sel.(type) {
		case *ast.Field:
			// sels = append(sels, fmt.Sprintf("%s as %s", sel.Name, sel.Alias))
			// sels = append(sels, sel.Name)
			if sel.Name == schema.COUNT || sel.Name == schema.ITEMS || sel.Name == schema.RELATED {
				selMap[sel.Name] = true
			}
		case *ast.InlineFragment:
			// sels = append(sels, fmt.Sprintf("inline fragment on %s", sel.TypeCondition))
			selMap[sel.TypeCondition] = true

		case *ast.FragmentSpread:
			fragment := reqCtx.Doc.Fragments.ForName(sel.Name)
			// sels = append(sels, fmt.Sprintf("named fragment %s on %s", sel.Name, fragment.TypeCondition))
			selMap[sel.Name+""+fragment.TypeCondition] = true

		}
	}
	klog.Infoln("***** sels for search: ", selMap)
	// for i, sel := range sels {
	// 	klog.Infoln(i, sel)
	// }
	return selMap
}

func (r *queryResolver) Search(ctx context.Context, input []*model.SearchInput) ([]*model.SearchResult, error) {
	// var count int
	klog.Infof("--------- Received Search query with %d inputs ---------\n", len(input))
	selMap := fieldSelections(ctx)
	switch {
	case selMap[schema.COUNT]:
		klog.Infoln("Only count")

	case selMap[schema.RELATED]:
		klog.Infoln("Only related")

	case selMap[schema.ITEMS]:
		klog.Infoln("Only items")
	}
	return schema.Search(ctx, input, selMap)
	// return nil, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	klog.Infoln("Received Messages query")

	messages := make([]*model.Message, 0)
	kind := "Informational"
	desc := "Trial search-v2-api"
	message1 := model.Message{ID: "1", Kind: &kind, Description: &desc}
	messages = append(messages, &message1)
	return messages, nil
}

func (r *queryResolver) SearchSchema(ctx context.Context) (map[string]interface{}, error) {
	klog.Infoln("Received SearchSchema query")
	return schema.SearchSchema(ctx)
}

func (r *queryResolver) SavedSearches(ctx context.Context) ([]*model.UserSearch, error) {
	klog.Infoln("Received SavedSearches query")
	savedSrches := []*model.UserSearch{}
	// savedSrches := make([]*model.UserSearch, 0)
	// id := "1"
	// name := "savedSrch1"
	// srchText := "Trial savedSrch1"
	// desc := "Trial search-v2-api savedSrch1"
	// savedSrch1 := model.UserSearch{ID: &id, Name: &name, Description: &desc, SearchText: &srchText}
	// savedSrches = append(savedSrches, &savedSrch1)
	// return savedSrches, nil
	return savedSrches, nil
}

func (r *queryResolver) SearchComplete(ctx context.Context, property string, query *model.SearchInput, limit *int) ([]*string, error) {
	klog.Infof("Received SearchComplete query with input property **%s** and limit %d", property, limit)
	return schema.SearchComplete(ctx, property, query, limit)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
