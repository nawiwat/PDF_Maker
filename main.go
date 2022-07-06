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

//Main
func main(){
	//Import Json Values
	jsonFile, err_r := os.Open("grade.json")
	if err_r != nil {
        fmt.Println("error")
    }
    defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Record)

	//Values
	firstPagePosX,firstPagePosY := 7.0 , 51.0
	TableSpace := 250.0
	PagePosX,PagePosY := firstPagePosX,firstPagePosY
	CurrentSite := "Left"
	FirstPage := true
	TableSize := 20.0

	//Pdf Contains
	pdf := gofpdf.New("P", "mm", "A4", "")//210 x 297 mm
	pdf.SetMargins(0,0,0)
	pdf.SetAutoPageBreak(true,0)
	pdf.AddPage()
	BgAdd(pdf)
	Header(pdf)
	GradeHeader(pdf)
	PagePosY = 58.0

	//Main Table Loop
	pdf.SetFont("Arial","B",7)
	for i := 0 ; i < len(Record.Report) ; i++ {
		//Position Set
		pdf.SetXY(PagePosX,PagePosY)
		TableSize = TableSpaceCal(i , pdf)

		//Page Cases
		if TableSpace - PagePosY < TableSize {
			if FirstPage == true {
				PagePosX,PagePosY = firstPagePosX + 100 , firstPagePosY
				pdf.SetXY(PagePosX,PagePosY)
				CurrentSite = "Right"
				FirstPage = false
			} else {
				if CurrentSite == "Left" {
					PagePosX,PagePosY = firstPagePosX + 100 ,firstPagePosY
					pdf.SetXY(PagePosX,PagePosY)
					CurrentSite = "Right"
				} else {
					//Footer(pdf)
					PagePosX,PagePosY = firstPagePosX,firstPagePosY
					pdf.SetXY(PagePosX,PagePosY)
					CurrentSite = "Left"
					pdf.AddPage()
					BgAdd(pdf)
					Header(pdf)
				}
			}
		}

		//Build Table
		BuildTable(pdf,i,PagePosX,PagePosY)
		pdf.Line(PagePosX,PagePosY + TableSize -2 ,PagePosX + 96 ,PagePosY + TableSize-2)
		PagePosY += TableSize
	}

	//Grade footer & Last page footer
	Gradefooter(pdf,PagePosX,PagePosY)
	//Footer(pdf)

	//info page
	pdf.AddPage()
	pdf.ImageOptions("images/Info.jpg",0,0,210,297,false,gofpdf.ImageOptions{ImageType: "jpg",ReadDpi: true},0,"")

	//Pdfs output
	err := pdf.OutputFileAndClose("pdfs/hello.pdf")
	if err != nil{
		fmt.Println("error")
	}
}

//Functions

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

//Background
func BgAdd(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.ImageOptions("images/Bg1.jpg",0,0,210,297,false,gofpdf.ImageOptions{ImageType: "jpg",ReadDpi: true},0,"")
	return pdf
}


