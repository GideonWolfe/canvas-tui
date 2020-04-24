package main

import (
	"log"
	// "math"
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

// called to handle when a user clicks a different tab
func  handleChoice(coursePages []ui.Grid, tabpane *widgets.TabPane, masterGrid *ui.Grid, contentGrid *ui.Grid) {
  switch tabpane.ActiveTabIndex {
  case 0:
    // contentGrid = createDashboardGrid("Dashboard YO!!!!!!")
    masterGrid.Items[1].Entry = contentGrid
    // masterGrid = updateMasterGrid(masterGrid,tabpane,contentGrid)
    ui.Render(masterGrid)
  case 1:
    contentGrid = &coursePages[tabpane.ActiveTabIndex-1]
    masterGrid.Items[1].Entry = contentGrid
    // masterGrid = updateMasterGrid(masterGrid,tabpane,contentGrid)
    ui.Render(masterGrid)
  case 2:
    contentGrid = &coursePages[tabpane.ActiveTabIndex-1]
    masterGrid.Items[1].Entry = contentGrid
    // masterGrid = updateMasterGrid(masterGrid,tabpane,contentGrid)
    ui.Render(masterGrid)
  case 3:
    contentGrid = &coursePages[tabpane.ActiveTabIndex-1]
    masterGrid.Items[1].Entry = contentGrid
    // masterGrid = updateMasterGrid(masterGrid,tabpane,contentGrid)
    ui.Render(masterGrid)
  }
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
  
  // declare master grid and set terminal dimensions
	masterGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	masterGrid.SetRect(0, 0, termWidth, termHeight)

  // declare tab widget
  tabpane := createMainTabPane(courses)

  // contentGrid := createDashboardGrid("front page")
  dashboard := createDashboardGrid(courses)
  contentGrid := dashboard

  // Do the initial drawing of the main dash
  masterGrid = updateMasterGrid(masterGrid, tabpane, contentGrid)

  var coursePages []ui.Grid
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      coursePages = append(coursePages, *createCourseGrid(crs))
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
      case "j":
        // variable that points to the list we're on
        // l := contentGrid.Items[0].Entry.(*widgets.List)
        // l.ScrollDown()
				ui.Render(masterGrid)
      case "k":
        // variable that points to the list we're on
        // l := contentGrid.Items[0].Entry.(*widgets.List)
        // l.ScrollUp()
				ui.Render(masterGrid)
      case "<Enter>":
        handleChoice(coursePages, tabpane, masterGrid, contentGrid)
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
