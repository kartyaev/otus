package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var SIZE = 10

type Word struct {
	Word   string
	Weight int
}

var reg = regexp.MustCompile(`[^a-zA-Z0-9А-Яа-я\-]+`)

func Top10(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	wordsMap := buildWordsMap(text)
	wordsSlice := convertMapToSlice(wordsMap)
	sortWords(wordsSlice)
	result := getResultFromWordsSlice(wordsSlice)
	return result
}

func getResultFromWordsSlice(wordsSlice []Word) []string {
	result := make([]string, 0, SIZE)
	tempValue := Word{}
	for _, value := range wordsSlice {
		if len(result) == SIZE {
			break
		}
		if value.Weight <= tempValue.Weight || tempValue.Weight == 0 {
			result = append(result, value.Word)
			tempValue = value
		}
	}
	return result
}

func convertMapToSlice(wordsMap map[string]int) []Word {
	words := make([]Word, 0)
	for key, elem := range wordsMap {
		word := Word{
			key,
			elem,
		}
		words = append(words, word)
	}
	return words
}

func buildWordsMap(text string) map[string]int {
	splitText := strings.Fields(text)
	mapValues := make(map[string]int)
	for _, word := range splitText {
		value := strings.ToLower(reg.ReplaceAllString(word, ""))
		if value == "-" {
			continue
		}
		if len(value) == 0 {
			continue
		}
		if mapValues[value] == 0 {
			mapValues[value] = 1
		} else {
			mapValues[value]++
		}
	}
	return mapValues
}

func sortWords(words []Word) {
	sort.SliceStable(words, func(i, j int) bool {
		if words[i].Weight > words[j].Weight {
			return true
		} else if words[i].Weight == words[j].Weight {
			scom := strings.Compare(words[i].Word, words[j].Word)
			return scom < 0
		}
		return false
	})
}
