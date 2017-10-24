package main

import (

	"fmt"
	"github.com/jung-kurt/gofpdf"
)



func main() {
	fmt.Println("hello world")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	if err:= pdf.OutputFileAndClose("hello.pdf"); err != nil {
		fmt.Println("Erreur pdt generation du pdf : ", err)
	}
}