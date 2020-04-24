package main

import (
  // "encoding/json"
  // "io/ioutil"
	// "log"
	// "net/http"
  // "github.com/spf13/viper"
  "fmt"
  // "math"
  // "time"
  // "bytes"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func createTodoTableDash(courses *[]Course) *widgets.Table {

  var tableData [][]string
  header := []string{"Name", "Course", "Due At", "Points"}
  tableData = append(tableData, header)
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      assignments := fetchAssignments(crs.ID)
      for _, assn := range *assignments {
        if assn.Submission.SubmittedAt.IsZero() {
          var assignmentData []string
          assignmentData = append(assignmentData, assn.Name)
          assignmentData = append(assignmentData, crs.CourseCode)
          assignmentData = append(assignmentData, assn.DueAt.Local().Format("1/2 3:04 PM"))
          assignmentData = append(assignmentData, fmt.Sprint(assn.PointsPossible))
          tableData = append(tableData, assignmentData)
        }
      }
    }
  }

  todoTable := widgets.NewTable()
  todoTable.Title = "To Do:"
  todoTable.Rows = tableData
	todoTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	todoTable.RowSeparator = true
  todoTable.FillRow = true
  todoTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  if len(tableData) >= 10 {
    todoTable.BorderStyle = ui.NewStyle(ui.ColorRed)
  } else if len(tableData) >=7 { 
    todoTable.BorderStyle = ui.NewStyle(ui.ColorYellow)
  } else if len(tableData) >=4 { 
    todoTable.BorderStyle = ui.NewStyle(ui.ColorBlue)
  } else if len(tableData) >=2 { 
    todoTable.BorderStyle = ui.NewStyle(ui.ColorCyan)
  } else if len(tableData) >= 0 { 
    todoTable.BorderStyle = ui.NewStyle(ui.ColorGreen)
  }
  return todoTable
}



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
  bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
  bc.BarGap = 0 

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

  todoTable := createTodoTableDash(courses)
  
	dashboardGrid := ui.NewGrid()
  dashboardGrid.Title = "Dashboard"
	termWidth, termHeight := ui.TerminalDimensions()
	dashboardGrid.SetRect(0, 0, termWidth, termHeight)
  dashboardGrid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/3, cl),
			ui.NewCol(2.0/3, bc),
		),
		ui.NewRow(3.0/4,
			ui.NewCol(1.0, todoTable),
		),
  )
  return dashboardGrid
}


