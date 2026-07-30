package entry

import (
	"errors"
	"fmt"
	"path/filepath"

	"menus/internal/menus"
	"menus/pkg/translations"

	polyhook "github.com/fr0nch/go-plugify-polyhook/v2"
	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	"github.com/untrustedmodders/go-plugify"
)

func (p *Plugin) start() error {
	t, err := translations.NewTranslator("translations")
	if err != nil {
		return fmt.Errorf("Error starting plugin with load translations: %v", err)
	}
	menus.T = t

	gamedataPath := filepath.Join(plugify.DataDir(), "gamedata", "menus.json")
	p.gameDataHandle = s2sdk.LoadGameConfigFile([]string{gamedataPath})
	if p.gameDataHandle == 0 || p.gameDataHandle == ^uint32(0) {
		return fmt.Errorf("Error starting plugin with load gamedata: %v", err)
	}

	sig := s2sdk.GetGameConfigSignature(p.gameDataHandle, "CCSPlayerController::ProcessUsercmds")
	if sig == 0 {
		return errors.New("Signature not found")
	}

	args := []polyhook.DataType{
		polyhook.DataType_Pointer,
		polyhook.DataType_Pointer,
	}

	p.hookHandle = polyhook.HookDetour(sig, polyhook.DataType_Int32, args, -1, "")
	if p.hookHandle == 0 {
		return errors.New("Error install hook")
	}

	polyhook.AddCallback(p.hookHandle, polyhook.CallbackType_Pre, menus.OnProcessUsercmds)
	s2sdk.OnClientDisconnect_Register(s2sdk.OnClientDisconnectCallback(p.handleDisconnect))

	return nil
}

func (p *Plugin) stop() error {
	if p.hookHandle != 0 {
		addr := polyhook.GetOriginalAddr(p.hookHandle)
		if addr != 0 {
			polyhook.UnhookDetour(addr)
		}
	}

	if p.gameDataHandle != 0 && p.gameDataHandle != ^uint32(0) {
		s2sdk.CloseGameConfigFile(p.gameDataHandle)
	}

	s2sdk.OnClientDisconnect_Unregister(s2sdk.OnClientDisconnectCallback(p.handleDisconnect))
	return nil
}

func (p *Plugin) handleDisconnect(slot int32) {
	if slot < 0 || slot >= 64 {
		return
	}
	menus.CloseMenu(slot)
}
