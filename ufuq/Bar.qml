import QtQuick

Item {
    id: root
    implicitWidth: parent ? parent.width : 1920
    implicitHeight: 44

    Row {
        anchors.right: parent.right
        anchors.rightMargin: 12
        anchors.verticalCenter: parent.verticalCenter
        spacing: 8
        Clock {}
    }
}
