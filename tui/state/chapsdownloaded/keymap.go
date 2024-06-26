package chapsdownloaded

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

type KeyMap struct {
	Quit,
	Open,
	Retry key.Binding

	state *State
}

func (k KeyMap) ShortHelp() []key.Binding {
	bindings := []key.Binding{
		k.Quit,
		k.Open,
	}

	if len(k.state.options.Failed) > 0 {
		bindings = append(bindings, k.Retry)
	}

	return bindings
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
