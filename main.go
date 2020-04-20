package main

import "fmt"
import "os"
import "github.com/spf13/viper"

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

  f, err := os.Open("config.yaml")
  if err != nil {
      panic(err)
  }
  defer f.Close()

  readConfig()

  var courses *[]Course = fetchCourses()

  // iterate through course structs
  for _, crs := range *courses {
    fmt.Println(crs.Name)
  }

}
