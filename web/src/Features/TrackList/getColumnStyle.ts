import { IColumn } from "./TrackListBase"

export default function getColumnStyle(column: IColumn): React.CSSProperties {
    const style: React.CSSProperties = {}

    if (typeof column.width === "number") {
        style.flexShrink = 0
        style.width = column.width
    } else if (column.width === "grow") {
        style.flexGrow = 1
    }

    if (column.centered) {
        style.textAlign = "center"
    }

    style.overflow = "hidden"

    return style
}