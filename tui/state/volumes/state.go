package volumes

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luevano/libmangal"
	"github.com/luevano/libmangal/mangadata"
	"github.com/luevano/mangal/tui/base"
	"github.com/luevano/mangal/tui/state/chapters"
	"github.com/luevano/mangal/tui/state/listwrapper"
	"github.com/luevano/mangal/tui/state/loading"
)

var _ base.State = (*State)(nil)

type State struct {
	client  *libmangal.Client
	manga   *mangadata.Manga
	volumes []*mangadata.Volume
	list    *listwrapper.State
	keyMap  KeyMap
}

func (s *State) Intermediate() bool {
	return false
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Title() base.Title {
	return base.Title{Text: (*s.manga).String()}
}

func (s *State) Subtitle() string {
	return s.list.Subtitle()
}

func (s *State) Status() string {
	return s.list.Status()
}

func (s *State) Backable() bool {
	return s.list.FilterState() == list.Unfiltered
}

func (s *State) Resize(size base.Size) {
	s.list.Resize(size)
}

func (s *State) Update(model base.Model, msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s.list.FilterState() == list.Filtering {
			goto end
		}

		item, ok := s.list.SelectedItem().(Item)
		if !ok {
			return nil
		}

		switch {
		case key.Matches(msg, s.keyMap.Confirm):
			return tea.Sequence(
				func() tea.Msg {
					return loading.New("Searching", fmt.Sprintf("Getting chapters for volume %s", *item.volume))
				},
				func() tea.Msg {
					cL, err := s.client.VolumeChapters(model.Context(), *item.volume)
					if err != nil {
						return err
					}

					var chapterList []*mangadata.Chapter
					for _, c := range cL {
						chapterList = append(chapterList, &c)
					}

					return chapters.New(s.client, s.manga, item.volume, chapterList)
				},
			)
		}
	}
end:
	return s.list.Update(model, msg)
}

func (s *State) View(model base.Model) string {
	return s.list.View(model)
}

func (s *State) Init(model base.Model) tea.Cmd {
	return s.list.Init(model)
}
