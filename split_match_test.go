package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitMatch(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		arg     string
		counts  map[string]int
	}{
		{
			name:    "Parenthesis pattern match",
			content: []byte("dog: 12\ndog: 13\n"),
			arg:     "dog: (\\d+)",
			counts: map[string]int{
				NO_MATCH: 0,
				"12":     1,
				"13":     1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "example")
			if err != nil {
				log.Fatal(err)
			}

			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(test.content); err != nil {
				log.Fatal(err)
			}

			if _, err := tmpfile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }() // Restore original Stdin

			os.Stdin = tmpfile
			counts, err := splitMatch(test.arg, false)
			if err != nil {
				t.Fatalf("Expected nil, got %v", err)
			}

			assert.Equal(t, test.counts, counts)

			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}
		})
	}
}
