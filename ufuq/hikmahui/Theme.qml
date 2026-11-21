pragma Singleton
import QtQuick

QtObject {
    property color surface: "#111318"
    property color surfaceContainer: "#1d2024"
    property color surfaceContainerHigh: "#282a2f"
    property color textOnSurface: "#e2e2e9"
    property color textOnSurfaceVariant: "#c4c6d0"
    property color primary: "#a8c7fa"
    property color textOnPrimary: "#00315f"
    property color outline: "#8e9099"

    property int cornerSm: 8
    property int cornerMd: 12
    property int cornerLg: 16
    property int cornerFull: 9999

    property real springStiffness: 380
    property real springDamping: 28
}
