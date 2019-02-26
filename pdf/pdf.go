package pdf

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/signintech/gopdf"
)

const (
	PageWidth                             = 595.28
	PageHeight                            = 841.89
	MarginWidth                           = 48.0
	MarginHeight                          = 56.0
	ColorTableR, ColorTableG, ColorTableB = 0x01, 0x01, 0x01
	ColorLineR, ColorLineG, ColorLineB    = 0xcc, 0xcc, 0xcc
	ColorTextR, ColorTextG, ColorTextB    = 0xaa, 0xaa, 0xaa
)

var (
	// LineWords 每行的文字数量
	LineWords    = 12
	fontSize     = 32
	TableLeft    = 0.0
	TableRight   = 0.0
	TableTop     = 0.0
	tableBottom  = 0.0
	TableDivider = 0.0
	LineHeight   = 18.0
)

// MakePdfCn : Generate copybook file
//
// 生成中文字帖文件
func MakePdfCn(filePath, content string) {

	pdf, err := getPdf()

	if err != nil {
		fmt.Printf("err : %s\n", err.Error())
		return
	}

	// TODO 绘制表格
	drawTable(pdf)

	// TODO 文本内容部分
	err = insertText(pdf, content)
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
		Title:        "Copybook",
		Subject:      "Copybook",
		Producer:     "qingwa.ink",
		Author:       "qingwa.ink",
		Creator:      "qingwa.ink",
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

	pdf.SetStrokeColor(0x01, 0x01, 0x01)
	pdf.SetLineWidth(0.5)
	TableLeft = MarginWidth
	TableRight = PageWidth - MarginWidth
	TableDivider = (TableRight - TableLeft) / float64(LineWords)
	fontSize = int(TableDivider * 9 / 10)

	err = pdf.SetFont("good", "", fontSize)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX(MarginWidth)
	pdf.SetY(MarginHeight)

	return
}

func drawTable(pdf *gopdf.GoPdf) (err error) {

	size := (PageHeight - MarginHeight) / TableDivider
	size = math.Floor(size)
	height := 0.0

	pdf.SetStrokeColor(ColorLineR, ColorLineG, ColorLineB)
	for i := 0; i < int(size); i++ {
		height = MarginHeight + (float64(i)+0.5)*TableDivider
		pdf.Line(TableLeft, height, TableRight, height)
	}
	height += TableDivider * 0.5
	for i := 0; i < LineWords; i++ {
		width := TableLeft + (float64(i)+0.5)*TableDivider
		pdf.Line(width, MarginHeight, width, height)
	}

	pdf.SetStrokeColor(ColorTableR, ColorTableG, ColorTableB)
	for i := 0; i < int(size+1); i++ {
		height = MarginHeight + float64(i)*TableDivider
		pdf.Line(TableLeft, height, TableRight, height)
	}
	for i := 0; i < LineWords+1; i++ {
		width := TableLeft + float64(i)*TableDivider
		pdf.Line(width, MarginHeight, width, height)
	}
	tableBottom = height

	return nil
}

func insertText(pdf *gopdf.GoPdf, text string) (err error) {

	pdf.SetTextColor(ColorTextR, ColorTextG, ColorTextB)

	datas := []rune(text)
	height := 0.0
	for index, data := range datas {
		col := index % LineWords
		if height == 0 {
			height = MarginHeight + (TableDivider-float64(fontSize))/2
		} else {
			if col == 0 {
				height += TableDivider
			}
			if height >= tableBottom {
				pdf.AddPage()
				drawTable(pdf)
				height = MarginHeight + (TableDivider-float64(fontSize))/2
			}
		}
		pdf.SetY(height)
		pdf.SetX(float64(col)*TableDivider + TableLeft + (TableDivider-float64(fontSize))/2)
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
