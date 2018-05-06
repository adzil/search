package search

import (
	"bytes"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

var stringSpace = " "

var bytesSpace = []byte(stringSpace)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// String searches for string pattern.
func String(data StringInterface, pattern string, callback func(index int)) {
	// Create new pattern matcher from pattern string
	rawpats := strings.Split(pattern, stringSpace)
	pats := make([]*search.Pattern, len(rawpats))
	for i, rawpat := range rawpats {
		pats[i] = search.New(language.English, search.IgnoreCase).CompileString(rawpat)
	}
	// Start find from any matches
	datalen := data.Len()
	indexes := make([]int, 0, min(64, datalen))
	for i := 0; i < datalen; i++ {
		// Calculate match number
		var nmatch int
		for _, pat := range pats {
			if s, e := pat.IndexString(data.At(i)); s >= 0 && e >= 0 {
				nmatch++
			}
		}
		// Check for matches
		if nmatch > 0 && nmatch < len(pats) && indexes != nil {
			indexes = append(indexes, i)
		} else if nmatch == len(pats) {
			if indexes != nil {
				indexes = nil
			}
			// Directly call callbacks
			callback(i)
		}
	}
	// Run alike match
	for _, idx := range indexes {
		callback(idx)
	}
}

// Bytes searches for byte slice pattern.
func Bytes(data BytesInterface, pattern []byte, callback func(index int)) {
	// Create new pattern matcher from pattern string
	rawpats := bytes.Split(pattern, bytesSpace)
	pats := make([]*search.Pattern, len(rawpats))
	for i, rawpat := range rawpats {
		pats[i] = search.New(language.English, search.IgnoreCase).Compile(rawpat)
	}
	// Start find from any matches
	datalen := data.Len()
	indexes := make([]int, 0, min(64, datalen))
	for i := 0; i < datalen; i++ {
		// Calculate match number
		var nmatch int
		for _, pat := range pats {
			if s, e := pat.Index(data.At(i)); s >= 0 && e >= 0 {
				nmatch++
			}
		}
		// Check for matches
		if nmatch > 0 && nmatch < len(pats) && indexes != nil {
			indexes = append(indexes, i)
		} else if nmatch == len(pats) {
			if indexes != nil {
				indexes = nil
			}
			// Directly call callbacks
			callback(i)
		}
	}
	// Run alike match
	for _, idx := range indexes {
		callback(idx)
	}
}
