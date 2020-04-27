package main

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "github.com/spf13/viper"
	"fmt"
  // "log"
	"math"
	"strconv"
  // "time"
	// "reflect"
	// "time"
	// "bytes"
  strip "github.com/grokify/html-strip-tags-go"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func createAnnouncementWindow(course Course, announcements *[]Announcement) *widgets.Paragraph {
  // announcements := fetchAnnouncements(course.ID)
  p4 := widgets.NewParagraph()
	p4.Title = "Latest Announcement"
	p4.BorderStyle.Fg = ui.ColorBlue
  var recentAnnouncement Announcement
  i := 0
  for _, ann := range *announcements {
    if i == 0 {
      recentAnnouncement = ann
      break
    }
    i++
  }
  p4.Text = "[Title](fg:blue,mod:bold): "+recentAnnouncement.Title+"\n"+
            "[Date](fg:blue,mod:bold): "+recentAnnouncement.PostedAt.Local().Format("Jan 2, 2006")+"\n\n"+
            strip.StripTags(recentAnnouncement.Message)+"\n"
  return p4
}

func createSyllabusWindow(course Course) *widgets.Paragraph {
  p4 := widgets.NewParagraph()
	p4.Title = "Syllabus"
	p4.BorderStyle.Fg = ui.ColorYellow
  p4.Text = strip.StripTags(course.SyllabusBody)
  return p4
}


func createScorePlot(course Course, assignments *[]Assignment) *widgets.Plot {
  
  p0 := widgets.NewPlot()
	p0.Title = "Not enough data"
  p0.Data = make([][]float64, 1)
	p0.Data =[][]float64{{1, 2, 3, 4, 5}}
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen
  var placeholder []float64
  placeholder = append(placeholder, 5.0)



  p1 := widgets.NewPlot()
	p1.Title = "No graded assignments found"
	p1.Marker = widgets.MarkerDot
	p1.Data = [][]float64{[]float64{1, 2, 3, 4, 5}}
	p1.SetRect(50, 0, 75, 10)
	p1.DotMarkerRune = '+'
	p1.AxesColor = ui.ColorWhite
	p1.LineColors[0] = ui.ColorYellow
	p1.DrawDirection = widgets.DrawLeft


  
  var dataList [][]float64
  var plotData []float64
  for _, assn := range *assignments {
    if !assn.Submission.SubmittedAt.IsZero() { 
      if !assn.Submission.GradedAt.IsZero()  { // filter out nongraded assignments
        percent := assn.Submission.Score/float64(assn.PointsPossible)*100
        if assn.PointsPossible > 0  && percent != 0 {
          plotData = append(plotData, percent)
        }
      }
    }
  }
  dataList = append(dataList, plotData)

  // hack if no graded assignments
  if len(plotData) <= 1 {
    return p1
  }

  // if course.CourseCode == "CSCI 347" {
    // log.Panic(dataList)
  // }
  p3 := widgets.NewPlot()
	p3.Title = "Assignment Score(%) Over Time"
  p3.Data = dataList

  p3.AxesColor = ui.ColorWhite
  p3.LineColors[0] = ui.ColorCyan
  p3.Marker = widgets.MarkerBraille
  p3.PlotType = widgets.LineChart
  p3.MaxVal = 105
  p3.HorizontalScale = 3

  return p3
}

