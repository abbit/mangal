package providers

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/luevano/libmangal"
)

var _ list.Item = (*Item)(nil)

type Item struct {
	libmangal.ProviderLoader
}

func (i Item) FilterValue() string {
	return i.String()
}

func (i Item) Title() string {
	return i.FilterValue()
}

func (i Item) Description() string {
	info := i.Info()
	return fmt.Sprintf(
		"%s v%s\n%s",
		info.ID,
		info.Version,
		info.Website,
	)
}
