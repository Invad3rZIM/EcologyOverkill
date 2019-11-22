package engine

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Sequence struct {
	classification map[string]string
}

//newSequence constructor maps an array of lines to a classification map
func newSequence(lines []string) *Sequence {
	classification := make(map[string]string)

	for _, line := range lines {
		if len(line) > 0 {
			li := strings.SplitN(line, " ", 2)
			classification[li[0]] = li[1]
		}
	}

	return &Sequence{
		classification: classification,
	}
}

/* RenderTreeFromFile reads from a source and writes to a target.
It will parse the source data, build a map of nestedSequences to DFS,
and then traverse said DFS while applying correct labelling / formatting
*/
func RenderTreeFromFile(source string, target string) {
	sequences := buildSequencesFromFile(source)

	nestedSequences := deriveRelationshipsSet(sequences)

	//create a new target to load data into
	f, err := os.Create(target)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	generateTree(nestedSequences, f, "Plantae", 0)

}

/* DFS to traverse the tree...

 */

var count = 0

func generateTree(speciesTable map[string]map[string]struct{}, file *os.File, start string, depth int) {
	//pre-header-spacing...

	line := ""

	for i := 0; i < (depth); i += 1 {
		line = line + "    "
	}

	switch depth { //output a different header depending on the recursion level
	case 0:
		line = line + "Kingdom:  "
	case 1:
		line = line + "Phylum:   "
	case 2:
		line = line + "Class:    "
	case 3:
		line = line + "Order:    "
	case 4:
		line = line + "Family:   "
	}

	//more spacing...
	for i := 0; i < (5 - depth); i += 1 {
		line = line + "    "
	}

	line = line + start

	saveOutputData(file, line)

	for cat, _ := range speciesTable[start] {
		//Handle single edge case... not worth rewriting this program to fix...
		/*
			This single case needs to be handled because in 1 instance, Class names are the same
			For different Phylums (PhylumA:CCC vs PhylumB:CCC) and so this graph has 1 path that needs
			to be killed. Because this program is simply to help me with my termpaper, I'm choosing to
			leave this edgecase in with an examplanation as to why it's needed...
		*/
		if start == "Dicotyledonae" && cat == "Asterales" {
			continue
		}

		if depth < 4 {
			generateTree(speciesTable, file, cat, depth+1)
		} else { //species is depth=4, and so this becomes the base case
			count = count + 1
			line = fmt.Sprint("                                        ", count, ". ", cat)

			saveOutputData(file, line)
		}
	}
}

/*	Generates an adjacency list of categories to graphically traverse...
 */
func deriveRelationshipsSet(sequences []*Sequence) map[string]map[string]struct{} {
	relationshipSet := make(map[string]map[string]struct{})

	//bind parent and child together
	order := map[string]int{"Kingdom": 0, "Phylum": 1, "Class": 2, "Order": 3, "Family": 4, "Species": 5}
	inverse := map[int]string{0: "Kingdom", 1: "Phylum", 2: "Class", 3: "Order", 4: "Family", 5: "Species"}

	for _, seq := range sequences {
		for childKey, _ := range seq.classification {
			val := order[childKey] - 1

			if val >= 0 {
				parentKey := inverse[val]
				parent := seq.classification[parentKey]
				child := seq.classification[childKey]

				if _, ok := relationshipSet[parent]; !ok {
					relationshipSet[parent] = make(map[string]struct{})
				}

				relationshipSet[parent][child] = struct{}{}
			}
		}
	}

	return relationshipSet
}
