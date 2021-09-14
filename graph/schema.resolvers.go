package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/daffinito/3wpapi/graph/generated"
	"github.com/daffinito/3wpapi/graph/model"
)

func (r *queryResolver) FindPhrases(ctx context.Context, in string) ([]*model.Phrase, error) {
	return PhraseFinder(in), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
