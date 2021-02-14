package main

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "github.com/spf13/viper"
	// "fmt"
  // "log"
	// "math"
  // "strconv"
  // "time"
	// "reflect"
	// "time"
	// "bytes"
  // strip "github.com/grokify/html-strip-tags-go"
  ui "github.com/gizak/termui"
  "github.com/gizak/termui/widgets"
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
)


// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createCourseGrid(course Course, assignments *[]Assignment, announcements *[]Announcement, assignmentGroups *[]AssignmentGroup) *ui.Grid {

  // var assignments *[]Assignment = fetchAssignments(course.ID)

  
  // list to select view of course
	l := widgets.NewList()
	l.Title = "Pages"
  l.Rows = []string{}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

  for _, tab := range course.Tabs {
    l.Rows = append(l.Rows, tab.Label)
  }


	courseGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	courseGrid.SetRect(0, 0, termWidth, termHeight)
  courseGrid.Title = "Course Master Grid"
  
  courseGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0/6, l), // left column for pages
			ui.NewCol(5.0/6,createCourseOverviewGrid(course, assignments, announcements, assignmentGroups)), // column for everything else
    ),
  )

  return courseGrid
}

