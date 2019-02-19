package main

import (
	"copybook/pdf"
	"fmt"
)

func main() {
	fmt.Println("------------------Start-------------------")
	defer fmt.Println("-------------------End--------------------")

	pdf.MakePdf("hello.pdf", "自定义字帖，随便写什么字想要什么字，快来定制")
}
