import QtQuick
import HikmahUI
import Ufuq

// Ufuq: one strip, islands shape (M01). Notches/full arrive in Phase 4.
Item {
    id: root

    Row {
        anchors.right: parent.right
        anchors.rightMargin: 12
        anchors.verticalCenter: parent.verticalCenter
        spacing: 8

        Clock {}
    }
}
