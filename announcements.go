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
	// ui "github.com/GideonWolfe/termui/v3"
	// "github.com/GideonWolfe/termui/v3/widgets"
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
        var announcementData []string
        announcementData = append(announcementData, strip.StripTags(ann.Message))
        tableData = append(tableData, announcementData)
  }
 
  announcementData := widgets.NewTable()
  announcementData.Title = "Announcements:"
  announcementData.Rows = tableData
  announcementData.TextStyle = ui.NewStyle(ui.ColorWhite)
  announcementData.RowSeparator = true
  announcementData.FillRow = false
  // assignmentTable.RowStyles
  // assignmentTable.ColumnWidths = []int{10, width}
  announcementData.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
  return announcementData

}


func createAnnouncementGrid(course Course, announcements *[]Announcement) *ui.Grid {

  // var announcements *[]Announcement = fetchAnnouncements(course.ID)
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
