package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner reads the font file and maps characters to ASCII art
func LoadBanner(font string) (map[rune][]string, error) {
	filename := "fonts/" + font + ".txt"

	// Check if the file exists before trying to open it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Font file \"%s\" not found.", font)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open banner file \"%s\".", filename)
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var currentChar rune = ' ' // Tracks the current character being processed (starting with space)
	var charLines []string
	linesRead := 0 // Track the number of lines read to detect an empty file

	for scanner.Scan() {
		linesRead++ // Increment counter for each line read
		line := scanner.Text()

		if line == "" {
			if len(charLines) > 0 {
				bannerMap[currentChar] = append([]string{}, charLines...) // Store a copy of the character's ASCII representation
				currentChar++
				charLines = nil
			}
			continue
		}

		charLines = append(charLines, line) // Collect the ASCII art lines for the current character
	}

	// If no lines were read, return an error
	if linesRead == 0 {
		return nil, fmt.Errorf("Banner file \"%s\" is empty.", filename)
	}

	if len(charLines) > 0 {
		bannerMap[currentChar] = append([]string{}, charLines...) // Store the last character in the map
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to read banner file \"%s\": %s", filename, err)
	}

	return bannerMap, nil
}
