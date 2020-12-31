package mooze

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"os"
)

type MoozeScreen struct {
	s tcell.Screen
	r *Renderer
}

func NewMoozeScreen() *MoozeScreen {
	return &MoozeScreen{}
}

type MoozeWindow struct {
	x        int
	y        int
	sizeX    int
	sizeY    int
	hasTitle bool
	title    string
}

func NewMoozeWindow(x, y, sizeX, sizeY int, t bool) *MoozeWindow {
	return &MoozeWindow{
		x:        x,
		y:        y,
		sizeX:    sizeX,
		sizeY:    sizeY,
		hasTitle: t,
	}
}

func (w *MoozeWindow) Title(t string) {
	w.hasTitle = true
	w.title = t
}

func (m *MoozeScreen) InitScreen(mouse bool) {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = s.Init(); err != nil {
		panic(err)
	}
	if mouse {
		s.EnableMouse()
	} else {
		s.DisableMouse()
	}
	m.s = s
}

func (m *MoozeScreen) DefaultStyle() tcell.Style {
	return tcell.StyleDefault
}

func (m *MoozeScreen) Size() (int, int) {
	return m.s.Size()
}

func (m *MoozeScreen) Print(y, x int, str string, style tcell.Style) {
	for _, c := range str {
		w := runeWidth(c)
		var comb []rune
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		m.s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func (m *MoozeScreen) RenderWindow(w *MoozeWindow, style tcell.Style) {
	for col := w.y; col <= w.y+w.sizeY; col++ {
		m.s.SetContent(col, w.x, tcell.RuneHLine, nil, style)
		m.s.SetContent(col, w.x+w.sizeX, tcell.RuneHLine, nil, style)
	}
	for row := w.x + 1; row < w.x+w.sizeX; row++ {
		m.s.SetContent(w.y, row, tcell.RuneVLine, nil, style)
		m.s.SetContent(w.y+w.sizeY, row, tcell.RuneVLine, nil, style)
	}
	if w.sizeY != 0 && w.sizeX != 0 {
		m.s.SetContent(w.y, w.x, tcell.RuneULCorner, nil, style)
		m.s.SetContent(w.y+w.sizeY, w.x, tcell.RuneURCorner, nil, style)
		m.s.SetContent(w.y, w.x+w.sizeX, tcell.RuneLLCorner, nil, style)
		m.s.SetContent(w.y+w.sizeY, w.x+w.sizeX, tcell.RuneLRCorner, nil, style)
	}
	for row := w.x + 1; row < w.x+w.sizeX; row++ {
		for col := w.y + 1; col < w.y+w.sizeY; col++ {
			m.s.SetContent(col, row, ' ', nil, style)
		}
	}
	if w.hasTitle && len(w.title) < w.sizeY {
		m.Print(w.x, w.y+1, w.title, style)
	}
}

func (m *MoozeScreen) Clear() {
	m.s.Clear()
}

func (m *MoozeScreen) Show() {
	m.s.Show()
}

func (m *MoozeScreen) Sync() {
	m.s.Sync()
}

func (m *MoozeScreen) Reload() {
	m.r.ClearConsoleUnix()
	m.s.Show()
	m.s.Sync()
}

func (m *MoozeScreen) EmitEvent() tcell.Event {
	return m.s.PollEvent()
}

func runeWidth(r rune) int {
	return runewidth.RuneWidth(r)
}

func GetColor(n string) tcell.Color {
	return tcell.ColorNames[n]
}

func ToStyle(f ...string) tcell.Style {
	s := tcell.StyleDefault
	if len(f) == 2 {
		return s.Foreground(GetColor(f[0])).
			Background(GetColor(f[1]))
	} else {
		return s.Foreground(GetColor(f[0]))
	}
}

func (m *MoozeScreen) Exit(code int) {
	m.s.Fini()
	os.Exit(code)
}