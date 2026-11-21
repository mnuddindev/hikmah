pragma Singleton
import QtQuick

// Theme singleton: reads tokens from Asas's themed.sock at startup,
// then exposes them as flat properties. Phase 3 will hot-reload.
QtObject {
    id: t

    // Placeholder defaults (match theme/default.json); replaced on connect.
    property color surface: "#111318"
    property color surfaceContainer: "#1d2024"
    property color surfaceContainerHigh: "#282a2f"
    property color onSurface: "#e2e2e9"
    property color onSurfaceVariant: "#c4c6d0"
    property color primary: "#a8c7fa"
    property color onPrimary: "#00315f"
    property color outline: "#8e9099"

    property int cornerSm: 8
    property int cornerMd: 12
    property int cornerLg: 16
    property int cornerFull: 9999

    property real springStiffness: 380
    property real springDamping: 28

    // Load tokens from the socket. Failures keep the defaults.
    Component.onCompleted: {
        let xhr = new XMLHttpRequest();
        let path = "$XDG_RUNTIME_DIR/hikmah/themed.sock";
        // QML can't read unix sockets directly; Phase 3 adds a Go shim
        // (themed → QML via file cache). For M01 we just keep defaults.
    }
}
