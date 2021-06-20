import React, { useState } from "react"
import styled from "styled-components"

import { AlbumLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { useEffect } from "react"
import ArtistList from "utilities/ArtistList"

import textEllipsisMixin from "utilities/ellipsisText"
import ITrack from "types/track"
import IPlaybackState from "types/playbackState"

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

interface IPlayerBar {
    track?: ITrack
    play: () => void
    pause: () => void
    playbackState: IPlaybackState
    seek: (ms: number) => void
    currentTime: number
    duration: number
}


export default function PlayerBar({
    track,
    play,
    pause,
    playbackState,
    seek,
    currentTime,
    duration
}: IPlayerBar) {
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
        <Controls
            play={play}
            pause={pause}
            playbackState={playbackState}
            seek={seek}
            currentTime={currentTime}
            duration={duration}
        />
    </Bar>
}

interface IControls {
    play: () => void
    pause: () => void
    playbackState: IPlaybackState
    seek: (ms: number) => void
    currentTime: number
    duration: number
}


function Controls({ play, pause, playbackState, seek, currentTime, duration }: IControls) {
    return <div>
        <div>
            <button
                onClick={playbackState === "paused"
                    ? play
                    : pause
                }
                style={{ width: 72 }}
            >
                {playbackState === "paused" ? "play" : "pause"}
            </button>
        </div>
        <div>
            <Slider
                duration={duration}
                currentTime={currentTime}
                seek={seek}
            />
        </div>
    </div>
}

interface ISlider {
    seek: (ms: number) => void
    duration: number
    currentTime: number
}

function Slider({ seek, duration, currentTime }: ISlider) {
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

type IState = "duration" | "timeLeft"

function DurationOrRemainingTime({ duration, currentTime }: { duration: number, currentTime: number }) {
    const [state, setState] = useState<IState>("duration")

    function toggle() {
        const newState = (() => {
            switch (state) {
                case "duration":
                    return "timeLeft"
                case "timeLeft":
                    return "duration"
            }
        })()

        setState(newState)
    }

    const time = (() => {
        switch (state) {
            case "duration":
                return secondsToTimeFormat(duration)
            case "timeLeft":
                return `-${secondsToTimeFormat(duration - currentTime)}`
        }
    })()

    return <div
        style={{ padding: "0 0 0 10px", width: "6ch" }}
        onClick={toggle}
    >
        {time}
    </div>
}
