package house

import "fmt"

const testVersion = 1

var sub = []string{
	"",
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

var verb = []string{
	"",
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

func Song() string {
	var r string

	for v := 1; v < 13; v++ {
		if v != 1 {
			r += "\n\n"
		}
		r += Verse(v)
	}

	return r
}

func Verse(i int) string {
	v := fmt.Sprintf("This is %s", sub[i])

	if i > 1 {
		for ; i > 1; i-- {
			v += fmt.Sprintf("\nthat %s %s", verb[i-1], sub[i-1])
		}
	}

	return v
}
