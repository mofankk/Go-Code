package main

// 项目地址: https://github.com/signintech/gopdf

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("wts11", "font/wts11.ttf")

	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "U", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.Cell(&gopdf.Rect{H:600, W:1000}, "您好")
	if err != nil {
		log.Println("PDF编辑失败")
	}
	err = pdf.WritePdf("file/hello.pdf")
	if err != nil {
		log.Println("文件创建失败")
	}
}