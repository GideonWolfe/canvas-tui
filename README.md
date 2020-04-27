# canvas-tui

canvas-tui is a terminal interface that allows you to view statistics scraped from the [Canvas API](https://canvas.instructure.com/doc/api/). 

![Some details obscured for privacy](https://i.imgur.com/Klao8nD.png)

The goal of this project is to not only allow easy viewing of upcoming assignments and due dates, but to provide new metrics and data visualizations that are not included in the default canvas interface. 

# Usage

The top bar is called the `tab bar`, and there should be one tab per enrolled course, plus a dashboard. You can use the `h` and `l` keys to navigate left and right respectively.

On a course page, you can use `j` and `k` to scroll down and up the list of pages provided by the canvas course. press `Space` to select a page. Currently only a few generic page types are supported for native viewing. I will soon implement an "open in browser" functionality for those pages that are impractical to display.

# TODO list

- [ ] Fix score line chart (waiting on PR in termui)
- [ ] Scrollable widgets (syllabus, grades, etc)
- [ ] Fix old classes (mostly labs) sticking around in courses
- [ ] Fix grade table formatting bug
