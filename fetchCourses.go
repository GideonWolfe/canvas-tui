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
	ID                               int            `json:"id"`
	Name                             string         `json:"name"`
	AccountID                        int            `json:"account_id"`
	UUID                             string         `json:"uuid"`
	StartAt                          time.Time      `json:"start_at"`
	GradingStandardID                interface{}    `json:"grading_standard_id"`
	IsPublic                         bool           `json:"is_public"`
	CreatedAt                        time.Time      `json:"created_at"`
	SyllabusBody                     string         `json:"syllabus_body"`
	CourseCode                       string         `json:"course_code"`
	DefaultView                      string         `json:"default_view"`
	RootAccountID                    int            `json:"root_account_id"`
	EnrollmentTermID                 int            `json:"enrollment_term_id"`
	License                          string         `json:"license"`
	PublicDescription                string         `json:"public_description"`
	GradePassbackSetting             interface{}    `json:"grade_passback_setting"`
	EndAt                            time.Time      `json:"end_at"`
	PublicSyllabus                   bool           `json:"public_syllabus"`
	PublicSyllabusToAuth             bool           `json:"public_syllabus_to_auth"`
	StorageQuotaMb                   int            `json:"storage_quota_mb"`
	IsPublicToAuthUsers              bool           `json:"is_public_to_auth_users"`
	CourseProgress                   CourseProgress `json:"course_progress"`
	ApplyAssignmentGroupWeights      bool           `json:"apply_assignment_group_weights"`
	Sections                         []Sections     `json:"sections"`
	TotalStudents                    int            `json:"total_students"`
	IsFavorite                       bool           `json:"is_favorite"`
	Teachers                         []Teachers     `json:"teachers"`
	Tabs                             []Tabs         `json:"tabs"`
	Calendar                         Calendar       `json:"calendar"`
	TimeZone                         string         `json:"time_zone"`
	ImageDownloadURL                 string         `json:"image_download_url"`
	Concluded                        bool           `json:"concluded"`
	Blueprint                        bool           `json:"blueprint"`
	Enrollments                      []Enrollments  `json:"enrollments"`
	HideFinalGrades                  bool           `json:"hide_final_grades"`
	WorkflowState                    string         `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool           `json:"restrict_enrollments_to_course_dates"`
}
type CourseProgress struct {
	RequirementCount          int         `json:"requirement_count"`
	RequirementCompletedCount int         `json:"requirement_completed_count"`
	NextRequirementURL        interface{} `json:"next_requirement_url"`
	CompletedAt               interface{} `json:"completed_at"`
}
type Sections struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	StartAt        interface{} `json:"start_at"`
	EndAt          interface{} `json:"end_at"`
	EnrollmentRole string      `json:"enrollment_role"`
}
type Teachers struct {
	ID             int         `json:"id"`
	DisplayName    string      `json:"display_name"`
	AvatarImageURL string      `json:"avatar_image_url"`
	HTMLURL        string      `json:"html_url"`
	Pronouns       interface{} `json:"pronouns"`
}
type Tabs struct {
	ID         string `json:"id"`
	HTMLURL    string `json:"html_url"`
	FullURL    string `json:"full_url"`
	Position   int    `json:"position"`
	Visibility string `json:"visibility"`
	Label      string `json:"label"`
	Type       string `json:"type"`
}
type Calendar struct {
	Ics string `json:"ics"`
}
type Enrollments struct {
	Type                           string      `json:"type"`
	Role                           string      `json:"role"`
	RoleID                         int         `json:"role_id"`
	UserID                         int         `json:"user_id"`
	EnrollmentState                string      `json:"enrollment_state"`
	LimitPrivilegesToCourseSection bool        `json:"limit_privileges_to_course_section"`
	ComputedCurrentGrade           interface{} `json:"computed_current_grade"`
	ComputedCurrentScore           float64     `json:"computed_current_score"`
	ComputedFinalGrade             interface{} `json:"computed_final_grade"`
	ComputedFinalScore             float64     `json:"computed_final_score"`
}

func fetchCourses() *[]Course {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string) + "api/v1/courses?per_page=60&enrollment_state=active"+
                                                                          "&include[]=syllabus_body"+
                                                                          "&include[]=total_scores"+
                                                                          "&include[]=public_description"+
                                                                          "&include[]=course_progress"+
                                                                          "&include[]=sections"+
                                                                          "&include[]=total_students"+
                                                                          "&include[]=favorites"+
                                                                          "&include[]=teachers"+
                                                                          "&include[]=tabs"+
                                                                          "&include[]=course_image"+
                                                                          "&include[]=concluded"

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
