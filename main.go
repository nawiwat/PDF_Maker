package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"github.com/jung-kurt/gofpdf"
)

//Globe value
var Record report

func main(){
	jsonFile, err_r := os.Open("grade.json")
	if err_r != nil {
        fmt.Println("error")
    }
    defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Record)

	//Values
	firstPagePosX,firstPagePosY := 10.0 , 46.0
	NextPagePosX,NextPagePosY := 10.0, 10.0
	TableSpace := 250.0
	PagePosX,PagePosY := firstPagePosX,firstPagePosY
	CurrentSite := "Left"
	FirstPage := true
	TableSize := 20.0

	//Pdf Contains
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(true,0)
	pdf.AddPage()
	Header(pdf)
	GradeHeader(pdf)
	PagePosY = 54.0

	//Main Table Loop
	pdf.SetFont("Arial","B",7)
	for i := 0 ; i < len(Record.Report) ; i++ {
		pdf.SetXY(PagePosX,PagePosY)
		TableSize = TableSpaceCal(i)
		if TableSpace - PagePosY < TableSize {
			if FirstPage == true {
				PagePosX,PagePosY = firstPagePosX + 100 , firstPagePosY
				pdf.SetXY(PagePosX,PagePosY)
				CurrentSite = "Right"
				FirstPage = false
			} else {
				if CurrentSite == "Left" {
					PagePosX,PagePosY = NextPagePosX + 100 ,NextPagePosY
					pdf.SetXY(PagePosX,PagePosY)
					CurrentSite = "Right"
				} else {
					Footer(pdf)
					PagePosX,PagePosY = NextPagePosX,NextPagePosY
					pdf.SetXY(PagePosX,PagePosY)
					CurrentSite = "Left"
					pdf.AddPage()
				}
			}
			
		}

		BuildTable(pdf,i,PagePosX,PagePosY)

		pdf.Line(PagePosX,PagePosY + TableSize - 4 ,PagePosX + 90 ,PagePosY + TableSize - 4 )

		PagePosY += TableSize
	}
	
	Gradefooter(pdf,PagePosX,PagePosY)
	Footer(pdf)

	err := pdf.OutputFileAndClose("pdfs/hello.pdf")
	if err != nil{
		fmt.Println("error")
	}
}

//Import grade info
type report struct {
	Report []struct{
		Grades []struct{
			CourseCode 	string 	`json:"courseCode"`
			Final_score string 	`json:"final_score"`
			Points 		float64 	`json:"points"`
			Short_title string 	`json:"short_title"`
			Title 		string 	`json:"title"`
			Units 		int 	`json:"units"`
		} 	`json:"grades"`
		Semester 	string 	`json:"semester"`
		Summary struct{
			Unitspassed float64 	`json:"units"`
			UnitsFact 	float64		`json:"units_fac"`
			FinalQpa 	float64 	`json:"qpa"`
			TotalPoints float64		`json:"points"`
		}	`json:"summary"`
		Cumulative struct{
			Unitspassed float64 	`json:"units"`
			UnitsFact 	float64		`json:"units_fac"`
			FinalQpa 	float64 	`json:"qpa"`
			TotalPoints float64		`json:"points"`
		}	`json:"cumulative"`
	} 	`json:"report"`
}


//Pdf/Grade header & footer
func Header(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetFont("Arial","B",9)
	pdf.Cell(200,3,"Officail Academic Record as of 20 Jun 2022")
	pdf.Ln(5)
	pdf.Cell(100,3 ,"Student Name: Omega Hub Student")
	pdf.Cell(100,3,"Student ID: omega.student@cmkl.ac.th")
	pdf.Line(10,19,200,19)
	pdf.Ln(6)
	pdf.Cell(20,3,"University: ")
	pdf.SetFont("Arial","",9)
	pdf.Cell(150,3,"CMKL University")
	pdf.Ln(5)
	pdf.SetFont("Arial","B",9)
	pdf.Cell(20,3,"Degree: ")
	pdf.SetFont("Arial","",9)
	pdf.Cell(150,3,"Master of Science")
	pdf.Ln(5)
	pdf.SetFont("Arial","B",9)
	pdf.Cell(20,3,"Program: ")
	pdf.SetFont("Arial","",9)
	pdf.Cell(150,3,"M.S. in Electrical and Computer Engeering")
	pdf.Line(10,36,200,36)
	pdf.Ln(6)
	pdf.SetFont("Arial","B",9)
	pdf.Cell(200,5,"Carnegie Mellon University - CMKL | THAILAND")
	pdf.Line(10,43,200,43)
	pdf.Ln(6)

	return pdf
}
func Footer(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetXY(5,287)
	pdf.SetFont("Arial","",9)
	pdf.Cell(200,3,"This information is not intended for external distribution.")
	pdf.ImageOptions("images/PresSign.JPG",160,268,35,12,false,gofpdf.ImageOptions{ImageType: "JPG",ReadDpi: true},0,"")
	pdf.SetXY(150,282)
	pdf.SetTextColor(0,0,255)
	pdf.Cell(50,3,"Supan Tungjitkusolmun, President")
	pdf.SetTextColor(0,0,0)
	return pdf
}
func GradeHeader(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetXY(10,46)
	pdf.SetFont("Arial","B",9)
	pdf.SetFillColor(128,128,128)
	pdf.CellFormat(90,6,"Beginning of Graduate Record","",0,"CM",true,0,"")
	return pdf
}
func Gradefooter(pdf *gofpdf.Fpdf , PagePosX float64 , PagePosY float64) *gofpdf.Fpdf{
	pdf.SetXY(PagePosX,PagePosY)
	pdf.SetFont("Arial","B",9)
	pdf.SetFillColor(128,128,128)
	pdf.CellFormat(90,6,"End of Graduate Record","",0,"CM",true,0,"")
	return pdf
}

