package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

func splitMatch(arg string, logNonMatches bool) (map[string]int, error) {
	reader := bufio.NewReader(os.Stdin)

	var total, maxLen int

	lastPrint := time.Now()
	counts := make(map[string]int)
  counts[NO_MATCH] = 0

	reg := regexp.MustCompile(arg)

	writeOutput := func() {
		if total == 0 {
			return
		}

		log.Println("\n\nMatches:")

		var maxValLen int
		for _, v := range counts {
			x := len(fmt.Sprintf("%d", v))
			if x > maxValLen {
				maxValLen = x
			}
		}

		pct := float64(counts[NO_MATCH]) / float64(total) * 100
		log.Printf("%"+fmt.Sprintf("%d", maxLen+1)+"s: %s%-"+fmt.Sprintf("%d", maxValLen)+"d%s %s(%.0f%%)%s\n", NO_MATCH, Red, counts[NO_MATCH], Reset, Gray, pct, Reset)

		type Line struct {
			Sort  int
			Value string
		}

		var lines []Line

		for k, v := range counts {
			if k == NO_MATCH {
				continue
			}

			col := Reset
			if v > 0 {
				col = Cyan
			}

			pct := float64(v) / float64(total) * 100

			lines = append(
				lines,
				Line{
					Sort: v,
					Value: fmt.Sprintf("%s%"+fmt.Sprintf("%d", maxLen+1)+"s%s: %s%-"+fmt.Sprintf("%d", maxValLen)+"d%s %s(%.0f%%)%s",
						Yellow,
						k,
						Reset,
						col,
						v,
						Reset,
						Gray,
						pct,
						Reset,
					),
				})
		}

		sort.Slice(lines, func(i, j int) bool {
			return lines[i].Sort > lines[j].Sort
		})

		for _, line := range lines {
			log.Println(line.Value)
		}
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		matches := reg.FindStringSubmatch(line)

		if matches == nil || len(matches) < 2 {
			counts[NO_MATCH]++
			total++

			if logNonMatches {
				log.Printf("%s[No Match]:%s: %s\n", Red, Reset, line)
			}
		} else {
			counts[matches[1]]++
			total++

			if len(matches[1]) > maxLen {
				maxLen = len(matches[1])
			}
		}

		if time.Since(lastPrint) > time.Second*5 {
			lastPrint = time.Now()
			writeOutput()
		}
	}

	writeOutput()

	return counts, nil
}
