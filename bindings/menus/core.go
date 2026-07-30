package menus

/*
#include "core.h"
#cgo noescape AddButton
#cgo noescape AddSubmenu
#cgo noescape CloseMenu
#cgo noescape CreateMenu
#cgo noescape IsPlayerMenuOpen
#cgo noescape OpenMenu
#cgo noescape SetAutoCloseTime
#cgo noescape SetMaxVisibleItems
#cgo noescape SetTitle
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from menus (group: core)

var _AddButton = func(slot int32, back string, text string) {
	__slot := C.int32_t(slot)
	__back := plugify.ConstructString(back)
	__text := plugify.ConstructString(text)
	plugify.Block {
		Try: func() {
			C.AddButton(__slot, (*C.String)(unsafe.Pointer(&__back)), (*C.String)(unsafe.Pointer(&__text)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__back)
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// AddButton 
func AddButton(slot int32, back string, text string) {
	_AddButton(slot, back, text)
}

var _AddSubmenu = func(slot int32, back string, text string) {
	__slot := C.int32_t(slot)
	__back := plugify.ConstructString(back)
	__text := plugify.ConstructString(text)
	plugify.Block {
		Try: func() {
			C.AddSubmenu(__slot, (*C.String)(unsafe.Pointer(&__back)), (*C.String)(unsafe.Pointer(&__text)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__back)
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// AddSubmenu 
func AddSubmenu(slot int32, back string, text string) {
	_AddSubmenu(slot, back, text)
}

var _CloseMenu = func(slot int32) {
	__slot := C.int32_t(slot)
	C.CloseMenu(__slot)
}

// CloseMenu 
func CloseMenu(slot int32) {
	_CloseMenu(slot)
}

var _CreateMenu = func(slot int32) {
	__slot := C.int32_t(slot)
	C.CreateMenu(__slot)
}

// CreateMenu 
func CreateMenu(slot int32) {
	_CreateMenu(slot)
}

var _IsPlayerMenuOpen = func(slot int32) bool {
	var __retVal bool
	__slot := C.int32_t(slot)
	__retVal = bool(C.IsPlayerMenuOpen(__slot))
	return __retVal
}

// IsPlayerMenuOpen 
func IsPlayerMenuOpen(slot int32) bool {
	return _IsPlayerMenuOpen(slot)
}

var _OpenMenu = func(slot int32, callback MenuCallback) {
	__slot := C.int32_t(slot)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OpenMenu(__slot, __callback)
}

// OpenMenu 
func OpenMenu(slot int32, callback MenuCallback) {
	_OpenMenu(slot, callback)
}

var _SetAutoCloseTime = func(slot int32, seconds float64) {
	__slot := C.int32_t(slot)
	__seconds := C.double(seconds)
	C.SetAutoCloseTime(__slot, __seconds)
}

// SetAutoCloseTime 
func SetAutoCloseTime(slot int32, seconds float64) {
	_SetAutoCloseTime(slot, seconds)
}

var _SetMaxVisibleItems = func(slot int32, count int32) {
	__slot := C.int32_t(slot)
	__count := C.int32_t(count)
	C.SetMaxVisibleItems(__slot, __count)
}

// SetMaxVisibleItems 
func SetMaxVisibleItems(slot int32, count int32) {
	_SetMaxVisibleItems(slot, count)
}

var _SetTitle = func(slot int32, title string) {
	__slot := C.int32_t(slot)
	__title := plugify.ConstructString(title)
	plugify.Block {
		Try: func() {
			C.SetTitle(__slot, (*C.String)(unsafe.Pointer(&__title)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__title)
		},
	}.Do()
}

// SetTitle 
func SetTitle(slot int32, title string) {
	_SetTitle(slot, title)
}

