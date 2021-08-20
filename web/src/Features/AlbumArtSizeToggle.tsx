import React from "react"

function AlbumArtSizeToggle({ onClick, icon }: { onClick: () => void, icon: React.ReactText }) {
    return <div
        style={{
            position: "absolute",
            top: 0,
            right: 0,
            width: 25,
            height: 25,
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            cursor: "pointer"
        }}
        onClick={e => {
            e.preventDefault()
            onClick()
        }}
    >{icon}</div>
}

export default AlbumArtSizeToggle