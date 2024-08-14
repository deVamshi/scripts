package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// used to create my files using format dd-MM-yyyy.md
// but wanted to have yyyy-MM-dd.md format for proper ordering reasons
func main() {
	// dir in which the files are to be renamed
	dir := `/Users/devamshi/Documents/repos/scripts/test_in_dir/`

	items, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for i, item := range items {

		oldname := item.Name()
		splits := strings.Split(oldname, "-")
		day, month, yr := splits[0], splits[1], splits[2]

		if len(day) != 2 {
			// if files is already following the desired format, then skip this
			continue
		}

		yr, _ = strings.CutSuffix(yr, ".md")

		newName := strings.Join([]string{yr, month, day}, "-") + ".md"

		err := os.Rename(filepath.Join(dir, oldname), filepath.Join(dir, newName))

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(i, ". oldname: ", oldname, "-----", "new name:", newName)

	}

}
