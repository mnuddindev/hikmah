import QtQuick
import "."

Item {
    id: root
    property string name: "circle"
    property color color: Theme.textOnSurface
    property int size: 20

    implicitWidth: size
    implicitHeight: size

    Rectangle {
        anchors.centerIn: parent
        width: root.size * 0.55
        height: width
        radius: width / 2
        color: root.color
    }
}
