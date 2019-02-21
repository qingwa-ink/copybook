package main

import (
	"copybook/pdf"
	"fmt"
)

func main() {
	fmt.Println("------------------Start-------------------")
	defer fmt.Println("-------------------End--------------------")

	// 每行字多一点，字小一点，练习钢笔字
	pdf.LineWords = 12
	pdf.MakePdfCn("hello.pdf", "自定义字帖，随便写什么字想要什么字，快来定制")

	// 每行字少一点，字大一点，练习毛笔字
	pdf.LineWords = 3
	pdf.MakePdfCn("hello2.pdf", "自定义字帖，随便写什么字想要什么字，快来定制")
}
