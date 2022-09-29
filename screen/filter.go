package screen

import (
	"fmt"
	"grover/code"
	"grover/syntax"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var grid = ui.NewGrid()
var fileList = widgets.NewList()
var area = widgets.NewList()
var tab = "flavors"
var insertMode = false

type FilterScreen struct {
	Paths []*Path
}

func OnePath(path *Path) {
	existing := LoadPaths()
	existing = append([]*Path{path}, existing...)
	WritePaths(existing)
	Setup(existing)
}

func Filter(paths []*Path) {
	Setup(paths)
}

func Setup(paths []*Path) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	fs := FilterScreen{}
	fs.Paths = paths

	setListColors(fileList)
	setListColors(area)

	for _, p := range paths {
		fileList.Rows = append(fileList.Rows, p.Filename)
	}

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, fileList),
			ui.NewCol(0.66, area),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if insertMode {
			handleInsert(e)
		} else {
			fs.normalEvents(e)
		}
		ui.Render(grid)
	}
}

func handleInsert(e ui.Event) {
	if e.ID == "<Enter>" || e.ID == "<Escape>" {
		insertMode = false
	} else if e.ID == "<Backspace>" {
	} else if e.ID == "<Space>" {
	} else {
	}
}

func (fs *FilterScreen) normalEvents(e ui.Event) {
	switch e.ID {
	case "q", "<C-c>":
		ui.Close()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		os.Exit(1)
		return
	case "s":
	case "j", "<Down>":
		selectedList().ScrollDown()
	case "k", "<Up>":
		selectedList().ScrollUp()
	case "i":
		insertMode = true
	case "<Right>":
	case "<Left>":
	case "<Tab>":
	case "<Enter>":
		fs.handleEnter()
	case "<Resize>":
		payload := e.Payload.(ui.Resize)
		grid.SetRect(0, 0, payload.Width, payload.Height)
		ui.Clear()
	}
}

func selectedList() *widgets.List {
	if tab == "flavors" {
		return fileList
	} else if tab == "selected" {
		return fileList
	}

	return fileList
}

func (fs *FilterScreen) handleEnter() {
	p := fs.Paths[fileList.SelectedRow]
	lines := code.ReadFile(p.Fullpath)
	area.Rows = []string{}
	//for i := 0; i < 10; i++ {
	for _, line := range lines {
		area.Rows = append(area.Rows, syntax.Highlight(line))
	}
}

func setListColors(s *widgets.List) {
	s.SelectedRowStyle.Fg = ui.ColorWhite
	s.SelectedRowStyle.Bg = ui.ColorMagenta
	s.TextStyle.Fg = ui.ColorWhite
	s.TextStyle.Bg = ui.ColorBlack
}
