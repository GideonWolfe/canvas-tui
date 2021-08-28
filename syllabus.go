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
  strip "github.com/grokify/html-strip-tags-go"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
)

// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createSyllabusGrid(course Course) *ui.Grid {

  p0 := widgets.NewParagraph()
  p0.Title = "Syllabus"
  p0.Text = strip.StripTags(course.SyllabusBody)
	p0.Border = true

	syllabusGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	syllabusGrid.SetRect(0, 0, termWidth, termHeight)
  syllabusGrid.Title = "Course Syllabus Grid"

  syllabusGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0, p0), // left column for pages
		),
  )

  return syllabusGrid
}

