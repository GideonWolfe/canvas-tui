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



func createAssignmentTable(assignments *[]Assignment) *widgets.Table {

  var tableData [][]string
  header := []string{"Name", "Due at" ,"Score"}
  for _, assn := range *assignments {
    var assignmentData []string
    assignmentData = append(assignmentData, assn.Name)
    // percentScored := float64(assn.Submission.Score/assn.PointsPossible)*100
    scoreString := fmt.Sprint(assn.Submission.EnteredScore)+"/"+fmt.Sprint(assn.PointsPossible)
    dueAt := assn.DueAt.Local().Format("1/2 4:05 PM")
    assignmentData = append(assignmentData, dueAt, scoreString)
    tableData = append(tableData, assignmentData)
  }
  tableData = append(tableData, header)
  // reverse the list
  for i, j := 0, len(tableData)-1; i < j; i, j = i+1, j-1 {
    tableData[i], tableData[j] = tableData[j], tableData[i]
  }

  assignmentTable := widgets.NewTable()
  assignmentTable.Title = "Assignments:"
  assignmentTable.Rows = tableData
	assignmentTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	assignmentTable.RowSeparator = true
  assignmentTable.FillRow = true
  assignmentTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  return assignmentTable
}

// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createAssignmentGrid(course Course, assignments *[]Assignment) *ui.Grid {

  assignmentGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	assignmentGrid.SetRect(0, 0, termWidth, termHeight)
  assignmentGrid.Title = "Course Assignment Grid"
  assignmentGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0, createAssignmentTable(assignments)), 
		),
  )
  return assignmentGrid
}
