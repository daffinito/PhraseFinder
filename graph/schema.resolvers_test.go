package graph

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/daffinito/PhraseFinder/graph/generated"
	"github.com/daffinito/PhraseFinder/graph/model"
)

func Test_queryResolver_FindPhrasesFromText(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
		in  string
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Phrase
		wantErr bool
	}{
		{
			name: `Phrases are parsed and sorted`,
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: ctx,
				in:  "Test is test.is te'st   !",
			},
			want: []*model.Phrase{
				{
					Text:  `test is test`,
					Count: 2,
				}, {
					Text:  `is test is`,
					Count: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.FindPhrasesFromText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.FindPhrasesFromText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				gotJson, _ := json.Marshal(got)
				wantJson, _ := json.Marshal(tt.want)
				t.Errorf("queryResolver.FindPhrasesFromText() = %s,\n want %s", gotJson, wantJson)
			}
		})
	}
}

func TestResolver_Query(t *testing.T) {
	tests := []struct {
		name string
		r    *Resolver
		want generated.QueryResolver
	}{
		{
			name: `Query returns valid response`,
			r:    &Resolver{},
			want: &queryResolver{&Resolver{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resolver{}
			if got := r.Query(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolver.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolver_Mutation(t *testing.T) {
	tests := []struct {
		name string
		r    *Resolver
		want generated.MutationResolver
	}{
		{
			name: `Mutation returns valid response`,
			r:    &Resolver{},
			want: &mutationResolver{&Resolver{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resolver{}
			if got := r.Mutation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolver.Mutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_FindPhrasesFromFile(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx  context.Context
		file graphql.Upload
	}

	ctx := context.Background()
	f, err := os.Open("../fixtures/testfile")
	if err != nil {
		t.Errorf("mutationResolver.FindPhrasesFromFile() - unable to open testfile")
	}
	defer f.Close()
	file := graphql.Upload{
		File:        f,
		Filename:    "testfile",
		Size:        int64(5),
		ContentType: "text/plain",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Phrase
		wantErr bool
	}{
		{
			name: `Parses from file are parsed and sorted`,
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx:  ctx,
				file: file,
			},
			want: []*model.Phrase{
				{
					Text:  `test is test`,
					Count: 2,
				}, {
					Text:  `is test is`,
					Count: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.FindPhrasesFromFile(tt.args.ctx, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.FindPhrasesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.FindPhrasesFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