//Tables
func TableSpaceCal(i int) float64{
	Space := float64(40 + (14 * len(Record.Report[i].Grades)))
	return Space
}

func BuildTable(pdf *gofpdf.Fpdf , i int ,PosX float64 , PosY float64) *gofpdf.Fpdf{
	TEXT := [][]byte{}

	//Semester Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",8)
	pdf.CellFormat(90,4, Record.Report[i].Semester ,"",0,"LM",false,0,"")
	PosY += 6

	//Geade Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(14,4, "PROGRAM" ,"",0,"CM",false,0,"")
	pdf.CellFormat(10,4, "CRS#" ,"",0,"CM",false,0,"")
	pdf.CellFormat(30,4, "COURSE TITLE" ,"",0,"CM",false,0,"")
	pdf.CellFormat(12,4, "UNITS" ,"",0,"CM",false,0,"")
	pdf.CellFormat(12,4, "GRADE" ,"",0,"CM",false,0,"")
	pdf.CellFormat(12,4, "POINTS" ,"",0,"CM",false,0,"")
	PosY += 6

	//Grade Rows
	for j := 0 ; j < len(Record.Report[i].Grades); j++ {
		pdf.SetFont("Arial","",6)
		pdf.SetXY(PosX,PosY)
		pdf.CellFormat(14,14, "ECE" ,"",0,"CM",false,0,"")

		pdf.SetXY(PosX + 14,PosY)
		TEXT = pdf.SplitLines([]byte(Record.Report[i].Grades[j].CourseCode),10)
		for _,lN := range TEXT {
			pdf.CellFormat(10, float64(14/len(TEXT)) ,string(lN),"", int(14/len(TEXT)) ,"CM",false,0,"")
		}
		
		pdf.SetXY(PosX + 24,PosY)
		TEXT = pdf.SplitLines([]byte(Record.Report[i].Grades[j].Title),30)
		for _,lN := range TEXT {
			pdf.CellFormat(30, float64(14/len(TEXT)) ,string(lN),"", int(14/len(TEXT)) ,"CM",false,0,"")
		}

		pdf.SetXY(PosX + 54,PosY)
		pdf.CellFormat(12,14, strconv.Itoa(Record.Report[i].Grades[j].Units) ,"",0,"CM",false,0,"")

		if Record.Report[i].Grades[j].Final_score != ""{
			pdf.SetXY(PosX + 66,PosY)
			pdf.CellFormat(12,14, Record.Report[i].Grades[j].Final_score ,"",0,"CM",false,0,"")
		}else{
			pdf.SetXY(PosX+66,PosY)
			pdf.CellFormat(12,14, "TBA" ,"",0,"CM",false,0,"")
		}
		
		pdf.SetXY(PosX + 78,PosY)
		pdf.CellFormat(12,14, strconv.FormatFloat(Record.Report[i].Grades[j].Points,'f',2,32) ,"",0,"CM",false,0,"")

		PosY += 14
	}
	PosY += 3

	//Summary Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(22,3, "" ,"",0,"CB",false,0,"")
	pdf.CellFormat(17,3, "UNITS" ,"",0,"CB",false,0,"")
	pdf.CellFormat(17,3, "UNITS" ,"",0,"CB",false,0,"")
	pdf.CellFormat(17,3, "FINAL" ,"",0,"CB",false,0,"")
	pdf.CellFormat(17,3, "TOTAL" ,"",0,"CB",false,0,"")
	PosY += 5
	pdf.SetXY(PosX,PosY)
	pdf.CellFormat(22,3, "" ,"",0,"CB",false,0,"")
	pdf.CellFormat(17,3, "PASSED" ,"",0,"CT",false,0,"")
	pdf.CellFormat(17,3, "FACTORABLE" ,"",0,"CT",false,0,"")
	pdf.CellFormat(17,3, "QPA" ,"",0,"CT",false,0,"")
	pdf.CellFormat(17,3, "POINTS" ,"",0,"CT",false,0,"")
	PosY += 5

	//Sumary Rows

	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(22,3, "Semester" ,"",0,"CM",false,0,"")
	pdf.SetFont("Arial","",6)
	pdf.CellFormat(17,3, strconv.FormatFloat( Record.Report[i].Summary.Unitspassed,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Summary.UnitsFact,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Summary.FinalQpa,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Summary.TotalPoints,'f',2,32) ,"",0,"CM",false,0,"")
	PosY += 5
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(22,3, "Cumulative" ,"",0,"CB",false,0,"")
	pdf.SetFont("Arial","",6)
	pdf.CellFormat(17,3, strconv.FormatFloat( Record.Report[i].Cumulative.Unitspassed,'f',2,64) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Cumulative.UnitsFact,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Cumulative.FinalQpa,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(17,3, strconv.FormatFloat(Record.Report[i].Cumulative.TotalPoints,'f',2,32) ,"",0,"CM",false,0,"")
	PosY += 5

	return pdf
}