package main

import (
  // "encoding/json"
  // "io/ioutil"
	// "log"
	// "net/http"
  // "github.com/spf13/viper"
  // "fmt"
  // "math"
  // "time"
  // "bytes"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)




func createDashboardGrid(courses *[]Course) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
	// p0.Text = someVal
	p0.Border = true

  cl := canvasLogo()

  bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
  
	dashboardGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	dashboardGrid.SetRect(0, 0, termWidth, termHeight)
  dashboardGrid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/3, cl),
			ui.NewCol(2.0/3, bc),
		),
		ui.NewRow(3.0/4,
			ui.NewCol(1.0, p0),
		),
  )
  return dashboardGrid
}


