import React from "react"
import styled from "styled-components"

import { AlbumLink, ArtistLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { useTrackContext } from "Contexts/TrackContext"
import { usePlaybackContext } from "Contexts/PlaybackContext"

const Bar = styled.div`
  display: flex;
  flex-direction: column;
  background: #333;
  padding: 10px;
`

export default function PlayerBar() {
    const { track } = useTrackContext()

    if (!track)
        return <Bar>Nothing is currently playing...</Bar>

    return <Bar>
        <div style={{ display: "flex", marginBottom: 10 }}>
            <AlbumLink {...track.album}>
                <AlbumImage album={track.album} size={72} />
            </AlbumLink>
            <div style={{ marginLeft: 10, overflow: "hidden" }}>
                <h1>{track.title}</h1>
                <h2 style={{ whiteSpace: "nowrap", textOverflow: "ellipsis" }}>
                    {track.artists.map((artist, i, arr) => <React.Fragment key={artist.id}>
                        <ArtistLink {...artist} >
                            {artist.name}
                        </ArtistLink>
                        {i !== arr.length - 1 && ", "}
                    </React.Fragment>)}
                    {" - "}
                    <AlbumLink {...track.album}>
                        {track.album.name}
                    </AlbumLink>
                </h2>
            </div>
        </div>
        <Controls />
    </Bar>
}

function Controls() {
    const { play, pause, playbackState: state } = usePlaybackContext()

    return <div>
        <button onClick={state === "paused" ? play : pause}>{state === "paused" ? "play" : "pause"}</button>
        <ProgressBar />
    </div>
}

function ProgressBar() {
    const { currentTime, duration } = usePlaybackContext()

    return <>{secondsToTimeFormat(currentTime)}/{secondsToTimeFormat(duration)}</>
}