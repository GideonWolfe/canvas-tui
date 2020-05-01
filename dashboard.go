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
  // ui "github.com/GideonWolfe/termui/v3"
  // "github.com/GideonWolfe/termui/v3/widgets"
)

func createCourseScorePlot(assignmentsMatrix [][]Assignment, courses *[]Course) *widgets.Plot {

  backup := widgets.NewPlot()
	backup.Title = "Not enough data"
  backup.Data = make([][]float64, 1)
	backup.Data =[][]float64{{1, 2, 3, 4, 5}}
	backup.AxesColor = ui.ColorWhite
	backup.LineColors[0] = ui.ColorCyan
  var placeholder []float64
  placeholder = append(placeholder, 5.0)

  p0 := widgets.NewPlot()
	p0.Title = "Score by Course"
  p0.LineColors[0] = ui.ColorCyan
  p0.HorizontalScale = 6
  var dataLabels []string

  count := 0
  courseDict := make(map[string][]float64)
  for _, crs := range *courses {
    if count <= len(assignmentsMatrix) {
      var scoreList []float64
      var assignments []Assignment = assignmentsMatrix[count]
      for _, assn := range assignments {
        if !assn.Submission.SubmittedAt.IsZero() {
          if assn.Submission.Score > 0 && assn.PointsPossible != 0{
            percentScored := float64(assn.Submission.Score/assn.PointsPossible)*100
            scoreList = append(scoreList, percentScored)
          }
        }
      }
      count++
      courseDict[crs.CourseCode] = scoreList
    }
  }


  var data [][]float64
  for course, scores := range courseDict {
    if len(scores) != 0 {
      dataLabels = append(dataLabels, course)
      data = append(data, scores)
    }
  }

  p0.Data = data
  p0.DataLabels = dataLabels
  if len(p0.Data) <= 1 {
    return backup
  }
  return p0
}

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
    if crs.EndAt.IsZero() && !crs.Term.EndAt.IsZero()  {
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

func createSummaryStackedBarchart(courses *[]Course) *widgets.StackedBarChart {
  var courseNames []string
  // var courseScores [][]float64
  // var barColors []ui.Color
  bc := widgets.NewStackedBarChart()
  bc.Data = make([][]float64, len(*courses))
  i := 0
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      currentScore := crs.Enrollments[0].ComputedCurrentScore
      finalScore := crs.Enrollments[0].ComputedFinalScore
      courseNames = append(courseNames, crs.CourseCode)
      // courseScores = append(courseScores, currentScore)
      bc.Data[i] = []float64{finalScore, currentScore}
      // if currentScore > 80 {
        // barColors = append(barColors, ui.ColorGreen)
      // } else if currentScore > 70 {
        // barColors = append(barColors, ui.ColorYellow)
      // } else if currentScore > 60 {
        // barColors = append(barColors, ui.ColorMagenta)
      // } else if currentScore > 50 {
        // barColors = append(barColors, ui.ColorRed)
      // }
    }
    i++
  }

  // bc.Data = courseScores
  bc.Labels = courseNames
	bc.Title = "Current Course Scores"
	bc.BarWidth = 15
  // bc.BarColors = barColors
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
  bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
  bc.BarGap = 0 

  return bc

}


func createDashboardGrid(courses *[]Course, assignmentsMatrix [][]Assignment) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
	// p0.Text = someVal
	p0.Border = true

  // render the logo in the top left
  cl := canvasLogo()

  // render the bar chart with current course grades
  bc := createSummaryBarchart(courses)
  // sbc := createSummayStackedBarchart(courses)

  todoTable := createTodoTableDash(courses)

  scorePlot := createCourseScorePlot(assignmentsMatrix, courses)
  
	dashboardGrid := ui.NewGrid()
  dashboardGrid.Title = "Dashboard"
	termWidth, termHeight := ui.TerminalDimensions()
	dashboardGrid.SetRect(0, 0, termWidth, termHeight)
  dashboardGrid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/3, cl),
			ui.NewCol(2.0/3, bc),
		),
    ui.NewRow(2.0/4, todoTable),
    ui.NewRow(1.0/4, scorePlot),
  )
  return dashboardGrid
}


