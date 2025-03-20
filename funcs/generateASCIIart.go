package funcs

import (
	"fmt"
)

// GenerateASCIIArt converts input into ASCII art based on the loaded banner
func GenerateASCIIArt(input string, banner map[rune][]string) ([]string, error) {
	var result []string
	lines := make([]string, 8)  // Temporary storage for ASCII art lines (each character is 8 lines tall)
	newLineEncountered := false // Track if we processed a newline

	for _, char := range input {
		if char == '\n' {
			// If there's content in lines, flush it to the result
			if len(lines[0]) > 0 {
				result = append(result, lines...)
				lines = make([]string, 8) // Reset lines
			}

			// Only add a blank line if a previous newline was already encountered
			if newLineEncountered {
				result = append(result, "")
			}

			newLineEncountered = true // Mark that a newline was processed to handle consecutive newlines correctly
			continue                  // Skip to the next character
		}

		// Reset tracker since we encountered a non-newline character
		newLineEncountered = false

		asciiArt, exists := banner[char]
		if !exists {
			return nil, fmt.Errorf("character '%c' is missing from the font definition", char)
		}

		// Must be exactly 8 lines to prevent index out of range errors
		if len(asciiArt) != 8 {
			return nil, fmt.Errorf("invalid ASCII art format: character '%c' has incorrect line count", char)
		}

		for i, artLine := range asciiArt {
			lines[i] += artLine + " " // Append the character's line with spacing
		}
	}

	// Append the last processed text if it's not empty
	if len(lines[0]) > 0 {
		result = append(result, lines...)
	}

	// Ensure the output ends with a blank line if the input ended with '\n'
	if len(input) > 0 && input[len(input)-1] == '\n' {
		result = append(result, "")
	}

	return result, nil // Return the generated ASCII art lines with no error
}
