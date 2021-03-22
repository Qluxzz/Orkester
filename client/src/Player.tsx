import React, { useState, useEffect } from "react"
import IPlaybackState from "./types/playbackState"
import PlaybackButton from "./PlaybackButton"

export default function Player({ id }: { id: number }) {
    const [playbackState, setPlaybackState] = useState<IPlaybackState>("paused")
    const audio = new Audio(`http://localhost:3000/stream/${id}`)

    useEffect(() => {
        async function togglePlayback() {
            switch (playbackState) {
                case "paused":
                    if (!audio.paused)
                        audio.pause()
                    break
                case "playing":
                    await audio.play()
                    break
            }
        }

        togglePlayback()
    }, [playbackState])

    function togglePlayback() {
        setPlaybackState(
            currentState => {
                switch (currentState) {
                    case "playing":
                        return "paused"
                    case "paused":
                        return "playing"
                }
            }
        )
    }

    return <PlaybackButton
        playbackState={playbackState}
        togglePlayback={togglePlayback}
    />
}
