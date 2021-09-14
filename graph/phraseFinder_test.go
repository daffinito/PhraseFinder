package graph

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/daffinito/PhraseFinder/graph/model"
)

func Test_sanitizeText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Replaces new lines with a single space",
			args: args{
				text: "Test\nTest\r\nTest\rTest\n\rTest",
			},
			want: "test test test test test",
		}, {
			name: `Replaces punc found between words with a space except '`,
			args: args{
				text: "This?is!a test 2.5, test",
			},
			want: "this is a test 2 5 test",
		}, {
			name: `Whitespace is trimmed`,
			args: args{
				text: "     test      ",
			},
			want: "test",
		}, {
			name: `Multiple spaces are fixed`,
			args: args{
				text: "test  test   test    test",
			},
			want: "test test test test",
		}, {
			name: `Contractions are replaced with no space`,
			args: args{
				text: "This isn't a test.it's a tribute",
			},
			want: "this isnt a test its a tribute",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeText(tt.args.text); got != tt.want {
				t.Errorf("sanitizeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPhrases(t *testing.T) {
	type args struct {
		text  string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: `Finds all phrases`,
			args: args{
				text:  "this is a test this is a test this is a test",
			},
			want: map[string]int{
				"this is a":    3,
				"is a test":    3,
				"a test this":  2,
				"test this is": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPhrases(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPhrases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortResponse(t *testing.T) {
	type args struct {
		phrases []*model.Phrase
		limit int
	}
	tests := []struct {
		name string
		args args
		want []*model.Phrase
	}{
		{
			name: `Response is sorted`,
			args: args{
				phrases: []*model.Phrase{
					{
						Text:  `test this is`,
						Count: 1,
					}, {
						Text:  `this is a`,
						Count: 4,
					}, {
						Text:  `a test this`,
						Count: 2,
					}, {
						Text:  `is a test`,
						Count: 3,
					},
				},
				limit: 100,
			},
			want: []*model.Phrase{
				{
					Text:  `this is a`,
					Count: 4,
				}, {
					Text:  `is a test`,
					Count: 3,
				}, {
					Text:  `a test this`,
					Count: 2,
				}, {
					Text:  `test this is`,
					Count: 1,
				},
			},
		}, {
			name: `Response is limited`,
			args: args{
				phrases: []*model.Phrase{
					{
						Text:  `test this is`,
						Count: 1,
					}, {
						Text:  `this is a`,
						Count: 4,
					}, {
						Text:  `a test this`,
						Count: 2,
					}, {
						Text:  `is a test`,
						Count: 3,
					},
				},
				limit: 2,
			},
			want: []*model.Phrase{
				{
					Text:  `this is a`,
					Count: 4,
				}, {
					Text:  `is a test`,
					Count: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortResponse(tt.args.phrases, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildResponse(t *testing.T) {
	type args struct {
		phrases map[string]int
	}
	tests := []struct {
		name string
		args args
		want []*model.Phrase
	}{
		{
			name: `Response is built`,
			args: args{
				phrases: map[string]int{
					"this is a": 2,
				},
			},
			want: []*model.Phrase{
				{
					Text:  `this is a`,
					Count: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildResponse(tt.args.phrases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhraseFinder(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []*model.Phrase
	}{
		{
			name: `Phrases are parsed and sorted`,
			args: args{
				text: "    Test i's test.is te'st   !    ",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhraseFinder(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				gotJson, _ := json.Marshal(got)
				wantJson, _ := json.Marshal(tt.want)
				t.Errorf("PhraseFinder() = %s, want %s", gotJson, wantJson)
			}
		})
	}
}
