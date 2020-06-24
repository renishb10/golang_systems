package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
		return
	}

	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("Cannot get the absolute path")
		return
	}

	fmt.Println("Abs", root)
	fmt.Println("Base", filepath.Base(root)) // Gives the Last directory
	fmt.Println("Clean", filepath.Clean(root + "//feed.go")) // Cleans up the URL
	fmt.Println("Dir", filepath.Dir(root + "/feed.go")) // Gives the last parent directory
	fmt.Println("Ext", filepath.Ext(root + "/feed.go")) // This gets the extension
	fmt.Println(filepath.Glob("go*")) // This finds the matching files in the path
	fmt.Println("IsAbs", filepath.IsAbs(root)) // Checks if the path is a Abs or not
	fmt.Println("Join", filepath.Join("/Documents", "feed", "renish")) // Just joins with OS specific seperator
	fmt.Println(filepath.Match("go*", "gomod.go")) // Matches the name with the pattern
	fmt.Println(filepath.Split(root)) // Splits the file or the last unit
	fmt.Println(filepath.SplitList(root)) // Splits based on separator NOT / its : or ;

	var c struct {
		files int
		dirs int
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})

	fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)
}