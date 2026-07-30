package exports

import (
	"menus/internal/menus"
)

type MenuCallback func(action string, playerSlot int32)

//plugify:export CreateMenu
func CreateMenu(slot int32) {
	menus.CreateMenu(slot)
}

//plugify:export SetTitle
func SetTitle(slot int32, title string) {
	menus.SetTitle(slot, title)
}

//plugify:export SetMaxVisibleItems
func SetMaxVisibleItems(slot int32, count int32) {
	menus.SetMaxVisibleItems(slot, count)
}

//plugify:export AddButton
func AddButton(slot int32, back string, text string) {
	menus.AddButton(slot, back, text)
}

//plugify:export AddSubmenu
func AddSubmenu(slot int32, back string, text string) {
	menus.AddSubmenu(slot, back, text)
}

//plugify:export SetAutoCloseTime
func SetAutoCloseTime(slot int32, seconds float64) {
	menus.SetAutoCloseTime(slot, seconds)
}

//plugify:export OpenMenu
func OpenMenu(slot int32, callback MenuCallback) {
	menus.OpenMenu(slot, func(action string, slot int32) {
		callback(action, slot)
	})
}

//plugify:export CloseMenu
func CloseMenu(slot int32) {
	menus.CloseMenu(slot)
}

//plugify:export IsPlayerMenuOpen
func IsPlayerMenuOpen(slot int32) bool {
	return menus.IsPlayerMenuOpen(slot)
}
