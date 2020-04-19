package main

import (
	"fmt"

	"makarov.pw/prntscr/download"
	"makarov.pw/prntscr/parser"
	"makarov.pw/prntscr/randomizer"
)

func main() {
	path := randomizer.String(6)
	fmt.Println(path)
	//download.File("https://prnt.sc/"+path, path)
	if err := download.File("https://prnt.sc/"+path, path+".html"); err != nil {
		panic(err)
	}
	imgURL, ext := parser.Search(download.Code("https://prnt.sc/" + path))
	fmt.Println(imgURL, ext)
}
