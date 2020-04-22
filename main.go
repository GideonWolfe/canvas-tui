package main

import (
	"log"
  "fmt"
  "time"
  "os"
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


func renderGrid() *ui.Grid {
  // Dummy placeholder widget
  p0 := widgets.NewParagraph()
	p0.Text = "Some Text"
	p0.Border = true

	newGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	newGrid.SetRect(0, 0, termWidth, termHeight)
  newGrid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, p0),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, p0),
		),
  )

  return newGrid
}

func main() {
  
  // Initialize temui
  if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()


  f, err := os.Open("config.yaml")
  if err != nil {
      panic(err)
  }
  defer f.Close()

  readConfig()

  
  // declare grid and set terminal dimensions
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

  // Dummy placeholder widget
  p0 := widgets.NewParagraph()
	p0.Text = "Some Text"
	p0.Border = true

  // Titles for the tab widget
  var titles []string
  var courses *[]Course = fetchCourses()
  for _, crs := range *courses {
    if crs.EndAt.IsZero() {
      titles = append(titles, crs.CourseCode)
    }
  }

  // declare tab widget
  tabpane := widgets.NewTabPane(titles...)
	tabpane.Border = true

  // defining master grid layout
  grid.Set(
		ui.NewRow(1.0/20,
			ui.NewCol(1.0, tabpane),
		),
		ui.NewRow(19.0/20,
			ui.NewCol(1.0/1, renderGrid()),
		),
  )
  

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p0)
		case 1:
			ui.Render(p0)
		case 2:
			ui.Render(p0)
		}
	}

	ui.Render(grid)

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
        tabpane.FocusLeft()
        ui.Clear()
        ui.Render(tabpane)
        renderTab()
      case "l":
        tabpane.FocusRight()
        ui.Clear()
        ui.Render(tabpane)
        renderTab()
      case "<Enter>":
        return
      case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
      }
		case <-ticker:
      ui.Render(grid)
			tickerCount++
		}
	}

}