func createTodoTable(course Course, assignments *[]Assignment) *widgets.Table {

  var tableData [][]string
  header := []string{"Name", "Due At", "Points"}
  tableData = append(tableData, header)
  for _, assn := range *assignments {
    if assn.Submission.SubmittedAt.IsZero() {
      var assignmentData []string
      assignmentData = append(assignmentData, assn.Name)
      assignmentData = append(assignmentData, assn.DueAt.Local().Format("1/2 3:04 PM"))
      assignmentData = append(assignmentData, fmt.Sprint(assn.PointsPossible))
      tableData = append(tableData, assignmentData)
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

func createGradeSummaryTable(assignments *[]Assignment) *widgets.Table {

  var tableData [][]string
  header := []string{"Name",  "Score"}
  for _, assn := range *assignments {
    if !assn.Submission.SubmittedAt.IsZero() {
      if assn.Submission.Score > 0{
        var assignmentData []string
        assignmentData = append(assignmentData, assn.Name)
        // percentScored := float64(assn.Submission.Score/assn.PointsPossible)*100
        scoreString := fmt.Sprint(assn.Submission.EnteredScore)+"/"+fmt.Sprint(assn.PointsPossible)
        assignmentData = append(assignmentData, scoreString)
        tableData = append(tableData, assignmentData)
      }
    }
  }
  tableData = append(tableData, header)
  // reverse the list
  for i, j := 0, len(tableData)-1; i < j; i, j = i+1, j-1 {
    tableData[i], tableData[j] = tableData[j], tableData[i]
  }
  gradeTable := widgets.NewTable()
  gradeTable.Title = "Recent Scores:"
  gradeTable.Rows = tableData
	gradeTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	gradeTable.RowSeparator = true
  gradeTable.FillRow = true
  gradeTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  return gradeTable
}


func createAssignmentProgressBar(course Course, assignments *[]Assignment) *widgets.Gauge {

  var assignmentCount int64
  var completedAssignmentCount int64
  for _, assn := range *assignments {
    assignmentCount++
    if !assn.Submission.SubmittedAt.IsZero(){
      completedAssignmentCount++
    }
  }

  g1 := widgets.NewGauge()
  g1.Title = "Assignments: "+strconv.FormatInt(completedAssignmentCount, 10)+"/"+strconv.FormatInt(assignmentCount, 10)
  if completedAssignmentCount != 0 {
    currentPercent := int(((float64(completedAssignmentCount)/float64(assignmentCount))*100))
    g1.Percent = currentPercent
    if currentPercent > 80 {
      g1.BarColor = ui.ColorGreen
    } else if currentPercent > 70 {
      g1.BarColor = ui.ColorYellow
    } else if currentPercent > 60 {
      g1.BarColor = ui.ColorRed
    }
  } else {
    g1.Percent = 0
  }
	g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	g1.TitleStyle.Fg = ui.ColorMagenta
	g1.BorderStyle.Fg = ui.ColorWhite

  return g1
}


func createCoursePieChart(assignmentGroups *[]AssignmentGroup) *widgets.PieChart {
  var weights []float64
  var names []string
  pc := widgets.NewPieChart()
	pc.Title = "Course Breakdown"
	pc.AngleOffset = -.5 * math.Pi
  for _, ag := range *assignmentGroups {
    weights = append(weights, ag.GroupWeight)
    names = append(names, ag.Name)
  }
  if len(weights) == 0{
    weights = append(weights, 100)
  }
  pc.Data = weights
  pc.LabelFormatter = func(i int, v float64) string {
    return fmt.Sprintf("%s: %.f%%", names[i], v)
  }
  return pc
}




// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createCourseOverviewGrid(course Course, assignments *[]Assignment, announcements *[]Announcement, assignmentGroups *[]AssignmentGroup) *ui.Grid {


  var overviewText string = "Professor: "+course.Teachers[0].DisplayName+"\n" +
                            "Students: "+strconv.FormatInt(int64(course.TotalStudents), 10)+"\n" +
                            "Role: "+course.Enrollments[0].Type+"\n" + 
                            "Term: "+course.Term.Name+"\n" +
                            "Started At: "+course.StartAt.Local().Format("Jan 2, 2006")+"\n" + 
                            "Ends At: "+course.Term.EndAt.Local().Format("Jan 2, 2006")+"\n"
                            // "Calendar: "+course.Calendar.Ics+"\n" +
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
  p0.Title = "Overview"
  p0.Text = overviewText
	p0.Border = true

  assignmentProgressBar := createAssignmentProgressBar(course, assignments)

  // var assignmentGroups *[]AssignmentGroup = fetchAssignmentGroups(course.ID)

  pc := createCoursePieChart(assignmentGroups)

  sp := createScorePlot(course, assignments)

  announcementWindow := createAnnouncementWindow(course, announcements)

  syllabus := createSyllabusWindow(course)


  // list to select view of course
	l := widgets.NewList()
	l.Title = "Pages"
  l.Rows = []string{}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

  for _, tab := range course.Tabs {
    l.Rows = append(l.Rows, tab.Label)
  }

  todoTable := createTodoTable(course, assignments)
  gradeTable := createGradeSummaryTable(assignments)

	courseGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	courseGrid.SetRect(0, 0, termWidth, termHeight)
  courseGrid.Title = "Course Overview Grid"
  

  courseGrid.Set(
		ui.NewRow(1.0, 
      ui.NewRow(1.0/20, //maybe some stats here?
        ui.NewCol(1.0, assignmentProgressBar), // assignment completion progress
      ),
      ui.NewRow(7.0/20, //maybe some stats here?
        ui.NewCol(1.0/7, p0), // course overview
        ui.NewCol(3.0/7, todoTable), // assignment completion progress
        ui.NewCol(3.0/7, gradeTable), // assignment completion progress
      ),
      ui.NewRow(7.0/20,  // 
        ui.NewCol(2.0/4, announcementWindow),
        ui.NewCol(2.0/4, syllabus),
      ),
      ui.NewRow(5.0/20,  // 
        ui.NewCol(1.0/2, sp),
        ui.NewCol(1.0/2, pc),
      ),
    ),
  )
  return courseGrid
}

