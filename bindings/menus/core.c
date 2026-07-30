#include "shared.h"

PLUGIFY_EXPORT void (*__menus_AddButton)(int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT void (*__menus_AddSubmenu)(int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT void (*__menus_CloseMenu)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__menus_CreateMenu)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__menus_IsPlayerMenuOpen)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__menus_OpenMenu)(int32_t, void*) = NULL;


PLUGIFY_EXPORT void (*__menus_SetAutoCloseTime)(int32_t, double) = NULL;


PLUGIFY_EXPORT void (*__menus_SetMaxVisibleItems)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT void (*__menus_SetTitle)(int32_t, String*) = NULL;


