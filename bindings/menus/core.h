#pragma once

#include "shared.h"

extern void (*__menus_AddButton)(int32_t, String*, String*);

static void AddButton(int32_t slot, String* back, String* text) {
	__menus_AddButton(slot, back, text);
}

extern void (*__menus_AddSubmenu)(int32_t, String*, String*);

static void AddSubmenu(int32_t slot, String* back, String* text) {
	__menus_AddSubmenu(slot, back, text);
}

extern void (*__menus_CloseMenu)(int32_t);

static void CloseMenu(int32_t slot) {
	__menus_CloseMenu(slot);
}

extern void (*__menus_CreateMenu)(int32_t);

static void CreateMenu(int32_t slot) {
	__menus_CreateMenu(slot);
}

extern bool (*__menus_IsPlayerMenuOpen)(int32_t);

static bool IsPlayerMenuOpen(int32_t slot) {
	return __menus_IsPlayerMenuOpen(slot);
}

extern void (*__menus_OpenMenu)(int32_t, void*);

static void OpenMenu(int32_t slot, void* callback) {
	__menus_OpenMenu(slot, callback);
}

extern void (*__menus_SetAutoCloseTime)(int32_t, double);

static void SetAutoCloseTime(int32_t slot, double seconds) {
	__menus_SetAutoCloseTime(slot, seconds);
}

extern void (*__menus_SetMaxVisibleItems)(int32_t, int32_t);

static void SetMaxVisibleItems(int32_t slot, int32_t count) {
	__menus_SetMaxVisibleItems(slot, count);
}

extern void (*__menus_SetTitle)(int32_t, String*);

static void SetTitle(int32_t slot, String* title) {
	__menus_SetTitle(slot, title);
}

