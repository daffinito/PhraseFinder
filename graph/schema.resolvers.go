package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io/ioutil"

	"github.com/99designs/gqlgen/graphql"
	"github.com/daffinito/PhraseFinder/graph/generated"
	"github.com/daffinito/PhraseFinder/graph/model"
)

func (r *mutationResolver) FindPhrasesFromFile(ctx context.Context, file graphql.Upload) ([]*model.Phrase, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	return PhraseFinder(string(content)), nil
}

func (r *queryResolver) FindPhrasesFromText(ctx context.Context, text string) ([]*model.Phrase, error) {
	return PhraseFinder(text), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

