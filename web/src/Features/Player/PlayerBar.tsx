import React, { useEffect, useState } from "react"
import styled from "styled-components"

import { usePlayerContext } from "Context"
import { AlbumLink, ArtistLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"

const Bar = styled.div`
  display: flex;
  flex-direction: column;
  background: #333;
  padding: 10px;
`

export default function PlayerBar() {
    const { track } = usePlayerContext()

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
                    {track.artists.map((artist, i, arr) => <>
                        <ArtistLink {...artist} key={artist.id}>
                            {artist.name}
                        </ArtistLink>
                        {i !== arr.length - 1 && ", "}
                    </>)}
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
    const { togglePlayback, state } = usePlayerContext()

    return <div>
        <button onClick={togglePlayback}>{state === "paused" ? "play" : "pause"}</button>
        <ProgressBar />
    </div>
}

function ProgressBar() {
    const [data, setData] = useState<{ duration: number, timestamp: number }>()
    const { player } = usePlayerContext()

    function updateCurrentTimeInterval() {
        const interval = setInterval(() => {
            if (!player)
                return

            setData({
                duration: Math.round(player.duration),
                timestamp: Math.round(player.currentTime)
            })
        }, 1000)

        return () => {
            clearInterval(interval)
        }
    }

    useEffect(updateCurrentTimeInterval, [player])

    if (!data || !data.timestamp || !data.duration)
        return null

    return <>{secondsToTimeFormat(data.timestamp)}/{secondsToTimeFormat(data.duration)}</>
}