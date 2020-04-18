package main

import (
	"fmt"
	"makarov.pw/prntscr/randomizer"
	"makarov.pw/prntscr/download"
)

func main() {
	path := randomizer.String(6)
	fmt.Println(path)
	//download.File("https://prnt.sc/"+path, path)
	if err := download.File("https://prnt.sc/"+path, path+".html"); err != nil {
		        panic(err)
	}
}
