package main

import(
	"github.com/jung-kurt/gofpdf"
)

func header(pdf *gofpdf.Fpdf) *gofpdf.Fpdf{
	pdf.SetFont("Arial","B",7)
	pdf.Cell(200,3,"Officail Academic Record as of 20 Jun 2022")
	pdf.Ln(4)
	pdf.Cell(100,3 ,"Student Name: Omega Hub Student")
	pdf.Cell(100,3,"Student ID: omega.student@cmkl.ac.th")
	pdf.Line(10,18,200,18)
	pdf.Ln(5)
	pdf.Cell(20,3,"University: ")
	pdf.SetFont("Arial","",7)
	pdf.Cell(150,3,"CMKL University")
	pdf.Ln(4)
	pdf.SetFont("Arial","B",7)
	pdf.Cell(20,3,"Degree: ")
	pdf.SetFont("Arial","",7)
	pdf.Cell(150,3,"Master of Science")
	pdf.Ln(4)
	pdf.SetFont("Arial","B",7)
	pdf.Cell(20,3,"Program: ")
	pdf.SetFont("Arial","",7)
	pdf.Cell(150,3,"M.S. in Electrical and Computer Engeering")
	pdf.Line(10,31,200,31)
	pdf.Ln(5)
	pdf.SetFont("Arial","B",7)
	pdf.Cell(200,4,"Carnegie Mellon University - CMKL | THAILAND")
	pdf.Line(10,37,200,37)
	pdf.Ln(5)

	return pdf
}