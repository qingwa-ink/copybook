package main

import (
	"copybook/pdf"
	"fmt"
)

var enDemo = `
Previous work experience is precious.
The manager assigned him some tasks, but he resigned.
Would request you to add the features of alternate number in application, so that if customer is not picking the phone , we can try alternate number though.`

func main() {
	fmt.Println("------------------Start-------------------")
	defer fmt.Println("-------------------End--------------------")

	// 每行字多一点，字小一点，练习钢笔字
	pdf.LineWords = 12
	pdf.MakePdfCn("hello.pdf", "自定义字帖，随便写什么字想要什么字，快来定制")

	// 每行字少一点，字大一点，练习毛笔字
	pdf.LineWords = 3
	pdf.MakePdfCn("hello2.pdf", "自定义字帖，随便写什么字想要什么字，快来定制")

	// 生成英文字帖，1个汉字位置表示2个英文字符
	pdf.LineWords = 20
	pdf.MakePdfEn("hello3.pdf", enDemo)
}
