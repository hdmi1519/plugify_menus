package menus

import (
	"menus/internal/constants"

	polyhook "github.com/fr0nch/go-plugify-polyhook/v2"
	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

var sessions [64]Session

func getSession(slot int32) *Session {
	if slot < 0 || slot >= 64 {
		return nil
	}
	return &sessions[slot]
}

func OnProcessUsercmds(cb polyhook.HookHandle, paramsHandle polyhook.ParametersHandle, count int32, retHandle polyhook.ReturnHandle, cbType polyhook.CallbackType) polyhook.ResultType {
	params := polyhook.NewParameters(paramsHandle)

	controllerPtr, err := params.GetPointer(0)
	if err != nil || controllerPtr == 0 {
		return 1
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				slot := s2sdk.EntPointerToPlayerSlot(controllerPtr)
				if slot >= 0 && slot < 64 {
					s := &sessions[slot]
					s.IsActive = false
					s.Menu = nil
					s.IsLocked = false
					s2sdk.PrintCentreHtml(slot, " ", 0)
				}
			}
		}()

		slot := s2sdk.EntPointerToPlayerSlot(controllerPtr)
		if slot < 0 || slot >= 64 || !s2sdk.IsClientInGame(slot) {
			return
		}

		s := &sessions[slot]
		if !s.IsActive || s.IsLocked || s.Menu == nil {
			return
		}

		now := s2sdk.GetEngineTime()

		if s.Menu.AutoCloseAfter > 0 && now-s.Menu.OpenedAt >= s.Menu.AutoCloseAfter {
			closeInternal(slot, s)
			return
		}

		controller := s2sdk.PlayerSlotToEntHandle(slot)
		if controller != -1 {
			pawn := s2sdk.GetEntSchemaEnt(controller, "CCSPlayerController", "m_hPlayerPawn", 0)
			if pawn != -1 {
				s2sdk.SetEntSchemaFloat(pawn, "CCSPlayerPawn", "m_flVelocityModifier", 0.0, true, 0)
			}
		}

		if now-s.LastRenderTime >= 0.12 {
			Render(slot, s)
		}

		current := s2sdk.GetClientButtons(slot, 0) & constants.MenuButtons
		old := s.LastButtons & constants.MenuButtons
		s.LastButtons = current

		pressed := current & ^old
		if pressed == 0 {
			return
		}

		if now-s.LastClickTime < 0.20 {
			return
		}

		handleInput(slot, s, pressed)
	}()

	return 1
}

func handleInput(slot int32, s *Session, pressed uint64) {
	menu := s.Menu
	now := s2sdk.GetEngineTime()

	menu.OpenedAt = now

	total := len(menu.Options)
	if total == 0 {
		return
	}

	limit := int(menu.MaxVisibleItems)
	start := menu.CurrentPage * limit
	end := start + limit
	if end > total {
		end = total
	}
	itemsOnPage := end - start

	switch {
	case pressed&constants.IN_FORWARD != 0:
		menu.CurrentOption--
		if menu.CurrentOption < 0 {
			menu.CurrentOption = itemsOnPage - 1
		}
		s.LastClickTime = now
		Render(slot, s)

	case pressed&constants.IN_BACK != 0:
		menu.CurrentOption++
		if menu.CurrentOption >= itemsOnPage {
			menu.CurrentOption = 0
		}
		s.LastClickTime = now
		Render(slot, s)

	case pressed&constants.IN_MOVELEFT != 0:
		if menu.CurrentPage > 0 {
			menu.CurrentPage--
			menu.CurrentOption = 0
			s.LastClickTime = now
			Render(slot, s)
		} else if menu.HasBack && len(s.History) > 1 {
			s.IsLocked = true
			s.History = s.History[:len(s.History)-1]
			s.Menu = s.History[len(s.History)-1]
			s.IsLocked = false
			s.LastClickTime = now
			Render(slot, s)
		}

	case pressed&constants.IN_MOVERIGHT != 0:
		maxPages := total / limit
		if total%limit == 0 {
			maxPages--
		}
		if menu.CurrentPage < maxPages {
			menu.CurrentPage++
			menu.CurrentOption = 0
			s.LastClickTime = now
			Render(slot, s)
		}

	case pressed&constants.IN_USE != 0:
		idx := start + menu.CurrentOption
		if idx >= total {
			return
		}
		opt := menu.Options[idx]

		s.IsLocked = true
		s.LastClickTime = now

		if menu.Callback != nil {
			menu.Callback(opt.Back, slot)
		}

		if opt.Back == "close" || opt.Back == "back" || opt.Back == "back_from_nested" {
			closeInternal(slot, s)
		}

		s.IsLocked = false

	case pressed&constants.IN_RELOAD != 0:
		closeInternal(slot, s)

	case pressed&constants.IN_SCORE != 0:
		closeInternal(slot, s)
	}
}

func closeInternal(slot int32, s *Session) {
	if s == nil {
		return
	}

	s.IsLocked = true

	controller := s2sdk.PlayerSlotToEntHandle(slot)
	if controller != -1 {
		pawn := s2sdk.GetEntSchemaEnt(controller, "CCSPlayerController", "m_hPlayerPawn", 0)
		if pawn != -1 {
			s2sdk.SetEntSchemaFloat(pawn, "CCSPlayerPawn", "m_flVelocityModifier", 1.0, true, 0)
		}
	}

	s.IsActive = false
	s.Menu = nil
	s.History = nil
	s.LastButtons = 0
	s.IsLocked = false
	builderCache[slot] = nil

	s2sdk.PrintCentreHtml(slot, " ", 0)
}
