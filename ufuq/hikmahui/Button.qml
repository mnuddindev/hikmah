import QtQuick
import "."

Rectangle {
    id: root
    property string label: "Button"
    property bool filled: true
    signal clicked

    implicitWidth: labelItem.implicitWidth + 32
    implicitHeight: 36
    radius: height / 2
    color: filled ? Theme.primary : "transparent"
    border.color: filled ? "transparent" : Theme.outline
    border.width: filled ? 0 : 1

    LText {
        id: labelItem
        anchors.centerIn: parent
        text: root.label
        color: root.filled ? Theme.textOnPrimary : Theme.textOnSurface
        font.pixelSize: 13
        font.weight: Font.Medium
    }

    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: root.clicked()
    }
}
