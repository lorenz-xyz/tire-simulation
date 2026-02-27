package render

import (
	"fmt"
	"github.com/gogpu/gogpu"
	"github.com/gogpu/gogpu/gmath"
	"log/slog"
	"os"
)

func Render() {

	app := gogpu.NewApp(gogpu.DefaultConfig().
		WithGraphicsAPI(gogpu.GraphicsAPISoftware).
		WithTitle("Hello GoGPU").
		WithSize(800, 600))

	// Enable info-level logging
	gogpu.SetLogger(slog.Default())

	// Enable debug-level logging for full diagnostics
	gogpu.SetLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	app.OnDraw(func(dc *gogpu.Context) {
		dc.DrawTriangleColor(gmath.Green)
	})

	fmt.Println("run app")
	app.Run()
	fmt.Println("app exited")
}
