package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	path2 "path"
	"training.go/gencert/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error)  {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	background(pdf)

	header(pdf, &cert)
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0,50,cert.LabelPresented, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0,50,cert.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0,50,cert.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0,50,cert.LabelDate, "C")
	pdf.Ln(30)

	footer(pdf)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path2.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)

	return nil
}

func footer(pdf *gofpdf.Fpdf)  {
	opts := gofpdf.ImageOptions{
		ImageType:             "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	margin := 50.0
	imageSize := 30.0
	pdf.ImageOptions("img/stamp.png", pageWidth - margin, pageHeight - margin, imageSize, 0, false, opts, 0, "")
}

func background(pdf *gofpdf.Fpdf)  {
	opts := gofpdf.ImageOptions{
		ImageType:             "png",
	}

	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png", 0, 0, pageWidth, pageHeight, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert)  {
	opts := gofpdf.ImageOptions{
		ImageType:             "png",
	}
	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename, x+margin, 20, imageWidth, 0, false, opts, 0, "")
	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename, x-margin, 20, imageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}