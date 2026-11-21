package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/mnuddindev/hikmah/asas/internal/config"
	"github.com/mnuddindev/hikmah/asas/internal/hypr"
	"github.com/mnuddindev/hikmah/asas/internal/ipc"
	"github.com/mnuddindev/hikmah/asas/internal/log"
	"github.com/mnuddindev/hikmah/asas/internal/theme"
)

func main() {
	lg := log.New()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.Load("config/hikmah.toml")
	if err != nil {
		lg.Warn("config missing, using defaults", "err", err)
		cfg = config.Defaults()
	}
	lg.Info("config loaded", "theme", cfg.General.Theme, "bar", cfg.Bar.Position)

	th := theme.New(cfg.General.Theme)
	if err := th.Load("theme/" + cfg.General.Theme + ".json"); err != nil {
		lg.Error("theme load failed", "err", err)
		os.Exit(1)
	}
	lg.Info("theme loaded", "name", th.Name(), "colors", len(th.Colors()))

	// Start the theme service (unix socket).
	tsock, err := ipc.ListenTheme(th)
	if err != nil {
		lg.Error("theme socket", "err", err)
		os.Exit(1)
	}
	defer tsock.Close()
	lg.Info("theme socket up", "path", tsock.Path())

	// Hyprland: connect + subscribe to events (best-effort).
	h, err := hypr.New()
	if err != nil {
		lg.Warn("hypr not reachable; continuing", "err", err)
	} else {
		go h.Subscribe(ctx, func(ev hypr.Event) {
			lg.Debug("hypr event", "name", ev.Name, "data", ev.Data)
		})
		defer h.Close()
	}

	<-ctx.Done()
	lg.Info("shutdown")
}