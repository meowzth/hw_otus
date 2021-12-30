package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(st string) []string {
	st = strings.ToLower(st)
	escapeChars := []string{"!", ".", ",", "?", ":", ";", "'", `"`}
	for _, char := range escapeChars {
		st = strings.ReplaceAll(st, char, "")
	}
	stSlice := strings.Fields(st)

	// building frequency map
	freqMap := make(map[string]int)
	for _, word := range stSlice {
		_, ok := freqMap[word]
		if !ok {
			freqMap[word] = 1
		}
		if ok {
			freqMap[word] = freqMap[word] + 1
		}
	}

	// select top 10 keys from map
	type kv struct {
		Key   string
		Value int
	}
	var freqSlice []kv
	for k, v := range freqMap {
		freqSlice = append(freqSlice, kv{k, v})
	}
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Value > freqSlice[j].Value
	})
	result := []string{}
	for _, kv := range freqSlice[:10] {
		result = append(result, kv.Key)
	}

	return result
}
