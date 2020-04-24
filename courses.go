package main

import (
	// "encoding/json"
	// "io/ioutil"
  // "log"
	// "net/http"
	// "github.com/spf13/viper"
	// "fmt"
	"math"
	"strconv"
	// "time"

	// "reflect"
	// "time"
	// "bytes"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)


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
  // pie chart to eventually break down course points
  pc := widgets.NewPieChart()
	pc.Title = "Course Breakdown"
  // pc.Data = []float64{.5, .3}
	pc.AngleOffset = -.5 * math.Pi
  // pc.LabelFormatter = func(i int, v float64) string {
    // return fmt.Sprintf("%.02f", v)
  // }

  for _, ag := range *assignmentGroups {
    // f := strconv.FormatFloat(ag.GroupWeight, 'E', -1, 64)
    weights = append(weights, float64((float64(ag.GroupWeight)/float64(100.0))))
    // if ag.Name == "Take Home Midterm"{
      // log.Panic(ag)
    // }
    // log.Println(float64(ag.GroupWeight)/float64(100.0))
    // log.Panicf("%T", ag.GroupWeight)
  }
  // log.Panic(assignmentGroups)
  pc.Data = weights

  return pc

}




// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func createCourseGrid(course Course) *ui.Grid {

  var assignments *[]Assignment = fetchAssignments(course.ID)

  var overviewText string = "Professor: "+course.Teachers[0].DisplayName+"\n" +
                            "Students: "+strconv.FormatInt(int64(course.TotalStudents), 10)+"\n" +
                            "Role: "+course.Enrollments[0].Type+"\n" + 
                            "Started At: "+course.StartAt.Format("Jan 2, 2006")+"\n"
                            // "Calendar: "+course.Calendar.Ics+"\n" +
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
  p0.Title = "Overview"
  p0.Text = overviewText
	p0.Border = true

  assignmentProgressBar := createAssignmentProgressBar(course, assignments)

  var assignmentGroups *[]AssignmentGroup = fetchAssignmentGroups(course.ID)

  pc := createCoursePieChart(assignmentGroups)

  // list to select view of course
	l := widgets.NewList()
	l.Title = "Pages"
  l.Rows = []string{}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

  for _, tab := range course.Tabs {
    l.Rows = append(l.Rows, tab.Label)
  }

  // for _, ag := range *assignmentGroups {
    // f := strconv.FormatFloat(ag.GroupWeight, 'E', -1, 64)
    // overviewText = overviewText+f+"\n"
    // pc.Data = append(pc.Data, ag.GroupWeight)
  // }


	courseGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	courseGrid.SetRect(0, 0, termWidth, termHeight)
  courseGrid.Set(
		ui.NewRow(1.0, 
			ui.NewCol(1.0/6, l), // left column for pages
			ui.NewCol(5.0/6, // column for everything else
        ui.NewRow(1.0/4, //maybe some stats here?
          ui.NewCol(1.0/2, p0), // course overview
          ui.NewCol(1.0/2, assignmentProgressBar), // assignment completion progress
        ),
        ui.NewRow(1.0/3, pc), // 
      ),
		),
  )
  return courseGrid
}

