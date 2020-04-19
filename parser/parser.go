package parser

import (
	"regexp"
)

//Search for pattern before the img link and after, grabbing indexes, doing some math
func Search(input string) (string, string) {
	queryBegin, _ := regexp.Compile("<meta name=\"twitter:image:src\" content=")
	foundBegin := queryBegin.FindStringIndex(input)
	queryEnd, _ := regexp.Compile("<meta property=")
	foundEnd := queryEnd.FindStringIndex(input)
	runes := []rune(input)
	if (foundBegin != nil) && (foundEnd != nil) {
		return string(runes[foundBegin[1]+1 : foundEnd[0]-4]), string(runes[foundEnd[0]-7 : foundEnd[0]-4])
	} else {
		return "nothing", "here"
	}
}
