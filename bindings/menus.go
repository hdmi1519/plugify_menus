//go:build plugin
// +build plugin
package main

//TODO: replace "__package__" by your package name
import (
	"unsafe"
	"__package__/menus"
	"github.com/untrustedmodders/go-plugify"
)

var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

//go:linkname menus_AddButton __package__/menus._AddButton
var menus_AddButton func(slot int32, back string, text string)
var Menus_AddButton = &menus_AddButton

//go:linkname menus_AddSubmenu __package__/menus._AddSubmenu
var menus_AddSubmenu func(slot int32, back string, text string)
var Menus_AddSubmenu = &menus_AddSubmenu

//go:linkname menus_CloseMenu __package__/menus._CloseMenu
var menus_CloseMenu func(slot int32)
var Menus_CloseMenu = &menus_CloseMenu

//go:linkname menus_CreateMenu __package__/menus._CreateMenu
var menus_CreateMenu func(slot int32)
var Menus_CreateMenu = &menus_CreateMenu

//go:linkname menus_IsPlayerMenuOpen __package__/menus._IsPlayerMenuOpen
var menus_IsPlayerMenuOpen func(slot int32) bool
var Menus_IsPlayerMenuOpen = &menus_IsPlayerMenuOpen

//go:linkname menus_OpenMenu __package__/menus._OpenMenu
var menus_OpenMenu func(slot int32, callback menus.MenuCallback)
var Menus_OpenMenu = &menus_OpenMenu

//go:linkname menus_SetAutoCloseTime __package__/menus._SetAutoCloseTime
var menus_SetAutoCloseTime func(slot int32, seconds float64)
var Menus_SetAutoCloseTime = &menus_SetAutoCloseTime

//go:linkname menus_SetMaxVisibleItems __package__/menus._SetMaxVisibleItems
var menus_SetMaxVisibleItems func(slot int32, count int32)
var Menus_SetMaxVisibleItems = &menus_SetMaxVisibleItems

//go:linkname menus_SetTitle __package__/menus._SetTitle
var menus_SetTitle func(slot int32, title string)
var Menus_SetTitle = &menus_SetTitle

/*func init() {
	menus.ModuleName = "__package__"
}*/
