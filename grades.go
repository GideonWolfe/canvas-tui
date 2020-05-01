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
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
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

func createAGBreakdown(assignments *[]Assignment, assignmentGroups *[]AssignmentGroup) *widgets.Plot {


  backup := widgets.NewPlot()
	backup.Title = "Not enough data"
  backup.Data = make([][]float64, 1)
	backup.Data =[][]float64{{1, 2, 3, 4, 5}}
	backup.AxesColor = ui.ColorWhite
	backup.LineColors[0] = ui.ColorGreen
  var placeholder []float64
  placeholder = append(placeholder, 5.0)


  p0 := widgets.NewPlot()
	p0.Title = "Score by Assignment Group"
  p0.LineColors[0] = ui.ColorCyan
  p0.HorizontalScale = 6
  var dataLabels []int

  // create a dictionary of assignment groups and their IDs
  count := 0
  aGDict := make(map[int]string)
  for _, ag := range *assignmentGroups {
    if len(ag.Assignments) > 0 {
      count++
      aGDict[ag.ID] = ag.Name
    }
  }

  p0.Data = make([][]float64, count)
  assignmentDict := make(map[int][]float64)

  for _, assn := range *assignments {
    if !assn.Submission.SubmittedAt.IsZero() {
      if assn.Submission.Score > 0 && assn.PointsPossible != 0{
        percentScored := float64(assn.Submission.Score/assn.PointsPossible)*100
        assignmentDict[assn.AssignmentGroupID] = append(assignmentDict[assn.AssignmentGroupID], percentScored)
      }
    }
  }

  i := 0
  for group, data := range assignmentDict {
    dataLabels = append(dataLabels, group)
    p0.Data = append(p0.Data, data)
    i++
  }

  // log.Panic(p0.Data, assignmentDict)
  if len(p0.Data) <= 1 {
    return backup
  }

  return p0
}

func createGradeGrid(course Course, assignments *[]Assignment, assignmentGroups *[]AssignmentGroup) *ui.Grid {

  // var assignments *[]Assignment = fetchAssignments(course.ID)
  gradeGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	gradeGrid.SetRect(0, 0, termWidth, termHeight)
  gradeGrid.Title = "Course Grade Grid"
  gradeGrid.Set(
		ui.NewRow(1.0, 
			ui.NewRow(2.0/3, createGradeTable(assignments)), 
			ui.NewRow(1.0/3, createAGBreakdown(assignments, assignmentGroups)), 
		),
  )

  return gradeGrid
}
