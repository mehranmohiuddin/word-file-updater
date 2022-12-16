package main

import (
	"github.com/nguyenthenguyen/docx"
)

func main() {
	r, err := docx.ReadDocxFile("./template.docx")
	if err != nil {
		panic(err)
	}

	docx1 := r.Editable()
	docx1.Replace("OLD_TEXT", "NEW_TEXT", -1)
	docx1.WriteToFile("./new_result.docx")

	r.Close()
}
