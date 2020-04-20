package main

import (
  "encoding/json"
  "io/ioutil"
	"log"
	"net/http"
  "github.com/spf13/viper"
  // "fmt"
)

type Calendar struct {
  Url string `json:"ics,omitempty"`
}

type Enrollments struct {
  Type string `json:"type,omitempty"`
  Role string `json:"role,omitempty"`
  RoleID int `json:"role_id,omitempty"`
  UserID int `json:"user_id,omitempty"`
  EnrollmentState string `json:"enrollment_state,omitempty"`
  Limit_privileges_to_course_section bool `json:"limit_privileges_to_course_section,omitempty"`
}

type Course struct {
  Id int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Account_id int `json:"account_id,omitempty"`
  Uuid string `json:"uuid,omitempty"`
  Start_at string `json:"start_at,omitempty"`
  Grading_standard_id string `json:"grading_standard_id,omitempty"`
  Is_public bool `json:"is_public,omitempty"`
  Created_at string `json:"created_at,omitempty"`
  Course_code string `json:"course_code,omitempty"`
  Default_view string `json:"default_view,omitempty"`
  Root_account_id int `json:"root_account_id,omitempty"`
  Enrollment_term_id int`json:"enrollment_term_id,omitempty"`
  License string `json:"license,omitempty"`
  Grade_passback_settting string `json:"grade_passback_settting,omitempty"`
  End_at string `json:"end_at,omitempty"`
  Public_syllabus bool `json:"public_syllabus,omitempty"`
  Public_syllabus_to_auth bool `json:"public_syllabus_to_auth,omitempty"`
  Storage_quota_mb int `json:"storage_quota_mb,omitempty"`
  Is_public_to_auth_users bool `json:"is_public_to_auth_users,omitempty"`
  Apply_assignment_group_weights bool `json:"apply_assignment_group_weights,omitempty"`
  Coursecalendar Calendar `json:"calendar"`
  Time_zone string `json:"time_zone,omitempty"`
  Blueprint bool `json:"blueprint,omitempty"`
  // Courseenrollments Enrollments `json:"enrollments,omitempty"`
  Hide_final_grades bool `json:"hide_final_grades,omitempty"`
  Workflow_state string `json:"workflow_state,omitempty"`
  Restrict_enrollments_to_course_dates bool `json:"restrict_enrollments_to_course_dates,omitempty"`
}

type Response struct {
  Collection []Course
}


func fetchCourses() *[]Course {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string) + "api/v1/courses/"

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
  courses := make([]Course,0)
  err = json.Unmarshal([]byte(body), &courses)
  if err != nil {
    panic(err)
  }
  
  // iterate through course structs
  // for _, crs := range courses {
    // fmt.Println(crs.Name)
  // }

  return &courses
	
}
