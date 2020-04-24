package main

import (
  "encoding/json"
  "io/ioutil"
	"log"
	"net/http"
  "github.com/spf13/viper"
  // "fmt"
  // "time"
  "strconv"
)

type AssignmentGroup struct {
  ID              int             `json:"id"`
  Name            string          `json:"name"`
  Position        int             `json:"position"`
  GroupWeight     float64         `json:"group_weight"`
  SisSourceID     interface{}     `json:"sis_source_id"`
  IntegrationData IntegrationData `json:"integration_data"`
  Rules           Rules           `json:"rules"`
  Assignments     []Assignment    `json:"assignments"`
  AnyAssignmentInClosedGradingPeriod bool            `json:"any_assignment_in_closed_grading_period"`
}

type IntegrationData struct {
}
type Rules struct {
	DropLowest int   `json:"drop_lowest"`
	NeverDrop  []int `json:"never_drop"`
}

func fetchAssignmentGroups(courseID int) *[]AssignmentGroup {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string)+"api/v1/courses/"+strconv.Itoa(courseID)+"/assignments?per_page=100&include[]=submission&include=all_dates"

  // Create a Bearer string by appending string access token
  var bearer = "Bearer " + viper.Get("canvastoken").(string)

  // Create a new request using http
  req, err := http.NewRequest("GET", url, nil)

  // add authorization header to the req
  req.Header.Add("Authorization", bearer)

  // Send req using http Client
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      log.Println("Error on response.\n[ERRO] -", err)
  }

  body, _ := ioutil.ReadAll(resp.Body)
  assignmentGroups := make([]AssignmentGroup,0)
  err = json.Unmarshal([]byte(body), &assignmentGroups)
  if err != nil {
    panic(err)
  }
  
  // iterate through assignment structs
  // for _, assignment := range assignmnts {
    // fmt.Println(crs.Name)
  // }
  // fmt.Println(assignments)
  return &assignmentGroups
	
}
