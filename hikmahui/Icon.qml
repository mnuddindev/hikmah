import QtQuick
import HikmahUI

// Minimal icon: draws a Material Symbols glyph from a font if present,
// otherwise falls back to a filled circle. Full symbol set lands in Phase 2.
Item {
    id: root
    property string name: "circle"
    property color color: Theme.onSurface
    property int size: 20

    implicitWidth: size
    implicitHeight: size

    Text {
        anchors.centerIn: parent
        text: "" // filled in Phase 2 with SymbolPaths.js
        visible: false
    }
    Rectangle {
        anchors.centerIn: parent
        width: root.size * 0.55
        height: width
        radius: width / 2
        color: root.color
    }
}
