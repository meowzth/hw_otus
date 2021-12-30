package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Top10(st string) []string {
	st = strings.ToLower(st)
	escapeChars := []string{"!", ".", ",", "?", ":", ";", "'", `"`}
	for _, char := range escapeChars {
		st = strings.ReplaceAll(st, char, "")
	}
	st = strings.ReplaceAll(st, "- ", "")
	stSlice := strings.Fields(st)

	// building frequency map
	freqMap := make(map[string]int)
	for _, word := range stSlice {
		_, ok := freqMap[word]
		if !ok {
			freqMap[word] = 1
		}
		if ok {
			freqMap[word]++
		}
	}

	// select top keys from map
	type kv struct {
		Key   string
		Value int
	}
	var freqSlice []kv
	for k, v := range freqMap {
		freqSlice = append(freqSlice, kv{k, v})
	}
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Key < freqSlice[j].Key
	})
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Value > freqSlice[j].Value
	})
	result := []string{}
	rng := min(len(freqSlice), 10)
	for _, kv := range freqSlice[:rng] {
		result = append(result, kv.Key)
	}

	return result
}
