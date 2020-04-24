package main

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "github.com/spf13/viper"
  // "log"
	// "strconv"
  // "time"
	// "reflect"
	// "time"
	// "bytes"
  // "fmt"
  strip "github.com/grokify/html-strip-tags-go"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)


func createAnnouncementParagraph(announcement Announcement) *widgets.Paragraph {

  paragraph := widgets.NewParagraph()
  paragraph.Text = announcement.Message

  return paragraph
}

func createAnnouncementTable(announcements *[]Announcement, width int) *widgets.Table {

  var tableData [][]string
  header := []string{"Message"}
  tableData = append(tableData, header)
  for _, ann := range *announcements {
        var assignmentData []string
        // assignmentData = append(assignmentData, ann.PostedAt.Local().Format("Jan 2"))
        assignmentData = append(assignmentData, strip.StripTags(ann.Message))
        tableData = append(tableData, assignmentData)
  }
 
  assignmentTable := widgets.NewTable()
  assignmentTable.Title = "Announcements:"
  assignmentTable.Rows = tableData
  assignmentTable.TextStyle = ui.NewStyle(ui.ColorWhite)
  assignmentTable.RowSeparator = true
  assignmentTable.FillRow = false
  // assignmentTable.RowStyles
  // assignmentTable.ColumnWidths = []int{10, width}
  assignmentTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  return assignmentTable

}


func createAnnouncementGrid(course Course) *ui.Grid {

  var announcements *[]Announcement = fetchAnnouncements(course.ID)
  // var announcmentParagraphs []widgets.Paragraph

  announcementGrid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	announcementGrid.SetRect(0, 0, termWidth, termHeight)
  announcementGrid.Title = "Announcement Grid"
  announcementGrid.Set(
    ui.NewCol(1.0, 
      ui.NewRow(1.0, createAnnouncementTable(announcements, termWidth)),
      // ui.NewRow(1.0, &announcmentParagraphs[0]),
		),
  )
  // for _, ann := range *announcements {
    // announcmentParagraphs = append(announcmentParagraphs, *createAnnouncementParagraph(ann))
  // }

  return announcementGrid
}
