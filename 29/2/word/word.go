package word

import (
	"strings"
)

// UseCount places the fields (words) from provided string into a map as a key with a value representing the number of words found
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of instances of groups of characters separated by spaces (words)
func Count(s string) int {
	xs := strings.Split(s, " ")
	return len(deleteEmptyStrings(xs))
}

// deleteEmptyStrings returns a copy of provided string slice without empty strings
func deleteEmptyStrings(xs []string) []string {
	var ns []string
	for _, str := range xs {
		if str != "" {
			ns = append(ns, str)
		}
	}
	return ns
}
