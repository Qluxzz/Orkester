import React, { useState } from "react"
import styled, { css } from "styled-components"

import { AlbumLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { useTrackContext } from "Contexts/TrackContext"
import { usePlaybackContext } from "Contexts/PlaybackContext"
import { useEffect } from "react"
import ArtistList from "utilities/ArtistList"

const Bar = styled.div`
  display: flex;
  flex-direction: column;
  background: #333;
  padding: 10px;
`

const textEllipsis = css`
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
`

const TrackTitle = styled.h1`
    ${_ => textEllipsis}
`

const ArtistAndAlbum = styled.h2`
    ${_ => textEllipsis}
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

function Controls() {
    const { play, pause, playbackState, currentTime, duration } = usePlaybackContext()

    return <div style={{ display: "flex" }}>
        <button
            onClick={playbackState === "paused"
                ? play
                : pause
            }
            style={{ width: 72 }}
        >
            {playbackState === "paused" ? "play" : "pause"}
        </button>
        <Slider duration={duration} currentTime={currentTime} />
    </div>
}

function Slider({ currentTime, duration }: { currentTime: number, duration: number }) {
    const { seek } = usePlaybackContext()
    const [value, setValue] = useState(currentTime)
    const [interacting, setInteracting] = useState(false)

    useEffect(() => {
        setValue(currentTime)
    }, [currentTime])

    return <div style={{ display: "flex", alignItems: "center", flexGrow: 1 }}>
        <div style={{ padding: "0 10px" }}>{secondsToTimeFormat(value)}</div>
        <input
            style={{ width: "100%" }}
            type="range"
            min={0}
            max={duration}
            value={!interacting ? value : undefined}
            defaultValue={interacting ? value : undefined}
            onMouseUp={e => {
                seek(e.currentTarget.valueAsNumber)
                setInteracting(false)
            }}
            onMouseDown={() => setInteracting(true)}
            onMouseMove={e => {
                setValue(e.currentTarget.valueAsNumber)
            }}
        />
        <DurationOrRemainingTime
            duration={duration}
            currentTime={currentTime}
        />
    </div>
}

function DurationOrRemainingTime({ duration, currentTime }: { duration: number, currentTime: number }) {
    const [inversed, setInversed] = useState(false)

    const time = inversed
        ? `-${secondsToTimeFormat(duration - currentTime)}`
        : secondsToTimeFormat(duration)

    return <div
        style={{ padding: "0 0 0 10px", width: "6ch" }}
        onClick={() => setInversed(!inversed)}
    >
        {time}
    </div>
}