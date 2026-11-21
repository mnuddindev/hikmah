import QtQuick
import HikmahUI

// Clock module: a pill. Format read from Hikmah config later; for M01 hard-coded.
Rectangle {
    id: root

    implicitWidth: timeText.implicitWidth + 28
    implicitHeight: 32
    radius: height / 2
    color: Theme.surfaceContainerHigh

    property var now: new Date()
    property string timeText: Qt.formatTime(now, "HH:mm")
    property string dateText: Qt.formatDate(now, "ddd dd MMM")

    Timer {
        interval: 1000
        running: true
        repeat: true
        onTriggered: root.now = new Date()
    }

    Row {
        anchors.centerIn: parent
        spacing: 8

        LText {
            id: timeText
            text: root.timeText
            color: Theme.onSurface
            font.pixelSize: 13
            font.weight: Font.Medium
        }
        LText {
            text: root.dateText
            color: Theme.onSurfaceVariant
            font.pixelSize: 12
        }
    }
}
