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

function readTrackInfoFromLocalStorage(): ILocalStorageTrack | undefined {
    const savedTrackJson = localStorage.getItem("track")
    if (!savedTrackJson)
        return

    const savedTrack: ILocalStorageTrack = JSON.parse(savedTrackJson)
    if (!savedTrack)
        return

    return savedTrack
}

function writeTrackInfoToLocalStorage(trackInfo: ILocalStorageTrack) {
    localStorage.setItem("track", JSON.stringify(trackInfo))
}


export function PlayerContextProvider({ children }: { children: React.ReactNode }) {
    const [track, setTrack] = useState<ITrack>()
    const [state, setState] = useState<IPlaybackState>("paused")

    const playerRef = useRef(new Audio())
    const player = playerRef.current

    const playTrack = useCallback((id: number) =>
        fetchTrackDetails(id)
            .then(track => {
                setTrack(track)
                writeTrackInfoToLocalStorage({ id: track.id, timestamp: 0 })
                document.title = `${track.title} - ${track.artists.map(artist => artist.name).join(", ")}`
                player.src = `/api/v1/track/${track.id}/stream`
                player.play()
                setState("playing")
            })
            .catch(error => {
                console.error("Something went wrong while loading track info", error)
            }),
        [player]
    )

    const togglePlayBack = useCallback(async () => {
        if (player.paused) {
            await player.play()
            setState("playing")
        } else {
            player.pause()
            setState("paused")
        }
    }, [player])

    function loadInitalPlaybackState() {
        const savedTrack = readTrackInfoFromLocalStorage()

        if (!savedTrack)
            return

        playTrack(savedTrack.id)
            .then(_ => {
                player.fastSeek(savedTrack.timestamp)
                player.pause()
            })
    }

    function writePlaybackStatusToLocalStorageWhilePlaying() {
        if (!track)
            return

        if (state !== "playing")
            return

        const interval = setInterval(
            () => {
                localStorage.setItem(
                    "track",
                    JSON.stringify({
                        id: track.id,
                        timestamp: Math.round(player.currentTime)
                    })
                )
            },
            1000
        )

        return () => { clearInterval(interval) }
    }

    /* Effects */

    useEffect(
        writePlaybackStatusToLocalStorageWhilePlaying,
        [track, state, player.currentTime]
    )

    useEffect(
        loadInitalPlaybackState,
        [playTrack, player]
    )


    const memoValues = useMemo(() => ({
        track: track,
        play: playTrack,
        togglePlayback: togglePlayBack,
        state: state,
        player: player
    }), [playTrack, player, state, togglePlayBack, track])

    return <PlayerContext.Provider
        value={memoValues}
    >
        {children}
    </PlayerContext.Provider>
}