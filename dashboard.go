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

func createSummaryBarchart(courses *[]Course) *widgets.BarChart {
var courseNames []string
  var courseScores []float64
  var barColors []ui.Color
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      currentScore := crs.Enrollments[0].ComputedCurrentScore
      courseNames = append(courseNames, crs.CourseCode)
      courseScores = append(courseScores, currentScore)
      if currentScore > 80 {
        barColors = append(barColors, ui.ColorGreen)
      } else if currentScore > 70 {
        barColors = append(barColors, ui.ColorYellow)
      } else if currentScore > 60 {
        barColors = append(barColors, ui.ColorMagenta)
      } else if currentScore > 50 {
        barColors = append(barColors, ui.ColorRed)
      }
    }
  }

  bc := widgets.NewBarChart()
  bc.Data = courseScores
  bc.Labels = courseNames
	bc.Title = "Current Course Scores"
	bc.BarWidth = 15
  bc.BarColors = barColors
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	// bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
  bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

  return bc

}


func createDashboardGrid(courses *[]Course) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
	// p0.Text = someVal
	p0.Border = true

  // render the logo in the top left
  cl := canvasLogo()

  // render the bar chart with current course grades
  bc := createSummaryBarchart(courses)
  
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


