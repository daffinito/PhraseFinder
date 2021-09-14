package graph

import (
	"regexp"
	"sort"
	"strings"

	"github.com/daffinito/3wpapi/graph/model"
)

func PhraseFinder(text string) []*model.Phrase {
	newText := sanitizeText(text)
	phraseMap := getPhrases(newText, 100)
	phrases := buildResponse(phraseMap)
	phrases = sortResponse(phrases)

	return phrases
}

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

func getPhrases(text string, limit int) map[string]int {
	phrases := make(map[string]int)
	var phrase strings.Builder
	words := strings.Split(text, " ")
	numWords := len(words)
	for n := range words {
		if n >= limit {
			break;
		}
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

func sortResponse(phrases []*model.Phrase) []*model.Phrase {
	sort.SliceStable(phrases, func(i, j int) bool {
		return phrases[i].Count > phrases[j].Count
	})
	return phrases
}