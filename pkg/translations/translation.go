package translations

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	"github.com/untrustedmodders/go-plugify"
)

const pluginName = "menus"

type Translator struct {
	phrases map[string]any
}

func NewTranslator(fileName string) (*Translator, error) {
	dirPath := filepath.Join(
		plugify.DataDir(),
		"translations",
		pluginName,
	)

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create translation directories: %w", err)
	}

	pathToFile := filepath.Join(dirPath, fmt.Sprintf("%s.json", fileName))

	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("translation file is missing at: %s", pathToFile)
	}

	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("failed read translation file: %w", err)
	}

	var jsonData map[string]any
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("failed parse translation JSON: %w", err)
	}

	return &Translator{phrases: jsonData}, nil
}

func (t *Translator) Get(playerSlot int32, key string, args ...any) string {
	lang := "en"

	clientLang := s2sdk.GetClientLanguage(playerSlot)
	if len(clientLang) >= 2 {
		lang = clientLang[:2]
	}

	text, found := t.getPhraseText(key, lang)
	if !found {
		return fmt.Sprintf("[Missing translation: %s]", key)
	}

	if len(args) > 0 {
		return fmt.Sprintf(text, args...)
	}

	return text
}

func (t *Translator) getPhraseText(compoundKey string, lang string) (string, bool) {
	keys := strings.Split(compoundKey, ".")

	var current any = t.phrases

	for _, key := range keys {
		if currentMap, ok := current.(map[string]any); ok {
			if next, found := currentMap[key]; found {
				current = next
			} else {
				return "", false
			}
		} else {
			return "", false
		}
	}

	langMap, ok := current.(map[string]any)
	if !ok {
		return "", false
	}

	if text, ok := langMap[lang].(string); ok {
		return text, true
	}

	if text, ok := langMap["en"].(string); ok {
		return text, true
	}

	return "", false
}
