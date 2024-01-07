package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	templateless = "Project(XXX)"
	middleman    = "Project(YYY)"
)

func main() {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(templateless, middleman, false)

	var output string

	for i, diff := range diffs {
		if diff.Type == diffmatchpatch.DiffEqual {
			output += diff.Text
		} else if diff.Type == diffmatchpatch.DiffInsert && diffs[i-1].Type == diffmatchpatch.DiffDelete {
			output += "123"
		}
	}

	fmt.Println(output) // "Project(123)"
}
