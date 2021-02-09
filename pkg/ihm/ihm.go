package ihm

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thomas-chastaingt/GoMonitor/pkg/cpu"
)

var run = true

//App return the UI
func App() {
	//init the MHI
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize user interface %v", err)
	}
	defer ui.Close()

	//CPU utilisation
	g0 := widgets.NewGauge()
	g0.Title = "CPU"
	g0.SetRect(0, 0, 50, 3)
	g0.Percent = cpu.ReturnCPUUsage().Pourcent
	g0.BarColor = ui.ColorWhite
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorWhite

	// CPU Frequency
	// var freqBuffer = make([]float64, 40)
	// freqBuffer[len(freqBuffer)-1] = float64(status.CPUFreq.Freq) / 1000

	// cpuFreqPlot := widgets.NewPlot()
	// cpuFreqPlot.Title = " CPU frequency GHz "
	// cpuFreqPlot.Data = [][]float64{freqBuffer}
	// cpuFreqPlot.SetRect(0, 12, 50, 24)
	// cpuFreqPlot.AxesColor = ui.ColorWhite
	// cpuFreqPlot.LineColors[0] = ui.ColorCyan

	// object rendering function
	render := func() {
		g0.Percent = cpu.ReturnCPUUsage().Pourcent
		ui.Render(g0)
	}
	// time to refresh screen every second
	ticker := time.NewTicker(1010 * time.Millisecond).C
	// event as keyboard and mouse
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			//loop rendering
			render()
		}

	}
}
