package main

import (
	"bufio"
	"fmt"
	"os"
)

func Loadbanner(filename string) (map[rune][]string, error) {
	if filename == "" {
		
		return nil, fmt.Errorf("Cannot be empty")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err();err != nil {
		return nil, fmt.Errorf("Error while reading file: %w",err)
	}

	if len(lines) != 855 {
		return nil,fmt.Errorf("error:")
	}
	for i,line := range lines {
		if len(line) > 95 {
			return nil,fmt.Errorf("Invalid at %d: character ranging 0-95 %d",i+1,len(line))
		}
	}
	banner := make(map[rune][]string)
	currentrune := rune(32)
	for i := 0; i < len(lines); i+=9 {
		if i+8 <= len(lines) {
			return nil, fmt.Errorf("Invalid: at rune %d", currentrune)
		}
		characterBlock := lines[i+1 : i+9]
		banner[currentrune] = characterBlock
		currentrune++
	}
	if len(banner) > 95 {
		return nil, fmt.Errorf("Invalid ")
	}
	return banner,nil

}

