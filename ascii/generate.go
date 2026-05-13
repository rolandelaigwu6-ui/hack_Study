package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	if strings.ReplaceAll(input, "\\n", "") == "" {
		return strings.Repeat("\n", len(input)/2)

	}
	lines := SplitInput(input)
	var output []string

	for i, line := range lines {
		if line == "" {
			if i == len(lines)-1 && input != `\n` {
				output = append(output, RenderLine("", banner)...)
			} else {
				output = append(output, "")
			}
			continue
		}
		render := RenderLine(line, banner)
		output = append(output, render...)
	}
	return strings.Join(output, "\n") + "\n"
}
