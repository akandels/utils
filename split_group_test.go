package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitGroup(t *testing.T) {
	tests := []struct {
    name string
		content         []byte
		args            []string
		allowMultiMatch bool
		counts          map[string]int
	}{
		{
      name: "Parenthesis pattern match",
			content:         []byte("I am a red dog\n"),
			args:            []string{"^yellow", "(red|green)", "blu.*"},
			allowMultiMatch: false,
			counts: map[string]int{
				NO_MATCH:      0,
				"^yellow":     0,
				"(red|green)": 1,
				"blu.*":       0,
			},
		},
		{
      name: "Case insensitive starting line regex",
			content:         []byte("Yellow is my color\n"),
			args:            []string{"^yellow", "(red|green)", "blu.*"},
			allowMultiMatch: false,
			counts: map[string]int{
				NO_MATCH:      0,
				"^yellow":     1,
				"(red|green)": 0,
				"blu.*":       0,
			},
		},
		{
      name: "No match",
			content:         []byte("Purple\n"),
			args:            []string{"^yellow", "(red|green)", "blu.*"},
			allowMultiMatch: false,
			counts: map[string]int{
				NO_MATCH:      1,
				"^yellow":     0,
				"(red|green)": 0,
				"blu.*":       0,
			},
		},
		{
      name: "Multi match",
			content:         []byte("Yellow is key\nRed is next best\nI am blue\nPurple too!\n"),
			args:            []string{"^yellow", "(red|green)", "blu.*"},
			allowMultiMatch: false,
			counts: map[string]int{
				NO_MATCH:      1,
				"^yellow":     1,
				"(red|green)": 1,
				"blu.*":       1,
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
			counts, err := splitGroup(test.args, test.allowMultiMatch, false)
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
