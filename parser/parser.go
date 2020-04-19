package parser

import (
	"regexp"
)

func Search(input string) (string, string) {
	queryBegin, _ := regexp.Compile("<meta name=\"twitter:image:src\" content=")
	foundBegin := queryBegin.FindStringIndex(input)
	queryEnd, _ := regexp.Compile("/>\n <meta property=")
	foundEnd := queryEnd.FindStringIndex(input)
	runes := []rune(input)
	return string(runes[foundBegin+2 : foundEnd-2]), string(runes[foundEnd-5 : foundEnd-2])
}
