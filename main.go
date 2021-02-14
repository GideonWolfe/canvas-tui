package main

import (
	"log"
  "fmt"
  "time"
  "github.com/spf13/viper"
  ui "github.com/gizak/termui"
  "github.com/gizak/termui/widgets"
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
  "runtime"
  "os/exec"
  // "strconv"
  // "strings"
)

// Reads the config file
func readConfig() {
  viper.SetConfigName("config") // name of config file (without extension)
  viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
  viper.AddConfigPath(".")               // optionally look for config in the working directory
  viper.AddConfigPath("$HOME/.config/canvas-tui/")              
  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
}

// called to generate navigation tabs for courses
func createMainTabPane(courses *[]Course) *widgets.TabPane {
  var titles []string
  titles = append(titles, "Dashboard")
  for _, crs := range *courses {
    if crs.EndAt.IsZero() && !crs.Term.EndAt.IsZero() { // second check to filter labs with no terms :(
      titles = append(titles, crs.CourseCode)
    }
  }
  tabpane := widgets.NewTabPane(titles...)
	tabpane.Border = true
  return tabpane
}


func  chooseTab(courseMasterGrids []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid, contentGrid *ui.Grid) {
  // Substitute the current grid for what the user has selected

  // If we click the dashboard 
  if tabpane.ActiveTabIndex == 0 {
    masterGrid.Items[1].Entry = contentGrid
  } else { // for other course pages
    contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
    masterGrid.Items[1].Entry = contentGrid
  }
  
  ui.Render(masterGrid)
}

func  handleSpace(courseMasterGrids []ui.Grid, courseOverviewGrids []ui.Grid, courseGradeGrids []ui.Grid, courseAnnouncementGrids []ui.Grid, courseSyllabusGrids []ui.Grid, courseAssignmentGrids []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid, contentGrid *ui.Grid) {
  // Substitute the current grid for what the user has selected

  // If we click the dashboard 
  if tabpane.ActiveTabIndex == 0 {
    masterGrid.Items[1].Entry = contentGrid

  } else { // for other course pages
    // contentGrid = &courseGradeGrids[tabpane.ActiveTabIndex-1]
    // get the currently selected item
    contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
    item := contentGrid.Items[0].Entry.(*widgets.List).SelectedRow
    itemStr := contentGrid.Items[0].Entry.(*widgets.List).Rows[item]
  
    if itemStr == "Home" {
      contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
      contentGrid.Items[1].Entry = &courseOverviewGrids[tabpane.ActiveTabIndex-1] 
      masterGrid.Items[1].Entry = contentGrid
    } else if itemStr == "Grades" {
      contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
      contentGrid.Items[1].Entry = &courseGradeGrids[tabpane.ActiveTabIndex-1]
      masterGrid.Items[1].Entry = contentGrid
    } else if itemStr == "Announcements" {
      contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
      contentGrid.Items[1].Entry = &courseAnnouncementGrids[tabpane.ActiveTabIndex-1]
      masterGrid.Items[1].Entry = contentGrid
    } else if itemStr == "Syllabus" {
      contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
      contentGrid.Items[1].Entry = &courseSyllabusGrids[tabpane.ActiveTabIndex-1]
      masterGrid.Items[1].Entry = contentGrid
    } else if itemStr == "Assignments" {
      contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
      contentGrid.Items[1].Entry = &courseAssignmentGrids[tabpane.ActiveTabIndex-1]
      masterGrid.Items[1].Entry = contentGrid
    } else {
      contentGrid.Items[1].Entry = placeholder()
      masterGrid.Items[1].Entry = contentGrid
    } 
  }
  
  ui.Render(masterGrid)
  // log.Panic(contentGrid.Title)
}

