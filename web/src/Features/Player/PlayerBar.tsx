import React from "react"
import styled from "styled-components"

import { AlbumLink } from "utilities/Links"
import ArtistList from "utilities/ArtistList"
import textEllipsisMixin from "utilities/ellipsisText"

import Controls from "./Controls"

import { useTrackContext } from "Context/TrackContext"
import AlbumImageWithToggle from "Features/AlbumImageWithToggle"

const Bar = styled.div`
  display: flex;
  flex-direction: column;
  background: #333;
  padding: 10px;
`


const TrackTitle = styled.h1`
    ${_ => textEllipsisMixin}
`

const ArtistAndAlbum = styled.h2`
    ${_ => textEllipsisMixin}
`

interface IProps {
    showAlbumCover: boolean
    hideAlbumCover: () => void
}

export default function PlayerBar({ showAlbumCover, hideAlbumCover }: IProps) {
    const albumArtSize = 72
    const track = useTrackContext()

    if (!track)
        return <Bar>Nothing is currently playing...</Bar>

    return <Bar>
        <div
            style={{
                display: "flex",
                marginBottom: 10,
                transition: "all 250ms",
                transform: showAlbumCover ? "translateX(0)" : `translateX(-${albumArtSize + 10}px)`
            }}
        >
            <AlbumLink {...track.album}>
                <AlbumImageWithToggle
                    album={track.album}
                    onClick={hideAlbumCover}
                    size={albumArtSize}
                    icon="⬆️"
                />
            </AlbumLink>
            <div style={{ marginLeft: 10, overflow: "hidden" }}>
                <TrackTitle>{track.title}</TrackTitle>
                <ArtistAndAlbum>
                    <ArtistList artists={track.artists} />
                    {" - "}
                    <AlbumLink {...track.album}>
                        {track.album.name}
                    </AlbumLink>
                </ArtistAndAlbum>
            </div>
        </div>
        <Controls />
    </Bar>
}