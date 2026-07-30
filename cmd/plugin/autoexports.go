package main

//#include "autoexports.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
	exports "menus/internal/exports"
)

var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

//export __CreateMenu
func __CreateMenu(slot int32) {
	exports.CreateMenu(slot)
}

//export __SetTitle
func __SetTitle(slot int32, title *C.String) {
	exports.SetTitle(slot, plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(title))))
}

//export __SetMaxVisibleItems
func __SetMaxVisibleItems(slot int32, count int32) {
	exports.SetMaxVisibleItems(slot, count)
}

//export __AddButton
func __AddButton(slot int32, back *C.String, text *C.String) {
	exports.AddButton(slot, plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(back))), plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(text))))
}

//export __AddSubmenu
func __AddSubmenu(slot int32, back *C.String, text *C.String) {
	exports.AddSubmenu(slot, plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(back))), plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(text))))
}

//export __SetAutoCloseTime
func __SetAutoCloseTime(slot int32, seconds float64) {
	exports.SetAutoCloseTime(slot, seconds)
}

//export __OpenMenu
func __OpenMenu(slot int32, callback unsafe.Pointer) {
	exports.OpenMenu(slot, plugify.GetDelegateForFunctionPointer(callback, reflect.TypeOf(exports.MenuCallback(nil))).(exports.MenuCallback))
}

//export __CloseMenu
func __CloseMenu(slot int32) {
	exports.CloseMenu(slot)
}

//export __IsPlayerMenuOpen
func __IsPlayerMenuOpen(slot int32) bool {
	__result := exports.IsPlayerMenuOpen(slot)
	return __result
}
