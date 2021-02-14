package main

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "github.com/spf13/viper"
  // "log"
	// "strconv"
  // "time"
	// "reflect"
	// "time"
	// "bytes"
  ui "github.com/gizak/termui"
  "github.com/gizak/termui/widgets"
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
)

// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func placeholder() *ui.Grid {

  // var assignments *[]Assignment = fetchAssignments(course.ID)

  // dummy placeholder widget
  p0 := widgets.NewParagraph()
  p0.Title = "Overview"
  p0.Text = "This is a placeholder dashboard"
	p0.Border = true

	assignmentGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	assignmentGrid.SetRect(0, 0, termWidth, termHeight)
  assignmentGrid.Title = "Course Something Grid"

  assignmentGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0, p0), // left column for pages
		),
  )



  return assignmentGrid
}

