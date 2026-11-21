// Package hypr talks to Hyprland 0.57 over its unix sockets.
package hypr

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type Event struct {
	Name string
	Data string
}

type Client struct {
	sig  string
	cmd  string // .socket.sock  (commands)
	evt  string // .socket2.sock (events)
	conn net.Conn
}

func New() (*Client, error) {
	sig := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if sig == "" {
		return nil, fmt.Errorf("HYPRLAND_INSTANCE_SIGNATURE not set")
	}
	// Prefer XDG_RUNTIME_DIR (0.40+); fall back to /tmp/hypr.
	base := filepath.Join(os.Getenv("XDG_RUNTIME_DIR"), "hypr", sig)
	if _, err := os.Stat(base); err != nil {
		base = filepath.Join("/tmp/hypr", sig)
	}
	c := &Client{
		sig: sig,
		cmd: filepath.Join(base, ".socket.sock"),
		evt: filepath.Join(base, ".socket2.sock"),
	}
	if _, err := os.Stat(c.evt); err != nil {
		return nil, fmt.Errorf("event socket missing at %s", c.evt)
	}
	return c, nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Subscribe reads Hyprland events until ctx is done.
func (c *Client) Subscribe(ctx context.Context, fn func(Event)) {
	conn, err := net.Dial("unix", c.evt)
	if err != nil {
		return
	}
	c.conn = conn
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		line := sc.Text()
		// Format: EVENT>>DATA
		for i := 0; i+1 < len(line); i++ {
			if line[i] == '>' && line[i+1] == '>' {
				fn(Event{Name: line[:i], Data: line[i+2:]})
				break
			}
		}
	}
}