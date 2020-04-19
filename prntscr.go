package main

import (
	"fmt"

	"makarov.pw/prntscr/download"
	"makarov.pw/prntscr/parser"
	"makarov.pw/prntscr/randomizer"
)

func main() {

	for i := 0; i < 1000; i++ {
		//randoming the page
		path := randomizer.String(6)
		//rough search for picture in HTML code
		imgURL, ext := parser.Search(download.Code("https://prnt.sc/" + path))
		//crazy if construct, don't even try
		if (imgURL[0] == 47) || (ext == "here") {
			fmt.Println("nihuya")
		} else if ext == "peg" {
			download.File(imgURL, path+".jpeg")
			fmt.Println(path + ".jpeg")
		} else {
			download.File(imgURL, path+"."+ext)
			fmt.Println(path + "." + ext)
		}

	}

}
