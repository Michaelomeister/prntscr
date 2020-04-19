package main

import (
	"fmt"
	"runtime"

	"makarov.pw/prntscr/download"
	"makarov.pw/prntscr/parser"
	"makarov.pw/prntscr/randomizer"
)

func main() {
	path := randomizer.String(6)
	fmt.Println(path)
	PrintMemUsage()
	//s := "https://prnt.sc/"
	//download.File(s+path, path+".html")
	/*if err := download.File("https://prnt.sc/"+path, path+".html"); err != nil {
		panic(err)
	}*/
	imgURL, ext := parser.Search(download.Code("https://prnt.sc/" + path))
	PrintMemUsage()
	fmt.Println(imgURL, ext)
	download.File(imgURL, path+"."+ext)
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
