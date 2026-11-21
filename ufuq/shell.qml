import QtQuick
import QtQuick.Window
import Quickshell
import Ufuq

ShellRoot {
    PanelWindow {
        id: barWin
        anchors.top: true
        anchors.left: true
        anchors.right: true
        implicitHeight: 44
        color: "transparent"
        Bar {}
    }
}
