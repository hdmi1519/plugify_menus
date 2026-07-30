package menus

import (
	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

var builderCache [64]*Instance

func CreateMenu(slot int32) {
	if slot < 0 || slot >= 64 {
		return
	}

	builderCache[slot] = &Instance{
		Title:           "Menu",
		MaxVisibleItems: 4,
		Options:         make([]Option, 0, 8),
	}
}

func SetTitle(slot int32, title string) {
	if slot < 0 || slot >= 64 || builderCache[slot] == nil {
		return
	}
	builderCache[slot].Title = title
}

func SetMaxVisibleItems(slot int32, count int32) {
	if slot < 0 || slot >= 64 || builderCache[slot] == nil || count <= 0 {
		return
	}
	builderCache[slot].MaxVisibleItems = count
}

func AddButton(slot int32, back, text string) {
	if slot < 0 || slot >= 64 || builderCache[slot] == nil {
		return
	}
	builderCache[slot].Options = append(builderCache[slot].Options, Option{Type: 0, Back: back, Text: text})
}

func AddSubmenu(slot int32, back, text string) {
	if slot < 0 || slot >= 64 || builderCache[slot] == nil {
		return
	}
	builderCache[slot].Options = append(builderCache[slot].Options, Option{Type: 1, Back: back, Text: text})
}

func SetAutoCloseTime(slot int32, seconds float64) {
	if slot < 0 || slot >= 64 || builderCache[slot] == nil || seconds <= 0 {
		return
	}
	builderCache[slot].AutoCloseAfter = seconds
}

func OpenMenu(slot int32, callback func(action string, slot int32)) {
	if slot < 0 || slot >= 64 {
		return
	}

	built := builderCache[slot]
	if built == nil || len(built.Options) == 0 {
		return
	}

	if built.MaxVisibleItems <= 0 {
		built.MaxVisibleItems = 4
	}
	if built.CurrentOption < 0 {
		built.CurrentOption = 0
	}
	if built.CurrentPage < 0 {
		built.CurrentPage = 0
	}

	s := getSession(slot)
	if s == nil {
		return
	}

	s.IsLocked = true

	if s.IsActive && len(s.History) > 0 {
		built.HasBack = true
	} else {
		s.History = make([]*Instance, 0, 4)
		s.IsActive = true
	}

	built.Callback = callback
	built.OpenedAt = s2sdk.GetEngineTime()
	s.Menu = built
	s.History = append(s.History, built)
	s.IsLocked = false

	controller := s2sdk.PlayerSlotToEntHandle(slot)
	if controller != -1 {
		pawn := s2sdk.GetEntSchemaEnt(controller, "CCSPlayerController", "m_hPlayerPawn", 0)
		if pawn != -1 {
			s2sdk.SetEntSchemaFloat(pawn, "CCSPlayerPawn", "m_flVelocityModifier", 0.0, true, 0)
		}
	}

	Render(slot, s)
}

func CloseMenu(slot int32) {
	if slot < 0 || slot >= 64 {
		return
	}
	s := getSession(slot)
	if s == nil {
		return
	}
	closeInternal(slot, s)
}

func IsPlayerMenuOpen(slot int32) bool {
	if slot < 0 || slot >= 64 {
		return false
	}
	s := getSession(slot)
	if s == nil {
		return false
	}
	return s.IsActive
}