// handles opeining a selection in the browser
func  handleOpen(courseMasterGrids []ui.Grid, courseOverviewGrids []ui.Grid, courseGradeGrids []ui.Grid, courseAnnouncementGrids []ui.Grid, courseSyllabusGrids []ui.Grid, courseAssignmentGrids []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid, contentGrid *ui.Grid, courses []Course) {
  // if we are on the dashboard, open canvas home
  var url string
  if tabpane.ActiveTabIndex == 0 {
    url = viper.Get("canvasdomain").(string)
  } else {
    contentGrid = &courseMasterGrids[tabpane.ActiveTabIndex-1]
    url = courses[tabpane.ActiveTabIndex-1].Tabs[contentGrid.Items[0].Entry.(*widgets.List).SelectedRow].FullURL
  }

  // actually open the URL
  var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func  menuScroll(coursePages []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid, contentGrid *ui.Grid, direction string) {
  // Substitute the current grid for what the user has selected
  // Don't try to scroll on the dashboard
  if tabpane.ActiveTabIndex != 0 {
    contentGrid = &coursePages[tabpane.ActiveTabIndex-1]
    l := contentGrid.Items[0].Entry.(*widgets.List)
    if direction == "down" {
      l.ScrollDown()
    } else if direction == "up"{
      l.ScrollUp()
    }
  }
  ui.Render(masterGrid)
}

// called if master grid needs to be updated
func updateMasterGrid(masterGrid *ui.Grid, tabpane *widgets.TabPane, contentGrid *ui.Grid) *ui.Grid {
  ui.Clear()
  // defining master grid layout
  masterGrid.Set(
    ui.NewRow(1.0/20,
      ui.NewCol(1.0, tabpane),
    ),
    ui.NewRow(19.0/20,
      ui.NewCol(1.0/1, contentGrid),
    ),
  )
  // ui.Render(masterGrid)
  return masterGrid
}

func main() {
  // Initialize temui
  if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

  // get the main config
  readConfig()

  var courses *[]Course = fetchCourses()
  var activeCourses []Course // variable needed to accurately choose/open courses from menu
  for _, crs := range *courses {
    if crs.EndAt.IsZero() && !crs.Term.EndAt.IsZero(){
      activeCourses = append(activeCourses, crs)
    }
  }

  
  // declare master grid and set terminal dimensions
	masterGrid := ui.NewGrid()
  masterGrid.Title = "Master Grid"
	termWidth, termHeight := ui.TerminalDimensions()
	masterGrid.SetRect(0, 0, termWidth, termHeight)

  // one list of assignments per course
  var assignmentsMatrix [][]Assignment

  // one list of announcements per course
  var announcementMatrix [][]Announcement

  // one list of announcements per course
  var assignmentGroupMatrix [][]AssignmentGroup

  // first fetch all the assignments to reduce redundant API calls
  for _, crs := range activeCourses {
  // for _, crs := range *courses {
    if crs.EndAt.IsZero() && !crs.Term.EndAt.IsZero() {
      assignmentsMatrix = append(assignmentsMatrix, *fetchAssignments(crs.ID))
      announcementMatrix = append(announcementMatrix, *fetchAnnouncements(crs.ID))
      assignmentGroupMatrix = append(assignmentGroupMatrix, *fetchAssignmentGroups(crs.ID))
    }
  }

  // one master grid per course
  var courseMasterGrids []ui.Grid

  // one grade grid per course 
  var courseGradeGrids []ui.Grid

  // one announcement grid per course 
  var courseAnnouncementGrids []ui.Grid

  // one course overview grid per course 
  var courseOverviewGrids []ui.Grid

  // one syllabus grid per course 
  var courseSyllabusGrids []ui.Grid

  // one assignment grid per course
  var courseAssignmentGrids []ui.Grid

  // declare tab widget
  tabpane := createMainTabPane(courses)

  dashboard := createDashboardGrid(&activeCourses, assignmentsMatrix)
  contentGrid := dashboard

  // Do the initial drawing of the main dash
  masterGrid = updateMasterGrid(masterGrid, tabpane, contentGrid)
  
  
  i := 0
  for _, crs := range *courses {
    if crs.EndAt.IsZero() && !crs.Term.EndAt.IsZero(){
      courseMasterGrids = append(courseMasterGrids, *createCourseGrid(crs, &assignmentsMatrix[i], &announcementMatrix[i], &assignmentGroupMatrix[i]))
      courseOverviewGrids = append(courseOverviewGrids, *createCourseOverviewGrid(crs, &assignmentsMatrix[i], &announcementMatrix[i], &assignmentGroupMatrix[i]))
      courseGradeGrids = append(courseGradeGrids, *createGradeGrid(crs, &assignmentsMatrix[i], &assignmentGroupMatrix[i]))
      courseAnnouncementGrids = append(courseAnnouncementGrids, *createAnnouncementGrid(crs, &announcementMatrix[i]))
      courseSyllabusGrids = append(courseSyllabusGrids, *createSyllabusGrid(crs))
      courseAssignmentGrids = append(courseAssignmentGrids, *createAssignmentGrid(crs, &assignmentsMatrix[i]))
      i++
    }
  }

  // log.Panic(courseMasterGrids)

  // Event polling loop
  tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
    select {
    case e := <-uiEvents:
      switch e.ID {
      case "q", "<C-c>":
        return
      case "o":
        // handleOpen(courseMasterGrids, courseOverviewGrids, courseGradeGrids, courseAnnouncementGrids, courseSyllabusGrids, courseAssignmentGrids, tabpane, masterGrid, contentGrid, *courses)
        handleOpen(courseMasterGrids, courseOverviewGrids, courseGradeGrids, courseAnnouncementGrids, courseSyllabusGrids, courseAssignmentGrids, tabpane, masterGrid, contentGrid, activeCourses)
      case "h":
        tabpane.FocusLeft() // changes the currently selected tab
        ui.Render(tabpane) // quickly redraws tabpane
      case "l":
        tabpane.FocusRight()
        ui.Render(tabpane)
      case "j":
        menuScroll(courseMasterGrids, tabpane, masterGrid, contentGrid, "down")
      case "k":
        menuScroll(courseMasterGrids, tabpane, masterGrid, contentGrid, "up")
      case "<Enter>":
        chooseTab(courseMasterGrids, tabpane, masterGrid, contentGrid)
      case "<Space>":
        handleSpace(courseMasterGrids, courseOverviewGrids, courseGradeGrids, courseAnnouncementGrids, courseSyllabusGrids, courseAssignmentGrids, tabpane, masterGrid, contentGrid)
      case "<Resize>":
				payload := e.Payload.(ui.Resize)
				masterGrid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(masterGrid)
      }
		case <-ticker:
      ui.Render(masterGrid)
			tickerCount++
		}
	}

}
