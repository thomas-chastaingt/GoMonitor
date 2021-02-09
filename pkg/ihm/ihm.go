package ihm

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thomas-chastaingt/GoMonitor/pkg/cpu"
)

var run = true

func App() {
	//init the MHI
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize user interface %v", err)
	}
	defer ui.Close()

	//CPU
	g0 := widgets.NewGauge()
	g0.Title = "CPU"
	g0.SetRect(0, 0, 50, 3)
	g0.Percent = cpu.GetCPUUsage()
	g0.BarColor = ui.ColorWhite
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorWhite

	// object rendering
	ui.Render(g0)

	// event as keyboard and mouse
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
