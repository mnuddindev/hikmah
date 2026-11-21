package theme

import (
	"encoding/json"
	"os"
	"sync"
)

type Theme struct {
	name string
	mu   sync.RWMutex
	data map[string]any
}

func New(name string) *Theme { return &Theme{name: name, data: map[string]any{}} }

func (t *Theme) Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var d map[string]any
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	t.mu.Lock()
	t.data = d
	t.mu.Unlock()
	return nil
}

func (t *Theme) Name() string { return t.name }

func (t *Theme) Colors() map[string]any {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if m, ok := t.data["colors"].(map[string]any); ok {
		return m
	}
	return nil
}

func (t *Theme) Snapshot() map[string]any {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.data
}