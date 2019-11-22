package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//buildSequencesFromFile takes in a source file, parses it, and returns an array of speciation sequences
func buildSequencesFromFile(source string) []*Sequence {
	sequences := []*Sequence{}

	//Read Source
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := strconv.Atoi(line); err == nil {
			if len(lines) > 0 {
				sequences = append(sequences, newSequence(lines))
				lines = []string{}
			}
		} else {
			lines = append(lines, line)
		}
	}

	return sequences
}

//utility function to save the data neatly into target file
func saveScraperData(file *os.File, data []string, index int) {
	file.WriteString(fmt.Sprintf("%d\n", index))
	for _, d := range data {
		file.WriteString(d + "\n")
	}
}

//utility function used to write a single line to the output file
func saveOutputData(file *os.File, line string) {
	file.WriteString(line + "\n")
}
