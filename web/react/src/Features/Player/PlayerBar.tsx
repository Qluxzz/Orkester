import React from "react"
import styled from "styled-components"

import { AlbumLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import ArtistList from "utilities/ArtistList"

import textEllipsisMixin from "utilities/ellipsisText"

import Controls from "./Controls"

import { useTrackContext } from "Context/TrackContext"

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

export default function PlayerBar() {
    const track = useTrackContext()

    if (!track)
        return <Bar>Nothing is currently playing...</Bar>

    return <Bar>
        <div style={{ display: "flex", marginBottom: 10 }}>
            <AlbumLink {...track.album}>
                <AlbumImage album={track.album} size={72} />
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