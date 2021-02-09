package ihm

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var run = true

func App() {
	//init the MHI
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize user interface %v", err)
	}
	defer ui.Close()

	g0 := widgets.NewGauge()
	g0.Title = "CPU"
	g0.SetRect(30, 30, 10, 10)
	g0.Percent = 75
	g0.BarColor = ui.ColorWhite
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorWhite

	ui.Render(g0)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
