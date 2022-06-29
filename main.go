package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/jung-kurt/gofpdf"
	
)


func main(){
	jsonFile, err_r := os.Open("grade.json")
	if err_r != nil {
        fmt.Println("error")
    }
    defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var Record report
	json.Unmarshal(byteValue, &Record)



	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	header(pdf)

	



	err := pdf.OutputFileAndClose("pdfs/hello.pdf")
	if err != nil{
		fmt.Println("error")
	}
}

