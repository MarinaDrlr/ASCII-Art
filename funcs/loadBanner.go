package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner reads the font file and maps characters to ASCII art
func LoadBanner(font string) (map[rune][]string, error) {
	filename := "fonts/" + font + ".txt"

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("font file \"%s\" not found", font)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open banner file \"%s\"", filename)
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var currentChar rune = ' ' // Tracks the current character being processed
	var charLines []string
	linesRead := 0 // Total lines read

	for scanner.Scan() {
		linesRead++
		line := scanner.Text()

		if line == "" {
			if len(charLines) > 0 {
				// Check if the character has exactly 8 lines
				if len(charLines) != 8 {
					return nil, fmt.Errorf("banner file \"%s\" is corrupted", filename)
				}
				bannerMap[currentChar] = append([]string{}, charLines...)
				currentChar++
				charLines = nil
			}
			continue
		}

		charLines = append(charLines, line)
	}

	// Final check: Last character must also have exactly 8 lines
	if len(charLines) > 0 && len(charLines) != 8 {
		return nil, fmt.Errorf("banner file \"%s\" is corrupted", filename)
	}

	// Check if the file was completely empty
	if linesRead == 0 {
		return nil, fmt.Errorf("banner file \"%s\" is empty", filename)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read banner file \"%s\": %s", filename, err)
	}

	return bannerMap, nil
}
