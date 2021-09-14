package graph

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/daffinito/3wpapi/graph/generated"
	"github.com/daffinito/3wpapi/graph/model"
)

func Test_queryResolver_FindPhrases(t *testing.T) {
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
			got, err := r.FindPhrases(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.FindPhrases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				gotJson, _ := json.Marshal(got)
				wantJson, _ := json.Marshal(tt.want)
				t.Errorf("queryResolver.FindPhrases() = %s,\n want %s", gotJson, wantJson)
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
			r: &Resolver{},
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
