package generator

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

type PaperSize string

const (
	Paper_A5 PaperSize = "A5"
	Paper_A4 PaperSize = "A4"
)

type ReceiptFileGenerator struct {
	ctx       context.Context
	pdf       *gofpdf.Fpdf
	paperSize PaperSize
}

func NewReceiptFileGenerator(ctx context.Context, paperSize PaperSize) (*ReceiptFileGenerator, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fontFilePath := filepath.Join(currentDir, "assets", "font")
	pdf := gofpdf.New("P", "mm", string(paperSize), fontFilePath)

	return &ReceiptFileGenerator{
		ctx:       ctx,
		pdf:       pdf,
		paperSize: paperSize,
	}, nil
}

func (file *ReceiptFileGenerator) GenerateReceipt(leftHeader, rightHeader string, isFirstHeaderVisible bool, listModelData []ListModelData) (bytes.Buffer, error) {
	var buf bytes.Buffer

	paper := GetPaperA5()
	var lineHt = paper.LineHt

	pagew, pageh := file.pdf.GetPageSize()

	file.pdf.SetMargins(paper.MarginSetup.LMargin*3, paper.MarginSetup.TMargin*3, paper.MarginSetup.RMargin*3)
	lf, tp, _, _ := file.pdf.GetMargins()
	file.pdf.SetHeaderFuncMode(func() {
		file.pdf.SetMargins(paper.MarginSetup.LMargin*3, paper.MarginSetup.TMargin*3, paper.MarginSetup.RMargin*3)
		x, y := file.pdf.GetXY()
		r, g, b := file.pdf.GetFillColor()
		log.Printf("RGB: %d%d%d", r, g, b)
		file.pdf.SetFillColor(235, 235, 235)
		file.pdf.Rect(0, 0, pagew, pageh, "F")
		file.pdf.SetFillColor(r, g, b)
		file.pdf.SetXY(x, y)
		file.pdf.SetFillColor(255, 255, 255)
		////file.pdf.Rect(12.5, 12.5, 135.3, 197.3, "F")
		file.pdf.RoundedRect(paper.RectSetup.X, paper.RectSetup.Y, paper.RectSetup.W, paper.RectSetup.H, 3, "1234", "F")
		file.pdf.SetFillColor(r, g, b)
		file.pdf.SetXY(x, y)
		file.pdf.SetAlpha(0.15, "Normal")

		file.pdf.SetTextColor(235, 235, 235)

		for i := paper.TransformSetup.I.Min; i < paper.TransformSetup.I.Max; i++ {
			for j := paper.TransformSetup.J.Min; j < paper.TransformSetup.J.Max; j++ {
				file.pdf.TransformBegin()
				file.pdf.TransformRotate(paper.TransformSetup.Angle, paper.TransformSetup.X.A+(i*paper.TransformSetup.X.B), paper.TransformSetup.Y.A+(j*paper.TransformSetup.Y.B))
				file.pdf.TransformRotate(0, paper.TransformSetup.X.A+(i*paper.TransformSetup.X.B), paper.TransformSetup.Y.A+(j*paper.TransformSetup.Y.B))
				file.pdf.TransformEnd()
			}
		}

		file.pdf.SetAlpha(1, "Normal")
	}, true)

	file.pdf.AddPage()
	file.pdf.SetMargins(paper.MarginSetup.LMargin*3, paper.MarginSetup.TMargin*3, paper.MarginSetup.RMargin*3)

	file.pdf.SetY(tp + paper.MarginSetup.TMargin*3)
	file.pdf.SetDrawColor(224, 224, 224)
	file.pdf.SetTextColor(66, 66, 66)

	file.pdf.SetY(file.pdf.GetY() - paper.TransactionTextSetup.UpperSpace)
	y := paper.TransactionTextSetup.LowerSpace
	if len(listModelData) > 0 {
		file.pdf.SetY(tp + y)
	printReceipt:
		for i, datum := range listModelData {
			spaceLeft := pageh - (file.pdf.GetY() + lineHt)
			if spaceLeft < paper.BottomSetup.BottomLimit {
				file.pdf.SetY(pageh - paper.BottomSetup.BottomLimitMinus)
				file.pdf.SetFont("Arial", "B", paper.BottomSetup.FontSize)
				file.pdf.CellFormat(0, lineHt, "...", "0", 0, "CM", false, 0, "")
				break
			}
			//log.Printf("file.pdf.GetX: %f", file.pdf.GetX())
			//log.Printf("file.pdf.GetY: %f", file.pdf.GetY())
			//xHeader, y = file.pdf.GetXY()
			if len(datum.ModelData) > 0 {
				if !isFirstHeaderVisible {
					if strings.Trim(datum.HeaderData, " ") == "" {
						datum.HeaderData = "-"
					}
					for _, modelDatum := range datum.ModelData {
						spaceLeft = pageh - (file.pdf.GetY() + lineHt)
						if spaceLeft < paper.BottomSetup.BottomLimit {
							file.pdf.SetY(pageh - paper.BottomSetup.BottomLimitMinus)
							file.pdf.SetFont("Arial", "B", paper.BottomSetup.FontSize)
							file.pdf.CellFormat(0, lineHt, "...", "0", 0, "CM", false, 0, "")
							break printReceipt
						}
						if modelDatum.IsTotalPayment {
							file.pdf.Ln(paper.ValueCellSetup.Ln1)
							file.pdf.SetFont("Arial", "B", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.CellFormat(paper.ValueCellSetup.W1, paper.ValueCellSetup.H1, modelDatum.Key, "0", 0, "LM", false, 0, "")
							file.pdf.CellFormat(paper.ValueCellSetup.W2, paper.ValueCellSetup.H2, "", "0", 0, "CM", false, 0, "")
							file.pdf.SetFont("Arial", "B", paper.TotalPaymentFont.ValueFontSize)
							file.pdf.MultiCell(paper.ValueCellSetup.WMultiCell, paper.ValueCellSetup.HMultiCell, modelDatum.Value, "0", "RM", false)
							file.pdf.SetFont("Arial", "", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.Ln(paper.ValueCellSetup.Ln2)
						} else {
							file.pdf.SetFont("Arial", "", paper.ValueFont.HeaderFontSize)
							file.pdf.CellFormat(paper.ValueCellSetup.W1, paper.ValueCellSetup.H1, modelDatum.Key, "0", 0, "LM", false, 0, "")
							file.pdf.CellFormat(paper.ValueCellSetup.W2, paper.ValueCellSetup.H2, "", "0", 0, "CM", false, 0, "")
							file.pdf.SetFont("Arial", "B", paper.ValueFont.ValueFontSize)
							file.pdf.MultiCell(paper.ValueCellSetup.WMultiCell, paper.ValueCellSetup.HMultiCell, modelDatum.Value, "0", "RM", false)
							file.pdf.SetFont("Arial", "", paper.ValueFont.HeaderFontSize)
							file.pdf.Ln(paper.ValueCellSetup.Ln2)
						}
					}
					if i < len(listModelData)-1 {
						file.pdf.SetY(file.pdf.GetY() + paper.HeaderSetup.Space1)
						file.pdf.SetFillColor(128, 128, 128)
						file.pdf.SetTextColor(255, 255, 255)
						file.pdf.SetFont("Arial", "B", paper.HeaderSetup.FontSize)
						file.pdf.CellFormat(paper.HeaderSetup.W, paper.HeaderSetup.H, "",
							"", 0, "C", true, 0, "")
						file.pdf.Text(paper.HeaderSetup.X, file.pdf.GetY()+paper.HeaderSetup.Y, listModelData[i+1].HeaderData)
						file.pdf.SetFillColor(255, 255, 255)
						file.pdf.SetX(lf)
					}
					file.pdf.SetTextColor(66, 66, 66)
					y = file.pdf.GetY()
					file.pdf.SetY(tp + y - paper.HeaderSetup.Space2)
				} else {
					if strings.Trim(datum.HeaderData, " ") == "" {
						datum.HeaderData = "-"
					}
					file.pdf.SetY(file.pdf.GetY() + paper.HeaderSetup.Space1)
					file.pdf.SetFillColor(128, 128, 128)
					file.pdf.SetTextColor(255, 255, 255)
					file.pdf.SetFont("Arial", "B", paper.HeaderSetup.FontSize)
					file.pdf.CellFormat(paper.HeaderSetup.W, paper.HeaderSetup.H, "",
						"", 0, "C", true, 0, "")
					file.pdf.Text(paper.HeaderSetup.X, file.pdf.GetY()+paper.HeaderSetup.Y, datum.HeaderData)
					file.pdf.SetFillColor(255, 255, 255)
					file.pdf.SetX(lf)
					file.pdf.SetTextColor(66, 66, 66)

					y = file.pdf.GetY()
					file.pdf.SetY(tp + y)

					for _, modelDatum := range datum.ModelData {
						if modelDatum.IsTotalPayment {
							file.pdf.Ln(paper.ValueCellSetup.Ln1)
							file.pdf.SetFont("Arial", "B", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.CellFormat(paper.ValueCellSetup.W1, paper.ValueCellSetup.H1, modelDatum.Key, "0", 0, "LM", false, 0, "")
							file.pdf.CellFormat(paper.ValueCellSetup.W2, paper.ValueCellSetup.H2, "", "0", 0, "CM", false, 0, "")
							file.pdf.SetFont("Arial", "B", paper.TotalPaymentFont.ValueFontSize)
							file.pdf.MultiCell(paper.ValueCellSetup.WMultiCell, paper.ValueCellSetup.HMultiCell, modelDatum.Value, "0", "RM", false)
							file.pdf.SetFont("Arial", "", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.Ln(paper.ValueCellSetup.Ln2)
						} else {
							file.pdf.SetFont("Arial", "", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.CellFormat(paper.ValueCellSetup.W1, paper.ValueCellSetup.H1, modelDatum.Key, "0", 0, "LM", false, 0, "")
							file.pdf.CellFormat(paper.ValueCellSetup.W2, paper.ValueCellSetup.H2, "", "0", 0, "CM", false, 0, "")
							file.pdf.SetFont("Arial", "B", paper.TotalPaymentFont.ValueFontSize)
							file.pdf.MultiCell(paper.ValueCellSetup.WMultiCell, paper.ValueCellSetup.HMultiCell, modelDatum.Value, "0", "RM", false)
							file.pdf.SetFont("Arial", "", paper.TotalPaymentFont.HeaderFontSize)
							file.pdf.Ln(paper.ValueCellSetup.Ln2)
						}
					}
				}
			}
		}
	}

	err := file.pdf.Output(&buf)
	if err != nil {
		return buf, err
	}

	return buf, nil

}
