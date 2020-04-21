package main

import (
  "encoding/json"
  "io/ioutil"
	"log"
	"net/http"
  "github.com/spf13/viper"
  // "fmt"
  "time"
  "bytes"
)

type Course struct {
	ID                               int           `json:"id"`
	Name                             string        `json:"name"`
	AccountID                        int           `json:"account_id"`
	UUID                             string        `json:"uuid"`
	StartAt                          time.Time     `json:"start_at"`
	GradingStandardID                interface{}   `json:"grading_standard_id"`
	IsPublic                         bool          `json:"is_public"`
	CreatedAt                        time.Time     `json:"created_at"`
	CourseCode                       string        `json:"course_code"`
	DefaultView                      string        `json:"default_view"`
	RootAccountID                    int           `json:"root_account_id"`
	EnrollmentTermID                 int           `json:"enrollment_term_id"`
	License                          string        `json:"license"`
	GradePassbackSetting             interface{}   `json:"grade_passback_setting"`
  EndAt                            time.Time     `json:"end_at"`
	PublicSyllabus                   bool          `json:"public_syllabus"`
	PublicSyllabusToAuth             bool          `json:"public_syllabus_to_auth"`
	StorageQuotaMb                   int           `json:"storage_quota_mb"`
	IsPublicToAuthUsers              bool          `json:"is_public_to_auth_users"`
	ApplyAssignmentGroupWeights      bool          `json:"apply_assignment_group_weights"`
	Calendar                         Calendar      `json:"calendar"`
	TimeZone                         string        `json:"time_zone"`
	Blueprint                        bool          `json:"blueprint"`
	Enrollments                      []Enrollments `json:"enrollments"`
	HideFinalGrades                  bool          `json:"hide_final_grades"`
	WorkflowState                    string        `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool          `json:"restrict_enrollments_to_course_dates"`
	OverriddenCourseVisibility       string        `json:"overridden_course_visibility,omitempty"`
}
type Calendar struct {
	Ics string `json:"ics"`
}
type Enrollments struct {
	Type                           string `json:"type"`
	Role                           string `json:"role"`
	RoleID                         int    `json:"role_id"`
	UserID                         int    `json:"user_id"`
	EnrollmentState                string `json:"enrollment_state"`
	LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
}


func fetchCourses() *[]Course {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string) + "api/v1/courses?per_page=60&enrollment_state=active"
  // url := viper.Get("canvasdomain").(string) + "api/v1/courses/"

  // Create a Bearer string by appending string access token
  var bearer = "Bearer " + viper.Get("canvastoken").(string)

  type Params struct{
    PerPage int `json:"per_page"`
    // EnrollmentType string `json:"enrollment_type"`
  }

  m := Params{50}
  rawbody, err := json.Marshal(m)
  // fmt.Println(string(rawbody))
  params := bytes.NewReader(rawbody)
  // params := bytes.NewBuffer(rawbody)

  // Create a new request using http
  req, err := http.NewRequest("GET", url, params)

  // add authorization header to the req
  req.Header.Add("Authorization", bearer)

    // Send req using http Client
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      log.Println("Error on response.\n[ERRO] -", err)
  }

  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  // fmt.Println(string(body))
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
