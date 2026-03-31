package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func process(line string) (string, bool) {
	line = strings.TrimSpace(line)

	if line == "" || strings.Trim(line, "-") == "" {
		return "", false
	}

	if strings.HasPrefix(strings.ToUpper(line), "TODO:") {
		line = "✦ ACTION:" + line[5:]
	}

	if strings.HasPrefix(strings.ToLower(line), "TODO:") {
		line = "✦ ACTION:" + line[5:]
	}

	if strings.HasPrefix(strings.ToUpper(line), "CLASSIFIED:") {
		line = "[REDACTED]:" + line[11:]
	}

	if strings.HasPrefix(strings.ToLower(line), "CLASSIFIED:") {
		line = "[REDACTED]:" + line[11:]
	}

	isUpper := false
	for _, ch := range line {
		if unicode.IsLetter(ch) {
			isUpper = true
			if unicode.IsLower(ch) {
				isUpper = false
				break
			}
		}
	}
	if isUpper {
		line = strings.Title(strings.ToLower(line))
	}

	return line, true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer inFile.Close()

	outFile, _ := os.Create(os.Args[2])
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	var count, removed int

	fmt.Fprintln(writer, "SENTINEL REPORT")
	fmt.Fprintf(writer, "Processed: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	for scanner.Scan() {
		line, ok := process(scanner.Text())
		if !ok {
			removed++
			continue
		}
		count++
		fmt.Fprintf(writer, "%03d. %s\n", count, line)
	}

	fmt.Fprintf(writer, "\nSummary:\nRead: %d\nWritten: %d\nRemoved: %d\n",
		count+removed, count, removed)

	writer.Flush()
}
