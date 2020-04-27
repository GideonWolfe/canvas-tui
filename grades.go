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
  "fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)


func createGradeTable(assignments *[]Assignment) *widgets.Table {

  gradeTable := widgets.NewTable()
  var tableData [][]string
  header := []string{"Name",  "Score", "%"}
  i := 1
  for _, assn := range *assignments {
    if !assn.Submission.SubmittedAt.IsZero() {
      if assn.Submission.Score > 0{
        var assignmentData []string
        assignmentData = append(assignmentData, assn.Name)
        percentScored := float64(assn.Submission.Score/assn.PointsPossible)*100
        scoreString := fmt.Sprint(assn.Submission.EnteredScore)+"/"+fmt.Sprint(assn.PointsPossible)
        assignmentData = append(assignmentData, scoreString)
        assignmentData = append(assignmentData, fmt.Sprintf("%.f%%", percentScored))
        tableData = append(tableData, assignmentData)
        if percentScored > 90 {
          gradeTable.RowStyles[i-1] = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
        } else if percentScored > 80 {
          gradeTable.RowStyles[i-1] = ui.NewStyle(ui.ColorBlue, ui.ColorClear, ui.ModifierBold)
        } else if percentScored > 70 {
          gradeTable.RowStyles[i-1] = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
        } else if percentScored <= 60 {
          // log.Panic(assn.Name)
          gradeTable.RowStyles[i-1] = ui.NewStyle(ui.ColorRed, ui.ColorClear, ui.ModifierBold)
        }
        i++
      }
    }
  }
  tableData = append(tableData, header)
  // reverse the list
  for i, j := 0, len(tableData)-1; i < j; i, j = i+1, j-1 {
    tableData[i], tableData[j] = tableData[j], tableData[i]
    gradeTable.RowStyles[i], gradeTable.RowStyles[j] = gradeTable.RowStyles[j], gradeTable.RowStyles[i]
  }

  gradeTable.Title = "Scores:"
  gradeTable.Rows = tableData
	gradeTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	gradeTable.RowSeparator = true
  gradeTable.FillRow = true
  gradeTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  return gradeTable
}


func createGradeGrid(course Course) *ui.Grid {

  var assignments *[]Assignment = fetchAssignments(course.ID)
  gradeGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	gradeGrid.SetRect(0, 0, termWidth, termHeight)
  gradeGrid.Title = "Course Overview Grid"
  gradeGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0, createGradeTable(assignments)), 
		),
  )

  return gradeGrid
}
