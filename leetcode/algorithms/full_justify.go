package algorithms

import (
	"strings"
)

func FullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	lineLength := 0

	for _, word := range words {
		if lineLength+len(currentLine)+len(word) > maxWidth {
			result = append(result, justifyLine(currentLine, maxWidth, false))
		}
	}

	return result
}

func justifyLine(line []string, maxWidth int, isLastLine bool) string {
	if len(line) == 1 || isLastLine {
		joinedLine := strings.Join(line, " ")
		return joinedLine + strings.Repeat(" ", maxWidth-len(joinedLine))
	}

	wordsLength := len(strings.Join(line, ""))
	totalSpaces := maxWidth - wordsLength
	spaceBetweenWords := totalSpaces / (len(line) - 1)
	extraSpaces := totalSpaces % (len(line) - 1)

	var justifiedLine strings.Builder
	for i, word := range line {
		justifiedLine.WriteString(word)
		if i < len(line)-1 {
			justifiedLine.WriteString(strings.Repeat(" ", spaceBetweenWords))
			if i < extraSpaces {
				justifiedLine.WriteString(" ")
			}
		}
	}

	return justifiedLine.String()
}
