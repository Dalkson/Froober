package main

import (
	"strings"
	"time"

	"github.com/f1bonacc1/glippy"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	for range tick {
		clipboardCheck()
	}

}

func clipboardCheck() {
	var newCopy strings.Builder
	// get clipboard text
	text, err := glippy.Get()
	if err != nil {
		panic(err)
	}

	if !strings.Contains(text, "http") {
		return
	}
	if strings.Contains(text, "http://froob.org/") {
		return
	}
	words := strings.Fields(text)
	for i, s := range words {
		switch {
		case i == 0:
			switch {
			case strings.HasPrefix(s, "https://"):
				newCopy.WriteString("http://froob.org/")
			case strings.HasPrefix(s, "http://"):
				newCopy.WriteString("http://froob.org/")
			default:
				newCopy.WriteString(s)
			}
		default:
			switch {
			case strings.HasPrefix(s, "https://"):
				newCopy.WriteString(" http://froob.org/")
			case strings.HasPrefix(s, "http://"):
				newCopy.WriteString(" http://froob.org/")
			default:
				newCopy.WriteString(" " + s)
			}
		}
	}

	// set clipboard text
	glippy.Set(newCopy.String())
}
