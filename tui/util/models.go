package util

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	"github.com/luevano/mangal/theme/color"
)

func NewList[T any](
	delegateHeight int,
	singular, plural string,
	items []T,
	transform func(T) list.DefaultItem,
) list.Model {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = transform(item)
	}

	border := lipgloss.ThickBorder()
	delegate := list.NewDefaultDelegate()

	// TODO: possibly use the current "window" (where the list is being displayed) accent color,
	// instead of always hardcoding color.Accent
	//
	// Styles don't use mangal/theme/style, as they are more specialized with paddings and whatnot
	styles := delegate.Styles
	styles.NormalTitle = styles.NormalTitle.Bold(true)
	styles.SelectedTitle = styles.SelectedTitle.Bold(true).
		Foreground(color.Accent).
		Border(border, false, false, false, true).
		BorderForeground(color.Accent)
	styles.SelectedDesc = styles.SelectedDesc.
		Foreground(delegate.Styles.NormalDesc.GetForeground()).
		Border(border, false, false, false, true).
		BorderForeground(color.Accent)
	delegate.Styles = styles

	if delegateHeight == 1 {
		delegate.ShowDescription = false
	}

	delegate.SetHeight(delegateHeight)

	l := list.New(listItems, delegate, 0, 0)
	l.SetShowHelp(false)
	l.SetShowFilter(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowPagination(false)
	l.InfiniteScrolling = true
	l.KeyMap.CancelWhileFiltering = Bind("cancel", "esc")

	l.Paginator.Type = paginator.Arabic

	l.SetStatusBarItemName(singular, plural)

	return l
}
