package parser

import (
	"regexp"
)

func Search(input string) (string, string) {
	queryBegin, _ := regexp.Compile("<meta name=\"twitter:image:src\" content=")
	foundBegin := queryBegin.FindStringIndex(input)
	queryEnd, _ := regexp.Compile("<meta property=")
	foundEnd := queryEnd.FindStringIndex(input)
	runes := []rune(input)
	return string(runes[foundBegin[1]+1 : foundEnd[0]-4]), string(runes[foundEnd[0]-7 : foundEnd[0]-4])
}
