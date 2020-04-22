package main

import (
  // "encoding/json"
  // "io/ioutil"
	// "log"
	// "net/http"
  // "github.com/spf13/viper"
  "fmt"
  "math"
  // "time"
  // "bytes"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)



// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createCourseGrid(someVal string) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
  p0.Title = "Syllabus"
	p0.Text = someVal
	p0.Border = true

  // pie chart to eventually break down course points
  pc := widgets.NewPieChart()
	pc.Title = "Course Breakdown"
	pc.Data = []float64{.10, .10, .05, .20, .05, .13, .14, .25}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

  // list to select view of course
	l := widgets.NewList()
	l.Title = "Pages"
  l.Rows = []string{
		"[0] Assignmets",
		"[1] Quizzes",
		"[2] Grades",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false


	courseGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	courseGrid.SetRect(0, 0, termWidth, termHeight)
  courseGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0/6, l), // left column for pages
			ui.NewCol(5.0/6, // column for everything else
        ui.NewRow(1.0/4, //maybe some stats here?
          ui.NewCol(1.0/2, pc), // bar chart
          ui.NewCol(1.0/2, pc), // bar chart
        ),
        ui.NewRow(1.0/3, p0), // paragraph
      ),
		),
  )
  return courseGrid
}

