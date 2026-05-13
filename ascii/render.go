package main

import(

)


func RenderLine(text string, banner map[rune][]string) []string {
	var result []string
	for i := 0; i < 8; i++ {
		word := ""
		for _, char := range text {
			word += banner[char][i]
		}
		result = append(result, word)
	}
	return result
}