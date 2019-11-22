package engine

/*
type Sequence struct {
	Kingdom string
	Phylum  string
	Class   string
	Order   string
	Family  string
	Genus   string
	Species string
}

var fields = map[string]bool{
	"Kingdom": true,
	"Phylum":  true,
	"Class":   true,
	"Order":   true,
	"Family":  true,
	"Genus":   true,
	"Species": true,
}

func newSequence(strs []string) *Sequence {
	Seq := Sequence{}

	index := 0

	for index < len(strs) {
		if _, ok := fields[strs[index]]; ok {
			val := strs[index+2]

			//some species take the 2-word form "P. occidentalis"...this accounts for that...
			if val[len(val)-1] == []byte(".")[0] {
				fmt.Println(strs[index], val, strs[index+3])
			}

			index = index + 1
		}

		return &Seq
	}
	return nil
}

/*
func BuildTree(sanitized []string) map[string]string {
	sequences := []*Sequence{}

	for _, str := range sanitized {
		//parses into an array, of the form [ key : val key : val key : val ]
		strs := regexp.MustCompile("[^\\s]+").FindAllString(str, -1)

		fmt.Println()
		seq := newSequence(strs)

		sequences = append(sequences, seq)
	}

	return nil
}*/
