package model

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luevano/mangal/log"
	"github.com/luevano/mangal/tui/base"
	"github.com/luevano/mangal/tui/state/errorstate"
	"github.com/luevano/mangal/tui/state/viewport"
	"github.com/pkg/errors"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.resize(base.Size{
			Width:  msg.Width,
			Height: msg.Height,
		})

		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.Back) && m.state.Backable():
			return m, m.back()
		case key.Matches(msg, m.keyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
			m.resize(m.size)
			return m, nil
		case key.Matches(msg, m.keyMap.Log):
			return m, m.pushState(viewport.New("Logs", log.Aggregate.String(), m.StateSize()))
		}
	case base.MsgBack:
		// this msg can override Backable() output
		return m, m.back()
	case base.State:
		return m, m.pushState(msg)
	case error:
		if errors.Is(msg, context.Canceled) || strings.Contains(msg.Error(), context.Canceled.Error()) {
			return m, nil
		}

		log.L.Err(msg).Msg("")

		return m, m.pushState(errorstate.New(msg))
	}

	cmd := m.state.Update(m, msg)
	return m, cmd
}
