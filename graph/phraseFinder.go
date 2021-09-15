package graph

import (
	"regexp"
	"sort"
	"strings"

	"github.com/daffinito/PhraseFinder/graph/model"
)

// PhraseFinder is the entry point, it takes the raw text provided in the request and returns the completed response
func PhraseFinder(text string) []*model.Phrase {
	newText := sanitizeText(text)
	phraseMap := getPhrases(newText)
	phrases := buildResponse(phraseMap)
	phrases = sortResponse(phrases, 100)

	return phrases
}

// sanitizeText removes all the new lines, punctuation, extra whitespace, and converts the text to lower case
func sanitizeText(text string) string {
	replaceWithSpace := " "
	var replacer = strings.NewReplacer(
		"\r\n", replaceWithSpace,
		"\n\r", replaceWithSpace,
		"\r", replaceWithSpace,
		"\n", replaceWithSpace,
		"\t", replaceWithSpace,
	)
	newText := replacer.Replace(text)
	newText = strings.ReplaceAll(newText, "'", "")
	rePunc := regexp.MustCompile(`[^\w\s]`)
	newText = rePunc.ReplaceAllString(newText, " ")
	reDblSpaces := regexp.MustCompile(`\s+`)
	newText = reDblSpaces.ReplaceAllString(newText, " ")

	return strings.ToLower(strings.TrimSpace(newText))
}

// getPhrases loops through the text and creates a map[string]int with all the 3 word phrases,
// where the string is the phrase and the int is the number of times the phrase appears
func getPhrases(text string) map[string]int {
	phrases := make(map[string]int)
	var phrase strings.Builder
	words := strings.Split(text, " ")
	numWords := len(words)
	for n := range words {
		if numWords > n+2 {
			phrase.WriteString(words[n])
			phrase.WriteString(" ")
			phrase.WriteString(words[n+1])
			phrase.WriteString(" ")
			phrase.WriteString(words[n+2])
			phraseString := phrase.String()

			if _, ok := phrases[phraseString]; !ok {
				phrases[phraseString] = 1
			} else {
				phrases[phraseString]++
			}

			phrase.Reset()
		}
	}
	return phrases
}

// buildResponse marshals the map[string]int to our Phrase struct
func buildResponse(phrases map[string]int) []*model.Phrase {
	var response []*model.Phrase
	for key, val := range phrases {
		response = append(response, &model.Phrase{
			Text:  key,
			Count: val,
		})
	}

	return response
}

// sortResponse sorts the Phrase struct from most common phrase to least common phrase, and
// limits the slice size by the limit variable
func sortResponse(phrases []*model.Phrase, limit int) []*model.Phrase {
	sort.SliceStable(phrases, func(i, j int) bool {
		return phrases[i].Count > phrases[j].Count
	})
	if len(phrases) <= limit {
		return phrases
	}
	return phrases[0:limit]
}
