package render

import (
	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

func Test2D() {
	// Create drawing context
	dc := gg.NewContext(512, 512)
	defer dc.Close()

	dc.ClearWithColor(gg.White)

	// Draw shapes
	dc.SetHexColor("#3498db")
	dc.DrawCircle(256, 256, 100)
	dc.Fill()

	// Render text
	source, err := text.NewFontSourceFromFile("assets/fonts/JetBrainsMonoNerdFont-Regular.ttf")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	dc.SetFont(source.Face(32))
	dc.SetColor(gg.Black.Color())
	dc.DrawString("Hello, GoGPU!", 180, 260)

	dc.SavePNG("output.png")
}
