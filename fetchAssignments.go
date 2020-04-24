package main

import (
  "encoding/json"
  "io/ioutil"
	"log"
	"net/http"
  "github.com/spf13/viper"
  // "fmt"
  "time"
  "strconv"
)


type Assignment struct {
	ID                              int              `json:"id"`
	Description                     string           `json:"description"`
	DueAt                           time.Time        `json:"due_at"`
	UnlockAt                        interface{}      `json:"unlock_at"`
	LockAt                          time.Time        `json:"lock_at"`
	PointsPossible                  float64          `json:"points_possible"`
	GradingType                     string           `json:"grading_type"`
	AssignmentGroupID               int              `json:"assignment_group_id"`
	GradingStandardID               interface{}      `json:"grading_standard_id"`
	CreatedAt                       time.Time        `json:"created_at"`
	UpdatedAt                       time.Time        `json:"updated_at"`
	PeerReviews                     bool             `json:"peer_reviews"`
	AutomaticPeerReviews            bool             `json:"automatic_peer_reviews"`
	Position                        int              `json:"position"`
	GradeGroupStudentsIndividually  bool             `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool             `json:"anonymous_peer_reviews"`
	GroupCategoryID                 interface{}      `json:"group_category_id"`
	PostToSis                       bool             `json:"post_to_sis"`
	ModeratedGrading                bool             `json:"moderated_grading"`
	OmitFromFinalGrade              bool             `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool             `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool             `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool             `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool             `json:"graders_anonymous_to_graders"`
	GraderCount                     int              `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool             `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   interface{}      `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool             `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int              `json:"allowed_attempts"`
	LockInfo                        LockInfo         `json:"lock_info"`
	SecureParams                    string           `json:"secure_params"`
	CourseID                        int              `json:"course_id"`
	Name                            string           `json:"name"`
	SubmissionTypes                 []string         `json:"submission_types"`
	HasSubmittedSubmissions         bool             `json:"has_submitted_submissions"`
	DueDateRequired                 bool             `json:"due_date_required"`
	MaxNameLength                   int              `json:"max_name_length"`
	InClosedGradingPeriod           bool             `json:"in_closed_grading_period"`
	VericiteEnabled                 bool             `json:"vericite_enabled"`
	VericiteSettings                VericiteSettings `json:"vericite_settings"`
	IsQuizAssignment                bool             `json:"is_quiz_assignment"`
	CanDuplicate                    bool             `json:"can_duplicate"`
	OriginalCourseID                interface{}      `json:"original_course_id"`
	OriginalAssignmentID            interface{}      `json:"original_assignment_id"`
	OriginalAssignmentName          interface{}      `json:"original_assignment_name"`
	OriginalQuizID                  interface{}      `json:"original_quiz_id"`
	WorkflowState                   string           `json:"workflow_state"`
	Muted                           bool             `json:"muted"`
	HTMLURL                         string           `json:"html_url"`
	QuizID                          int              `json:"quiz_id"`
	AnonymousSubmissions            bool             `json:"anonymous_submissions"`
	AllDates                        []AllDates       `json:"all_dates"`
	Published                       bool             `json:"published"`
	OnlyVisibleToOverrides          bool             `json:"only_visible_to_overrides"`
	Submission                      Submission       `json:"submission"`
	LockedForUser                   bool             `json:"locked_for_user"`
	LockExplanation                 string           `json:"lock_explanation"`
	SubmissionsDownloadURL          string           `json:"submissions_download_url"`
	PostManually                    bool             `json:"post_manually"`
	AnonymizeStudents               bool             `json:"anonymize_students"`
	RequireLockdownBrowser          bool             `json:"require_lockdown_browser"`
  ExternalToolTagAttributes       ExternalToolTagAttributes `json:"external_tool_tag_attributes,omitempty"`
  URL                             string                    `json:"url,omitempty"`
}
type LockInfo struct {
	LockAt      time.Time `json:"lock_at"`
	CanView     bool      `json:"can_view"`
	AssetString string    `json:"asset_string"`
}
type VericiteSettings struct {
	OriginalityReportVisibility string `json:"originality_report_visibility"`
	ExcludeQuoted               bool   `json:"exclude_quoted"`
	ExcludeSelfPlag             bool   `json:"exclude_self_plag"`
	StoreInIndex                bool   `json:"store_in_index"`
}
type AllDates struct {
	DueAt    time.Time   `json:"due_at"`
	UnlockAt interface{} `json:"unlock_at"`
	LockAt   time.Time   `json:"lock_at"`
	Base     bool        `json:"base"`
}
type Submission struct {
	ID                            int         `json:"id"`
	Body                          string      `json:"body"`
	URL                           interface{} `json:"url"`
	Grade                         string      `json:"grade"`
	Score                         float64     `json:"score"`
	SubmittedAt                   time.Time   `json:"submitted_at"`
	AssignmentID                  int         `json:"assignment_id"`
	UserID                        int         `json:"user_id"`
	SubmissionType                string      `json:"submission_type"`
	WorkflowState                 string      `json:"workflow_state"`
	GradeMatchesCurrentSubmission bool        `json:"grade_matches_current_submission"`
	GradedAt                      time.Time   `json:"graded_at"`
	GraderID                      int         `json:"grader_id"`
	Attempt                       int         `json:"attempt"`
	CachedDueDate                 time.Time   `json:"cached_due_date"`
	Excused                       interface{} `json:"excused"`
	LatePolicyStatus              interface{} `json:"late_policy_status"`
	PointsDeducted                interface{} `json:"points_deducted"`
	GradingPeriodID               interface{} `json:"grading_period_id"`
	ExtraAttempts                 interface{} `json:"extra_attempts"`
	PostedAt                      time.Time   `json:"posted_at"`
	Late                          bool        `json:"late"`
	Missing                       bool        `json:"missing"`
	SecondsLate                   int         `json:"seconds_late"`
	EnteredGrade                  string      `json:"entered_grade"`
	EnteredScore                  float64     `json:"entered_score"`
	PreviewURL                    string      `json:"preview_url"`
}
type ExternalToolTagAttributes struct {
  URL            string `json:"url"`
  NewTab         bool   `json:"new_tab"`
  ResourceLinkID string `json:"resource_link_id"`
}



func fetchAssignments(courseID int) *[]Assignment {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string)+"api/v1/courses/"+strconv.Itoa(courseID)+"/assignments?per_page=100&include[]=submission&include[]=all_dates"
  // url := viper.Get("canvasdomain").(string)+"api/v1/courses/"+strconv.Itoa(courseID)+"/assignments?per_page=100&include[]=submission&include=all_dates"

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
  assignments := make([]Assignment,0)
  err = json.Unmarshal([]byte(body), &assignments)
  if err != nil {
    panic(err)
  }
  
  return &assignments
	
}
