package main

import (
  "encoding/json"
  "io/ioutil"
	"log"
	"net/http"
  "github.com/spf13/viper"
  "time"
  "strconv"
)

type Announcement struct {
	ID                      int           `json:"id"`
	Title                   string        `json:"title"`
	LastReplyAt             time.Time     `json:"last_reply_at"`
	CreatedAt               time.Time     `json:"created_at"`
	DelayedPostAt           interface{}   `json:"delayed_post_at"`
	PostedAt                time.Time     `json:"posted_at"`
	AssignmentID            interface{}   `json:"assignment_id"`
	RootTopicID             interface{}   `json:"root_topic_id"`
	Position                int           `json:"position"`
	PodcastHasStudentPosts  bool          `json:"podcast_has_student_posts"`
	DiscussionType          string        `json:"discussion_type"`
	LockAt                  interface{}   `json:"lock_at"`
	AllowRating             bool          `json:"allow_rating"`
	OnlyGradersCanRate      bool          `json:"only_graders_can_rate"`
	SortByRating            bool          `json:"sort_by_rating"`
	IsSectionSpecific       bool          `json:"is_section_specific"`
	UserName                string        `json:"user_name"`
	DiscussionSubentryCount int           `json:"discussion_subentry_count"`
	Permissions             Permissions   `json:"permissions"`
	RequireInitialPost      interface{}   `json:"require_initial_post"`
	UserCanSeePosts         bool          `json:"user_can_see_posts"`
	PodcastURL              interface{}   `json:"podcast_url"`
	ReadState               string        `json:"read_state"`
	UnreadCount             int           `json:"unread_count"`
	Subscribed              bool          `json:"subscribed"`
	Attachments             []interface{} `json:"attachments"`
	Published               bool          `json:"published"`
	CanUnpublish            bool          `json:"can_unpublish"`
	Locked                  bool          `json:"locked"`
	CanLock                 bool          `json:"can_lock"`
	CommentsDisabled        bool          `json:"comments_disabled"`
	Author                  Author        `json:"author"`
	HTMLURL                 string        `json:"html_url"`
	URL                     string        `json:"url"`
	Pinned                  bool          `json:"pinned"`
	GroupCategoryID         interface{}   `json:"group_category_id"`
	CanGroup                bool          `json:"can_group"`
	TopicChildren           []interface{} `json:"topic_children"`
	GroupTopicChildren      []interface{} `json:"group_topic_children"`
	ContextCode             string        `json:"context_code"`
	LockedForUser           bool          `json:"locked_for_user"`
	LockInfo                LockInfo      `json:"lock_info"`
	LockExplanation         string        `json:"lock_explanation"`
	Message                 string        `json:"message"`
	SubscriptionHold        string        `json:"subscription_hold"`
	TodoDate                interface{}   `json:"todo_date"`
}
type Permissions struct {
	Attach bool `json:"attach"`
	Update bool `json:"update"`
	Reply  bool `json:"reply"`
	Delete bool `json:"delete"`
}
type Author struct {
	ID             int         `json:"id"`
	DisplayName    string      `json:"display_name"`
	AvatarImageURL string      `json:"avatar_image_url"`
	HTMLURL        string      `json:"html_url"`
	Pronouns       interface{} `json:"pronouns"`
}

func fetchAnnouncements(courseID int) *[]Announcement {

  // Create URL string from config file
  url := viper.Get("canvasdomain").(string)+"api/v1/announcements?context_codes[]=course_"+strconv.Itoa(courseID)

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
  announcements := make([]Announcement,0)
  err = json.Unmarshal([]byte(body), &announcements)
  if err != nil {
    panic(err)
  }
  
  return &announcements
	
}
