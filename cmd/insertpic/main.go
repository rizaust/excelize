package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test/Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	if err := f.AddPictureFromURI("Sheet1", "A1", "http://127.0.0.1/a.png", nil); err != nil {
		log.Fatal(err)
	}
	if err := f.SaveAs("test/Book1_pic.xlsx"); err != nil {
		log.Fatal(err)
	}
}
