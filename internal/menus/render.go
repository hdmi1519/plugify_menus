package menus

import (
	"fmt"
	"strings"
	"sync"

	"menus/pkg/translations"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

var (
	T *translations.Translator

	stringBuilderPool = sync.Pool{
		New: func() any {
			b := new(strings.Builder)
			b.Grow(1024)
			return b
		},
	}
)

func Render(slot int32, s *Session) {
	if s == nil || s.Menu == nil {
		return
	}

	now := s2sdk.GetEngineTime()
	if now-s.LastRenderTime < 0.12 {
		return
	}

	menu := s.Menu

	sb := stringBuilderPool.Get().(*strings.Builder)
	sb.Reset()
	defer stringBuilderPool.Put(sb)

	sep := "<font face='monospace' color='#444444'>─────────────────────────────────────────────</font><br>"

	fmt.Fprintf(sb, "<font face='monospace' class='fontSize-m' color='#ffffff'><b>%s</b></font><br>%s", menu.Title, sep)

	total := len(menu.Options)
	limit := int(menu.MaxVisibleItems)
	start := menu.CurrentPage * limit
	end := start + limit
	if end > total {
		end = total
	}

	for i := start; i < end; i++ {
		selected := (i - start) == menu.CurrentOption
		opt := menu.Options[i]
		if selected {
			fmt.Fprintf(sb, "<font face='monospace' class='fontSize-sm' color='#ffffff'>► <b>%s</b> ◄</font><br>", opt.Text)
		} else {
			fmt.Fprintf(sb, "<font face='monospace' class='fontSize-sm' color='#777777'>&nbsp;&nbsp;&nbsp;&nbsp;%s</font><br>", opt.Text)
		}
	}

	maxPages := total / limit
	if total%limit == 0 {
		maxPages--
	}

	fmt.Fprintf(sb, "%s<font face='monospace' class='fontSize-s'>", sep)

	hasPagination := maxPages > 0

	if hasPagination {
		canGoBack := menu.CurrentPage > 0
		canGoNext := menu.CurrentPage < maxPages
		hasBackButton := menu.HasBack && len(s.History) > 1

		if canGoBack {
			if T != nil {
				sb.WriteString(fmt.Sprintf("<font color='#ffffff'>[A] %s</font>", T.Get(slot, "menus.nav.back")))
			} else {
				sb.WriteString("<font color='#ffffff'>[A] Back</font>")
			}
		} else if hasBackButton {
			if T != nil {
				sb.WriteString(fmt.Sprintf("<font color='#ffffff'>[A] %s</font>", T.Get(slot, "menus.nav.menu_back")))
			} else {
				sb.WriteString("<font color='#ffffff'>[A] Back</font>")
			}
		} else {
			if T != nil {
				sb.WriteString(fmt.Sprintf("<font color='#777777'>[A] %s</font>", T.Get(slot, "menus.nav.back")))
			} else {
				sb.WriteString("<font color='#777777'>[A] Back</font>")
			}
		}

		sb.WriteString(" | ")

		if canGoNext {
			if T != nil {
				sb.WriteString(fmt.Sprintf("<font color='#ffffff'>[D] %s</font>", T.Get(slot, "menus.nav.next")))
			} else {
				sb.WriteString("<font color='#ffffff'>[D] Next</font>")
			}
		} else {
			if T != nil {
				sb.WriteString(fmt.Sprintf("<font color='#555555'>[D] %s</font>", T.Get(slot, "menus.nav.next")))
			} else {
				sb.WriteString("<font color='#777777'>[D] Next</font>")
			}
		}

		sb.WriteString("<br>")
	}

	if T != nil {
		sb.WriteString(T.Get(slot, "menus.footer_controls"))
	} else {
		sb.WriteString("[R] Close</font>")
	}

	s.LastRenderTime = now
	s2sdk.PrintCentreHtml(slot, sb.String(), 60)
}
