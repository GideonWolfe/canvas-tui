package main

import (
	"log"
	"math"
  "fmt"
  "time"
  "github.com/spf13/viper"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Reads the config file
func readConfig() {
  viper.SetConfigName("config") // name of config file (without extension)
  viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
  viper.AddConfigPath(".")               // optionally look for config in the working directory
  // viper.AddConfigPath("$HOME/.config/canvas-tui/")              
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
    if crs.EndAt.IsZero() {
      titles = append(titles, crs.CourseCode)
    }
  }
  tabpane := widgets.NewTabPane(titles...)
	tabpane.Border = true
  return tabpane
}

// based on an input course object, this function generates 
// a grid with widgets populated with data from the course
func renderCourseGrid(someVal string) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
	p0.Text = someVal
	p0.Border = true

  // pie chart to eventually break down course points
  pc := widgets.NewPieChart()
	pc.Title = "Pie Chart"
	pc.Data = []float64{.10, .10, .05, .20, .05, .13, .14, .25}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

	courseGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	courseGrid.SetRect(0, 0, termWidth, termHeight)
  courseGrid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, pc),
			ui.NewCol(3.0/4, p0),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, p0),
		),
  )
  return courseGrid
}

func createDashboardGrid(someVal string) *ui.Grid {
  // dummy placeholder widget
  p0 := widgets.NewParagraph()
	p0.Text = someVal
	p0.Border = true

  bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
  
	dashboardGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	dashboardGrid.SetRect(0, 0, termWidth, termHeight)
  dashboardGrid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, bc),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, p0),
		),
  )
  return dashboardGrid
}


func  handleChoice(coursePages []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid) {
  switch tabpane.ActiveTabIndex {
  case 0:
    // updateMasterGrid(masterGrid, tabpane, "dashboard")
    // ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex])
  case 1:
    ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex-1])
    // updateMasterGrid(masterGrid, tabpane, "course")
  case 2:
    ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex-1])
    // updateMasterGrid(masterGrid, tabpane, "course")
  case 3:
    ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex-1])
    // updateMasterGrid(masterGrid, tabpane, "course")
  }
}




// called if master grid needs to be updated
func updateMasterGrid(masterGrid *ui.Grid, tabpane *widgets.TabPane, contentGrid *ui.Grid) {
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
  ui.Render(masterGrid)
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
  
  // declare master grid and set terminal dimensions
	masterGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	masterGrid.SetRect(0, 0, termWidth, termHeight)

  // declare tab widget
  tabpane := createMainTabPane(courses)

  contentGrid := createDashboardGrid("suh dude")

  // Do the initial drawing of the main dash
  updateMasterGrid(masterGrid, tabpane, contentGrid)

  // handleChoice(coursePages) := func() {
    // switch tabpane.ActiveTabIndex {
    // case 0:
      // // updateMasterGrid(masterGrid, tabpane, "dashboard")
      // ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex])
    // case 1:
      // ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex])
      // // updateMasterGrid(masterGrid, tabpane, "course")
    // case 2:
      // ui.Render(masterGrid, &coursePages[tabpane.ActiveTabIndex])
      // // updateMasterGrid(masterGrid, tabpane, "course")
    // }
  // }

  // ui.Render(masterGrid)


  var coursePages []ui.Grid
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      coursePages = append(coursePages, *renderCourseGrid(crs.Name))
    }
  }


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
      case "h":
        tabpane.FocusLeft() // changes the currently selected tab
        ui.Render(tabpane) // quickly redraws tabpane
      case "l":
        tabpane.FocusRight()
        ui.Render(tabpane)
      case "<Enter>":
        ui.Clear() // Clear what we currently are displaying
        contentGrid = renderCourseGrid("u clicked m8")
        // updateMasterGrid(masterGrid, tabpane, testGrid) // TODO Master grid doesn't clear previous grid?
        // ui.Render(masterGrid, &coursePages[1])
        handleChoice(coursePages, tabpane, masterGrid)
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
