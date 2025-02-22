package cos418_hw1_1

import (
	"fmt"
	"sort"
	"strings"
	"io/ioutil"
    "log"
    "regexp"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.

func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"

	//reg := regexp.MustCompile(`[a-zA-Z0-9]+`)
    
	//reg := regexp.MustCompile(`[^0-9a-zA-Z]+`)
	reg := regexp.MustCompile(`'+`)
	bs, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
	text := strings.ToLower(string(bs))
   
	ms:=reg.ReplaceAllString(text,"")
	//fmt.Printf("%s",ms)
	
	reg2 := regexp.MustCompile(`[a-zA-Z0-9]+`)
	
	matches := reg2.FindAllString(ms, -1)
    groups := make(map[string]int)
    for _, match := range matches {
		if strings.Count(match,"")>charThreshold {
	      groups[match]++
		 }
    }
    var WordCounts []WordCount
    for k, v := range groups {
        WordCounts = append(WordCounts, WordCount{k, v})
    }
    
	sortWordCounts(WordCounts)
	
    for i := 1; i <= numWords; i++ {
		word := WordCounts[i-1].Word
        cnt := WordCounts[i-1].Count
	    fmt.Printf("%v: %v", word, cnt)
    }
	
	
	return nil
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
