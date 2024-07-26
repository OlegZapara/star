package main

import (
	"flag"
	"os"
)

func main() {
	search := flag.Bool("s", false, "search starred file or directory")
	delete := flag.Bool("d", false, "delete star from file or directory")
	editor := flag.String("e", "", "editor to open starred file")
	flag.Parse()
	file := openStar()
	defer file.Close()
	if *search {
		searchStar(file, SearchOptions{editor: *editor, delete: *delete})
	} else {
		arg := ""
		if len(os.Args) >= 2 {
			arg = os.Args[1]
		}
		toggleStar(file, arg)
	}
}
