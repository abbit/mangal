package chapters

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/luevano/mangal/tui/state/listwrapper"
)

var _ help.KeyMap = (*KeyMap)(nil)

type KeyMap struct {
	UnselectAll,
	SelectAll,
	ToggleChapterNumber,
	ToggleGroup,
	ToggleDate,
	Toggle,
	Read,
	OpenURL,
	Download,
	Anilist,
	Confirm,
	ChangeFormat key.Binding

	list listwrapper.KeyMap
}

func (k KeyMap) ShortHelp() []key.Binding {
	return append(
		k.list.ShortHelp(),
		k.Toggle,
		k.Read,
		k.Download,
	)
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{k.SelectAll, k.UnselectAll, k.ToggleChapterNumber, k.ToggleGroup, k.ToggleDate},
		{k.Anilist},
		{k.ChangeFormat, k.OpenURL},
	}
}
