# Hikmah

A Material 3 Expressive desktop for Hyprland. Rebuilt from scratch: **Go + QML + CSS + JS + Bash**.

**Status:** M00 + M01 — boots, one bar, one clock.

## Layout

| Path        | Package  | Role                                          |
| ----------- | -------- | --------------------------------------------- |
| `asas/`     | **Asas** | Go core: config, IPC, Hyprland, theme service |
| `ufuq/`     | **Ufuq** | QML bar + clock (Quickshell)                  |
| `hikmahui/` | HikmahUI | shared QML kit                                |
| `theme/`    | —        | theme token JSON                              |
| `config/`   | —        | default TOML                                  |

## Run

```bash
go run ./asas/cmd/asasd &     # core daemon
qs -p ufuq/                    # shell surface
```
