package loading

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luevano/mangal/theme/color"
	"github.com/luevano/mangal/theme/style"
	"github.com/luevano/mangal/tui/base"
)

var _ base.State = (*State)(nil)

type State struct {
	message  string
	subtitle string
	spinner  spinner.Model
	keyMap   KeyMap
}

func (s *State) Intermediate() bool {
	return true
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Title() base.Title {
	// return base.Title{Text: s.message, Background: color.Loading}
	return base.Title{Text: "Loading", Background: color.Loading}
}

func (s *State) Subtitle() string {
	return s.subtitle
}

func (s *State) Status() string {
	return s.spinner.View()
	// return ""
}

func (s *State) Backable() bool {
	return true
}

func (s *State) Resize(size base.Size) {
}

func (s *State) SetMessage(message string) {
	s.message = message
}

func (s *State) Update(model base.Model, msg tea.Msg) (cmd tea.Cmd) {
	s.spinner, cmd = s.spinner.Update(msg)
	return cmd
}

func (s *State) View(model base.Model) string {
	return fmt.Sprint(
		style.Bold.Accent.Render(s.spinner.View()),
		style.Normal.Secondary.Render(s.message),
	)
}

func (s *State) Init(model base.Model) tea.Cmd {
	return s.spinner.Tick
}
