package pdf

import (
	"fmt"
	"log"
	"time"

	"github.com/signintech/gopdf"
)

const (
	PageWidth                             = 595.28
	PageHeight                            = 841.89
	MarginWidth                           = 48.0
	MarginHeight                          = 56.0
	FontSize                              = 32.0
	LineHeight                            = 18.0
	LineWords                             = 12
	ColorTableR, ColorTableG, ColorTableB = 0x01, 0x01, 0x01
	ColorLineR, ColorLineG, ColorLineB    = 0xcc, 0xcc, 0xcc
	ColorTextR, ColorTextG, ColorTextB    = 0xaa, 0xaa, 0xaa
)

var (
	TableLeft    = 0.0
	TableRight   = 0.0
	TableTop     = 0.0
	TableBottom  = 0.0
	TableDivider = 0.0
)

// MakePdf : Generate copybook file
//
// 生成字帖文件
func MakePdf(filePath, content string) {

	pdf, err := getPdf()

	if err != nil {
		fmt.Printf("err : %s\n", err.Error())
		return
	}

	// TODO 绘制表格
	DrawTable(pdf)

	// TODO 借款表单部分
	err = InsertText(pdf, content)
	if err != nil {
		fmt.Printf("err : %s\n", err.Error())
		return
	}

	err = WritePdf(pdf, filePath)
	if err != nil {
		fmt.Printf("err : %s\n", err.Error())
		return
	}
}

func getPdf() (pdf *gopdf.GoPdf, err error) {

	pdf = &gopdf.GoPdf{}

	info := gopdf.PdfInfo{
		Title:        "Loan Agreement",
		Subject:      "Loan Agreement",
		Producer:     "Opera&MobiMagic",
		Author:       "Opera&MobiMagic",
		Creator:      "Opera&MobiMagic",
		CreationDate: time.Now(),
	}
	pdf.SetInfo(info)

	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: PageWidth, H: PageHeight}})
	pdf.AddPage()
	err = pdf.AddTTFFont("good", "xingzhebiji.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("good", "", FontSize)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetStrokeColor(0x01, 0x01, 0x01)
	pdf.SetLineWidth(0.5)
	TableLeft = MarginWidth
	TableRight = PageWidth - MarginWidth
	TableDivider = (TableRight - TableLeft) / LineWords

	pdf.SetX(MarginWidth)
	pdf.SetY(MarginHeight)

	return
}
func DrawTable(pdf *gopdf.GoPdf) (err error) {

	size := (PageHeight - MarginHeight) / TableDivider
	height := 0.0

	pdf.SetStrokeColor(ColorLineR, ColorLineG, ColorLineB)
	for i := 0; i < int(size-1); i++ {
		height = MarginHeight + (float64(i)+0.5)*TableDivider
		pdf.Line(TableLeft, height, TableRight, height)
	}
	height += TableDivider * 0.5
	for i := 0; i < LineWords; i++ {
		width := TableLeft + (float64(i)+0.5)*TableDivider
		pdf.Line(width, MarginHeight, width, height)
	}

	pdf.SetStrokeColor(ColorTableR, ColorTableG, ColorTableB)
	for i := 0; i < int(size); i++ {
		height = MarginHeight + float64(i)*TableDivider
		pdf.Line(TableLeft, height, TableRight, height)
	}
	for i := 0; i < LineWords+1; i++ {
		width := TableLeft + float64(i)*TableDivider
		pdf.Line(width, MarginHeight, width, height)
	}

	return nil
}

func InsertText(pdf *gopdf.GoPdf, text string) (err error) {

	pdf.SetTextColor(ColorTextR, ColorTextG, ColorTextB)

	datas := []rune(text)
	for index, data := range datas {
		row := index / LineWords
		col := index % LineWords
		pdf.SetX(float64(col)*TableDivider + TableLeft + (TableDivider-FontSize)/2)
		pdf.SetY(float64(row)*TableDivider + MarginHeight + (TableDivider-FontSize)/2)
		err = pdf.Cell(nil, string(data))
	}

	return
}

func WritePdf(pdf *gopdf.GoPdf, filePath string) (err error) {
	err = pdf.WritePdf(filePath)
	return
}

func AddLine(pdf *gopdf.GoPdf, y float64) {
	if pdf.GetY()+y > PageHeight-MarginHeight-LineHeight {
		pdf.AddPage()
		pdf.SetY(MarginHeight)
	} else {
		pdf.SetY(pdf.GetY() + y)
	}
}
