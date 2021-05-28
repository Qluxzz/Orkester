import React, { useState, useContext, useEffect, useRef, useCallback, useMemo } from "react"
import IPlaybackState from "types/playbackState"
import ITrack from "./types/track"

interface IPlayerContext {
    track?: ITrack
    play: (id: number) => void
    togglePlayback: () => void
    state: IPlaybackState
    player?: HTMLAudioElement
}

const PlayerContext = React.createContext<IPlayerContext>({
    track: undefined,
    play: () => { throw new Error("Tried to access context outside of provider") },
    togglePlayback: () => { throw new Error("Tried to access context outside of provider") },
    state: "paused",
    player: undefined
})

export const usePlayerContext = () => useContext(PlayerContext)

async function fetchTrackDetails(id: number) {
    const response = await fetch(`/api/v1/track/${id}`)
    if (!response.ok)
        throw new Error(`Http request failed with status code: ${response.status}`)

    const track: ITrack = await response.json()

    return track
}

interface ILocalStorageTrack {
    id: number
    timestamp: number
}


export function PlayerContextProvider({ children }: { children: React.ReactNode }) {
    const [track, setTrack] = useState<ITrack>()
    const [state, setState] = useState<IPlaybackState>("paused")

    const playerRef = useRef(new Audio())
    const player = playerRef.current

    const loadTrack = useCallback((id: number) =>
        fetchTrackDetails(id)
            .then(track => {
                setTrack(track)
                localStorage.setItem("track", JSON.stringify({ id: track.id, timestamp: 0 }))
                document.title = `${track.title} - ${track.artists.map(artist => artist.name).join(", ")}`
                player.src = `/api/v1/track/${track.id}/stream`
                player.play()
                setState("playing")
            })
            .catch(error => {
                console.error("Something went wrong while loading track info", error)
            })
        , [player])

    const togglePlayBack = useCallback(async () => {
        if (player.paused) {
            await player.play()
            setState("playing")
        } else {
            player.pause()
            setState("paused")
        }
    }, [player])

    useEffect(() => {
        const savedTrackJson = localStorage.getItem("track")
        if (!savedTrackJson)
            return

        const { id, timestamp }: ILocalStorageTrack = JSON.parse(savedTrackJson)
        if (!id)
            return

        loadTrack(id)
            .then(_ => {
                player.fastSeek(timestamp)
            })
    }, [loadTrack, player])

    useEffect(() => {
        if (!track)
            return

        if (state !== "playing")
            return

        const interval = setInterval(
            () => {
                localStorage.setItem("track", JSON.stringify({ id: track.id, timestamp: player.currentTime }))
            },
            1000
        )

        return () => { clearInterval(interval) }
    }, [track, state])


    const memoValues = useMemo(() => ({
        track: track,
        play: loadTrack,
        togglePlayback: togglePlayBack,
        state: state,
        player: player
    }), [loadTrack, player, state, togglePlayBack, track])

    return <PlayerContext.Provider
        value={memoValues}
    >
        {children}
    </PlayerContext.Provider>
}