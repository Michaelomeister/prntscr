package main

import (
	"fmt"

	"makarov.pw/prntscr/download"
	"makarov.pw/prntscr/parser"
	"makarov.pw/prntscr/randomizer"
)

func main() {
	path := randomizer.String(6)
	//fmt.Println(path)
	//s := "https://prnt.sc/"
	//download.File(s+path, path+".html")
	/*if err := download.File("https://prnt.sc/"+path, path+".html"); err != nil {
		panic(err)
	}*/
	imgURL, ext := parser.Search(download.Code("https://prnt.sc/" + path))
	//fmt.Println(imgURL, ext)
	if imgURL[0] == 47 {
		fmt.Println("nihuya")
	} else if ext == "peg" {
		download.File(imgURL, path+".jpeg")
		fmt.Println(path + ".jpeg")
	} else {
		download.File(imgURL, path+"."+ext)
		fmt.Println(path + ".jpeg")
	}
}
