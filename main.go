package main

import (
	"log"
  "fmt"
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

  // var courses *[]Course = fetchCourses()
  // var assignments *[]Assignment = fetchAssignments(1263956)
  // fetchAssignments(1263956)

  // Phony bar graph to show
  bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 35, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

  var titles []string
  var courses *[]Course = fetchCourses()
  for _, crs := range *courses {
    titles = append(titles, crs.CourseCode)
  }
  
  tabpane := widgets.NewTabPane(titles...)
	tabpane.SetRect(0, 1, 180, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(bc)
		case 1:
			ui.Render(bc)
		}
	}

	ui.Render(tabpane)

	uiEvents := ui.PollEvents()
  // Event polling loop
	for {
		e := <-uiEvents
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
		}
	}
}
