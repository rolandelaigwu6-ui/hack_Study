package main

import (
	"bufio"
	"fmt"
	"os"
)
func LoadBanner( filename string) (map[rune][]string, error) {
	if len(filename) == 0 {
		return nil, fmt.Errorf("filename is empty")
	}
	file,err := os.Open(filename) 
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v", err)
	}
	defer file.Close()
	var line []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %v", err)
	}
	if len(line) != 855 {
		return nil, fmt.Errorf("unexpected number of line got %d want 855", len(line))
	}
	banner := make(map[rune][]string)
	currentrune := rune(32)
	for i := 0; i < len(line); i += 9 {
		if i+8 > len(line) {
			return nil, fmt.Errorf("unsupported  rune %v at index %d", currentrune, i)
		}
		charblock := line[i+1 : i+9]
		banner[currentrune] = charblock
		currentrune++
	}
	if len(banner) != 95 {
		return nil, fmt.Errorf("unsupported char expected 95 got %v", len(banner))
	}
	return banner, nil

	
}