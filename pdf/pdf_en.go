package pdf

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

// MakePdfEn : Generate copybook file
//
// 生成英文字帖文件
func MakePdfEn(filePath, content string) {

	pdf, err := getPdfEn()

	if err != nil {
		fmt.Printf("err : %s\n", err.Error())
		return
	}

	// TODO 绘制表格
	drawTableEn(pdf)

	// TODO 文本内容部分
	err = insertTextEn(pdf, content)
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

func getPdfEn() (pdf *gopdf.GoPdf, err error) {

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
	LineHeight = float64(fontSize) * 5 / 2

	err = pdf.SetFont("good", "", fontSize)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX(MarginWidth)
	pdf.SetY(MarginHeight)

	return
}

func drawTableEn(pdf *gopdf.GoPdf) (err error) {

	size := (PageHeight - MarginHeight) / TableDivider
	size = math.Floor(size)
	height := 0.0

	pdf.SetStrokeColor(ColorLineR, ColorLineG, ColorLineB)
	for i := 0; i < int(size); i++ {
		height = MarginHeight + float64(i)*LineHeight + float64(fontSize)*0.43
		pdf.Line(TableLeft, height, TableRight, height)
		height = height + float64(fontSize)*0.44
		pdf.Line(TableLeft, height, TableRight, height)
	}

	pdf.SetStrokeColor(ColorTableR, ColorTableG, ColorTableB)
	for i := 0; i < int(size+1); i++ {
		height = MarginHeight + float64(i)*LineHeight
		pdf.Line(TableLeft, height, TableRight, height)
		height = height + float64(fontSize)*1.3
		pdf.Line(TableLeft, height, TableRight, height)
	}
	tableBottom = height

	return nil
}

func insertTextEn(pdf *gopdf.GoPdf, text string) (err error) {

	pdf.SetTextColor(ColorTextR, ColorTextG, ColorTextB)

	lines := strings.Split(text, "\n")
	// AddLine(pdf, LineHeight*0.5)

	for _, line := range lines {
		sum := 0
		tmp := (PageWidth - MarginWidth*2) / float64(fontSize)
		count := int(tmp)
		datas := []rune(line)
		for sum < len(datas) {
			if sum+count > len(datas) {
				count = len(datas) - sum
			}
			w, _ := pdf.MeasureTextWidth(string(datas[sum : sum+count]))
			for {
				if w > PageWidth-MarginWidth*2 {
					count--
				} else if sum+count < len(datas) && w < PageWidth-MarginWidth*2-float64(fontSize) {
					count++
				} else {
					break
				}
				w, _ = pdf.MeasureTextWidth(string(datas[sum : sum+count]))
			}
			for sum+count < len(datas) && datas[sum+count-1] != ' ' {
				count--
			}
			pdf.SetX(MarginWidth)
			err = pdf.Cell(nil, string(datas[sum:sum+count]))
			if err != nil {
				return err
			}
			sum += count
			AddLine(pdf, LineHeight)
		}
	}

	return
}
