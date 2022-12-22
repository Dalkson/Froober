package main

import (
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
	// Initialize a new strings.Builder object.
	var newCopy strings.Builder
	// Get the text on the clipboard.
	text, err := glippy.Get()
	if err != nil {
		// Panic if there was an error getting the clipboard text.
		panic(err)
	}

	// If the text does not contain "http", return.
	if !strings.Contains(text, "http") {
		return
	}
	// If the text contains "http://froob.org/", return.
	if strings.Contains(text, "http://froob.org/") {
		return
	}
	// Split the text into words.
	words := strings.Fields(text)
	// Iterate over the words.
	for i, s := range words {
		// If it is the first word:
		if i == 0 {
			// Check if it starts with "https://" or "http://".
			// If it does, append "http://froob.org/" to newCopy.
			// If not, append the word as is.
			if strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://") {
				newCopy.WriteString("http://froob.org/")
			} else {
				newCopy.WriteString(s)
			}
		} else {
			// If it is not the first word:
			// Check if it starts with "https://" or "http://".
			// If it does, append " http://froob.org/" to newCopy.
			// If not, append " " + the word.
			if strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://") {
				newCopy.WriteString(" http://froob.org/")
			} else {
				newCopy.WriteString(" " + s)
			}
		}
	}

	// Set the text on the clipboard to the modified text in newCopy.
	glippy.Set(newCopy.String())
}
