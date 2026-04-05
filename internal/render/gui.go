package render

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	//	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

var pointSize = 10

func Render() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	//pallet := material.Palette{Bg: color.NRGBA{R: 100, G: 100, B: 100, A: 255}}
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:

			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			colorBG(gtx.Ops)
			drawContactPatch(gtx.Ops)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

func drawRedRect(ops *op.Ops) {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func colorBG(ops *op.Ops) {
	bgColor := color.NRGBA{R: 20, G: 20, B: 20, A: 255}
	paint.ColorOp{Color: bgColor}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func drawContactPatch(ops *op.Ops) {
	blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}

	pT := f32.Pt(50, 25)
	pR := f32.Pt(75, 75)
	pB := f32.Pt(50, 125)
	pL := f32.Pt(25, 75)

	pTR := f32.Pt(pR.X, pT.Y)
	pBR := f32.Pt(pR.X, pB.Y)
	pBL := f32.Pt(pL.X, pB.Y)
	pTL := f32.Pt(pL.X, pT.Y)

	var path clip.Path
	path.Begin(ops)
	path.Move(pT)
	path.QuadTo(pTR, pR)
	path.QuadTo(pBR, pB)
	path.QuadTo(pBL, pL)
	path.QuadTo(pTL, pT)

	var patch = func() {
		defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
		paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
		paint.PaintOp{}.Add(ops)
	}

	patch()

	drawPoint(ops, blue, pT.Round())
	drawPoint(ops, blue, pB.Round())
	drawPoint(ops, blue, pL.Round())
	drawPoint(ops, blue, pR.Round())

	drawPoint(ops, green, pTR.Round())
	drawPoint(ops, green, pBR.Round())
	drawPoint(ops, green, pBL.Round())
	drawPoint(ops, green, pTL.Round())
}

func drawPoint(ops *op.Ops, color color.NRGBA, pt image.Point) {
	hPtSize := pointSize / 2.0
	defer clip.Ellipse{Min: image.Pt(pt.X-hPtSize, pt.Y-hPtSize), Max: image.Pt(pt.X+hPtSize, pt.Y+hPtSize)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
