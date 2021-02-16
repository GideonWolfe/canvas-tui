# canvas-tui

canvas-tui is a terminal interface that allows you to view statistics scraped from the [Canvas API](https://canvas.instructure.com/doc/api/). 

![](https://i.imgur.com/9W7CjQa.png)

![Some details obscured for privacy](https://i.imgur.com/Klao8nD.png)

The goal of this project is to not only allow easy viewing of upcoming assignments and due dates, but to provide new metrics and data visualizations that are not included in the default canvas interface. 

![grades](https://i.imgur.com/fd6Sz7t.png)

# Usage

The top bar is called the `tab bar`, and there should be one tab per enrolled course, plus a dashboard. You can use the `h` and `l` keys to navigate left and right respectively, and `Enter` to make a selection.

On a course page, you can use `j` and `k` to scroll down and up the list of pages provided by the canvas course. press `Space` to select a page. Currently only a few generic page types are supported for native viewing. To open any page in the browser, simply press `o` while viewing that page.

# Configuration

TO generate a new canvas token, head to your Canvas Settings -> Approved Integrations -> Generate New Token.

Name it whatever you want, but save the generated token. The `config.yaml` file should be in `$HOME/.config/canvas-tui`, with the following contents:

```yaml
canvasdomain: "https://your.schools.canvas.com"
canvastoken: "token from earlier"
```
Yes, it is quite slow to start up for a TUI application. This is because the **vast** majority of API calls and data processing are happening at startup, so it is limited to the speed of the API. However once loaded, it should be quite snappy.


# Installation

Run the following commands to install the dependencies:

* `GO111MODULE=on go get github.com/spf13/viper`
* `go get github.com/gizak/termui`
* `go get github.com/gizak/termui/widgets`
* `go get github.com/grokify/html-strip-tags-go`

# TODO list

- [ ] Fix score line chart (waiting on PR in termui)
- [ ] Scrollable widgets (syllabus, grades, etc, waiting on PR in termui)
- [x] Fix old classes (mostly labs) sticking around in courses
- [x] Fix grade table formatting bug
- [x] Open pages in browser
- [ ] Add labels to grade plot (waiting on PR in termui)