//Pdf/Grade header & footer
func OldHeader(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetFont("Arial","B",9)
	pdf.Cell(200,3,"Official Academic Record as of 20 Jun 2022")
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

func Header(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetFont("Arial","",7)
	pdf.SetXY(5,30)
	pdf.Cell(60,4,"Official Academic Record as of")
	pdf.SetXY(5,35)
	pdf.Cell(60,4,"Student Name: ")
	pdf.SetXY(5,39)
	pdf.Cell(60,4,"Student ID: ")

	pdf.SetXY(40,30)
	pdf.Cell(60,4,"20 Jun 2022")
	pdf.SetXY(22,35)
	pdf.Cell(60,4,"Omega Hub Student")
	pdf.SetXY(18,39)
	pdf.Cell(60,4,"Omega.student@cmkl.ac.th")

	pdf.SetXY(84,30)
	pdf.SetFont("Arial","B",7)
	pdf.CellFormat(120,4,"Electrical and Computer Engineering","",0,"R",false,0,"")
	CellMargin := pdf.GetStringWidth("Electrical and Computer Engineering")
	pdf.SetXY(84,30)
	pdf.SetFont("Arial","",7)
	pdf.CellFormat(120 - CellMargin,4,"Program: ","",0,"R",false,0,"")

	pdf.SetXY(84,34)
	pdf.SetFont("Arial","B",7)
	pdf.CellFormat(120,4,"Master of Science","",0,"R",false,0,"")
	CellMargin = pdf.GetStringWidth("Master of Science")
	pdf.SetXY(84,34)
	pdf.SetFont("Arial","",7)
	pdf.CellFormat(120 - CellMargin,4,"Degree: ","",0,"R",false,0,"")

	pdf.SetXY(84,38)
	pdf.SetFont("Arial","B",7)
	pdf.CellFormat(120,4,"31 Aug 2020","",0,"R",false,0,"")
	CellMargin = pdf.GetStringWidth("31 Aug 2020")
	pdf.SetXY(84,38)
	pdf.SetFont("Arial","",7)
	pdf.CellFormat(120 - CellMargin,4,"Enrolled Date: ","",0,"R",false,0,"")



	pdf.SetXY(84,42)
	pdf.SetFont("Arial","B",7)
	pdf.CellFormat(120,4,"31 May 2020","",0,"R",false,0,"")
	CellMargin = pdf.GetStringWidth("31 May 2020")
	pdf.SetXY(84,42)
	pdf.SetFont("Arial","",7)
	pdf.CellFormat(120 - CellMargin,4,"Anticipated Graduation Date: ","",0,"R",false,0,"")

	pdf.SetFont("Arial","B",8)
	pdf.SetXY(5,44)
	pdf.CellFormat(80,5,"Carnegie Mellon University - CMKL | THAILAND","",0,"L",false,0,"")

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
	pdf.SetXY(6,51)
	pdf.SetFont("Arial","B",6)
	pdf.SetFillColor(192,192,192)
	pdf.CellFormat(96,4,"BEGINNING OF ACADEMIC RECORD","",0,"CM",true,0,"")
	return pdf 
}

func Gradefooter(pdf *gofpdf.Fpdf , PagePosX float64 , PagePosY float64) *gofpdf.Fpdf{
	pdf.SetXY(PagePosX,PagePosY)
	pdf.SetFont("Arial","B",6)
	pdf.SetFillColor(192,192,192)
	pdf.CellFormat(96,4,"END OF ACADEMIC RECORD","",0,"CM",true,0,"")
	return pdf
}


//Tables
func TableSpaceCal(i int , pdf *gofpdf.Fpdf) float64{
	Lines := 0
	for  j := 0;  j < len(Record.Report[i].Grades) ;  j++ {
		TEXT := pdf.SplitLines([]byte(Record.Report[i].Grades[j].Title),40)
		if len(TEXT) > 2{
			Lines += len(TEXT)-2
		}
	}
	Space := float64(29 + (7 * len(Record.Report[i].Grades)) + (3*Lines))
	return Space
}

func BuildTable(pdf *gofpdf.Fpdf , i int ,PosX float64 , PosY float64) *gofpdf.Fpdf{
	//Semester Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",7)
	pdf.CellFormat(90,3, Record.Report[i].Semester ,"",0,"LM",false,0,"")
	PosY += 5

	//Geade Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.SetXY(PosX,PosY)
	pdf.CellFormat(14,4, "PROGRAM" ,"",0,"LM",false,0,"")
	pdf.SetXY(PosX + 14,PosY)
	pdf.CellFormat(10,4, "CRS#" ,"",0,"LM",false,0,"")
	pdf.SetXY(PosX + 24,PosY)
	pdf.CellFormat(40,4, "COURSE TITLE" ,"",0,"LM",false,0,"")
	pdf.SetXY(PosX + 64,PosY)
	pdf.CellFormat(10,4, "UNITS" ,"",0,"LM",false,0,"")
	pdf.SetXY(PosX + 74,PosY)
	pdf.CellFormat(10,4, "GRADE" ,"",0,"LM",false,0,"")
	pdf.SetXY(PosX + 84,PosY)
	pdf.CellFormat(10,4, "POINTS" ,"",0,"LM",false,0,"")
	PosY += 5

	//Grade Rows
	for j := 0 ; j < len(Record.Report[i].Grades); j++ {
		pdf.SetFont("Arial","",6)
		pdf.SetXY(PosX,PosY)
		pdf.CellFormat(14,4, "ECE" ,"",0,"LM",false,0,"")

		pdf.SetXY(PosX + 14,PosY)
		TEXT := pdf.SplitLines([]byte(Record.Report[i].Grades[j].CourseCode),10)
		for _,lN := range TEXT {
			pdf.CellFormat(10, 3 ,string(lN),"",2 ,"LM",false,0,"")
		}

		pdf.SetXY(PosX + 24,PosY)
		TEXT = pdf.SplitLines([]byte(Record.Report[i].Grades[j].Title),40)
		L := 0
		if len(TEXT) <= 2 {
			for _,lN := range TEXT {
				pdf.CellFormat(40, 3 ,string(lN),"", 2 ,"LM",false,0,"")
			}
		}else{
			L = -2
			for _,lN := range TEXT {
				pdf.CellFormat(40, 3 ,string(lN),"", 2 ,"LM",false,0,"")
				L++
			}
		}
		

		pdf.SetXY(PosX + 64,PosY)
		pdf.CellFormat(10,4, strconv.Itoa(Record.Report[i].Grades[j].Units) ,"",0,"CM",false,0,"")

		pdf.SetXY(PosX+74,PosY)
		if Record.Report[i].Grades[j].Final_score != ""{
			pdf.CellFormat(10,4, Record.Report[i].Grades[j].Final_score ,"",0,"CM",false,0,"")
		}else{
			pdf.CellFormat(10,4, "TBA" ,"",0,"CM",false,0,"")
		}
		
		pdf.SetXY(PosX + 84,PosY)
		pdf.CellFormat(10,4, strconv.FormatFloat(Record.Report[i].Grades[j].Points,'f',1,32) ,"",0,"CM",false,0,"")

		PosY += 7 + (3*float64(L))
	}

	//Summary Header
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(24,3, "" ,"",0,"RB",false,0,"")
	pdf.CellFormat(40,3, "UNITS" ,"",0,"LB",false,0,"")
	pdf.CellFormat(10,3, "UNITS" ,"",0,"RB",false,0,"")
	pdf.CellFormat(10,3, "FINAL" ,"",0,"RB",false,0,"")
	pdf.CellFormat(10,3, "TOTAL" ,"",0,"RB",false,0,"")
	PosY += 4
	pdf.SetXY(PosX,PosY)
	pdf.CellFormat(24,3, "" ,"",0,"RB",false,0,"")
	pdf.CellFormat(40,3, "PASSED" ,"",0,"LT",false,0,"")
	pdf.CellFormat(10,3, "FACTORABLE" ,"",0,"RT",false,0,"")
	pdf.CellFormat(10,3, "QPA" ,"",0,"CT",false,0,"")
	pdf.CellFormat(10,3, "POINTS" ,"",0,"RT",false,0,"")
	PosY += 4

	//Sumary Rows

	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(24,3, "Semester" ,"",0,"LM",false,0,"")
	pdf.SetFont("Arial","",6)
	pdf.CellFormat(40,3, strconv.Itoa(int( Record.Report[i].Summary.Unitspassed)) ,"",0,"LM",false,0,"")
	pdf.CellFormat(10,3, strconv.Itoa(int(Record.Report[i].Summary.UnitsFact)),"",0,"CM",false,0,"")
	pdf.CellFormat(10,3, strconv.FormatFloat(Record.Report[i].Summary.FinalQpa,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(10,3, strconv.FormatFloat(Record.Report[i].Summary.TotalPoints,'f',1,32) ,"",0,"CM",false,0,"")
	PosY += 4
	pdf.SetXY(PosX,PosY)
	pdf.SetFont("Arial","B",6)
	pdf.CellFormat(24,3, "Cumulative" ,"",0,"LM",false,0,"")
	pdf.SetFont("Arial","",6)
	pdf.CellFormat(40,3, strconv.Itoa(int(Record.Report[i].Cumulative.Unitspassed)) ,"",0,"LM",false,0,"")
	pdf.CellFormat(10,3, strconv.Itoa(int(Record.Report[i].Cumulative.UnitsFact)) ,"",0,"CM",false,0,"")
	pdf.CellFormat(10,3, strconv.FormatFloat(Record.Report[i].Cumulative.FinalQpa,'f',2,32) ,"",0,"CM",false,0,"")
	pdf.CellFormat(10,3, strconv.FormatFloat(Record.Report[i].Cumulative.TotalPoints,'f',1,32) ,"",0,"CM",false,0,"")
	PosY += 5

	return pdf
}