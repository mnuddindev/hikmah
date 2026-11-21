// Package ipc exposes Asas services on unix sockets under
// $XDG_RUNTIME_DIR/hikmah/.
package ipc

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"

	"github.com/mnuddindev/hikmah/asas/internal/theme"
)

type Socket struct {
	ln   net.Listener
	path string
}

func (s *Socket) Path() string { return s.path }
func (s *Socket) Close() error { return s.ln.Close() }

func socketPath(name string) (string, error) {
	dir := os.Getenv("XDG_RUNTIME_DIR")
	if dir == "" {
		dir = os.TempDir()
	}
	base := filepath.Join(dir, "hikmah")
	if err := os.MkdirAll(base, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(base, name), nil
}

// ListenTheme serves the current theme as a JSON document on connect.
func ListenTheme(th *theme.Theme) (*Socket, error) {
	p, err := socketPath("themed.sock")
	if err != nil {
		return nil, err
	}
	_ = os.Remove(p)
	ln, err := net.Listen("unix", p)
	if err != nil {
		return nil, err
	}
	s := &Socket{ln: ln, path: p}
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				_ = json.NewEncoder(c).Encode(th.Snapshot())
			}(c)
		}
	}()
	return s, nil
}