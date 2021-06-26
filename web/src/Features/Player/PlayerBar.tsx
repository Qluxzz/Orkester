import React, { useState } from "react"
import styled from "styled-components"

import { AlbumLink } from "utilities/Links"
import AlbumImage from "utilities/AlbumImage"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { useEffect } from "react"
import ArtistList from "utilities/ArtistList"

import textEllipsisMixin from "utilities/ellipsisText"
import { useTrackContext } from "Context/TrackContext"
import { useControlsContext } from "Context/ControlsContext"
import { useProgressContext } from "Context/ProgressContext"

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

function Controls() {
    const { play, pause, playbackState } = useControlsContext()

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
            <Slider />
        </div>
    </div>
}

function Slider() {
    const { duration, currentTime } = useProgressContext()
    const { seek } = useControlsContext()

    const [sliderValue, setSliderValue] = useState(currentTime)
    const [interacting, setInteracting] = useState(false)

    useEffect(() => {
        if (interacting)
            return

        setSliderValue(currentTime)
    }, [currentTime, interacting])

    return <div style={{ display: "flex", alignItems: "center", flexGrow: 1 }}>
        <div style={{ padding: "0 10px" }}>{secondsToTimeFormat(sliderValue)}</div>
        <input
            style={{ width: "100%" }}
            type="range"
            min={0}
            max={duration}
            value={sliderValue}
            onMouseUp={e => {
                seek(e.currentTarget.valueAsNumber)
                setInteracting(false)
            }}
            onChange={e => {
                setSliderValue(e.currentTarget.valueAsNumber)
                setInteracting(true)
            }}
        />
        <DurationOrRemainingTime
            duration={duration}
            currentTime={sliderValue}
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