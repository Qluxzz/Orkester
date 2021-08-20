import React, { useState } from "react"

import { IAlbum } from "types/track"
import AlbumImage from "utilities/AlbumImage"
import AlbumArtSizeToggle from "./AlbumArtSizeToggle"

interface IProps {
    album: IAlbum
    onClick: () => void
    size: number | string | "auto"
    icon: React.ReactText
}


function AlbumImageWithToggle({ album, onClick, size, icon }: IProps) {
    const [showToggle, setShowToggle] = useState(false)

    return <div
        onMouseEnter={() => setShowToggle(true)}
        onMouseLeave={() => setShowToggle(false)}
        style={{ position: "relative" }}
    >
        <AlbumImage album={album} size={size} />
        {showToggle &&
            <AlbumArtSizeToggle onClick={onClick} icon={icon} />
        }
    </div>
}

export default AlbumImageWithToggle