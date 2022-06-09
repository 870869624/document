package text

import (
	"log"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func Newtext() error {
	doc := document.New()
	defer doc.Close()

	para := doc.AddParagraph()
	para.SetStyle("Title")
	run := para.AddRun()
	run.AddBreak()
	para.SetAlignment(wml.ST_JcCenter)
	run.AddText("广告监测报告")

	para = doc.AddParagraph()
	para.SetAlignment(wml.ST_JcCenter)
	// para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("监测时间:2022年02月01日至2022年02月28日")

	para = doc.AddParagraph()
	run = para.AddRun()
	for i := 0; i < 14; i++ {
		run.AddBreak()
	}
	para = doc.AddParagraph()
	para.SetAlignment(wml.ST_JcCenter)
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("顶级广告监测中心")

	para = doc.AddParagraph()
	run = para.AddRun()

	para = doc.AddParagraph()
	run.Properties().SetEmboss(true)
	para.SetAlignment(wml.ST_JcCenter)
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("第一部分 广告监测情况综述")

	para = doc.AddParagraph()
	para.Properties().SetFirstLineIndent(0.4 * measurement.Inch)
	run = para.AddRun()
	run.AddText("2022年02月01日至2022年02月28日,顶级广告监测中心共监测各类媒体（电视、广播、报纸）发布的各类广告168078条次，其中电视广告84396条次；广播广告82962条次；报纸广告720条次。")

	img1, err := common.ImageFromFile("/Users/sharon/Desktop/图片1.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}
	imgref, err := doc.AddImage(img1)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}
	para = doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(imgref)
	if err != nil {
		log.Fatalf("unable to add anchored image:%s", err)
	}
	anchored.SetName("电视、广播、广告比例图")
	anchored.SetSize(5*measurement.Inch, 5*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)
	anchored.SetOffset(4*measurement.Centimeter, 3*measurement.Inch)

	// para.Properties().SetFirstLineIndent(0.4 * measurement.Inch)
	// run = para.AddRun()
	// run.AddBreak()
	// run.AddText("其中合法广告167934条次；违法广告144条次；总违法率为0.09%。")

	doc.SaveToFile("测试文件.docx")
	return nil
}
