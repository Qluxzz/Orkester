import React, { useState, useContext, useEffect, useRef, useCallback, useMemo } from "react"
import IPlaybackState from "types/playbackState"
import ITrack from "./types/track"

interface IPlayerContext {
    track?: ITrack
    play: (id: number) => void
    queueTrack: (id: number) => void
    queueTracks: (ids: number[]) => void
    togglePlayback: () => void
    state: IPlaybackState
    player: HTMLAudioElement
    queue: ITrack[]
}

const PlayerContext = React.createContext<IPlayerContext>({} as IPlayerContext)

export const usePlayerContext = () => {
    const context = useContext(PlayerContext)
    if (!context)
        throw new Error("PlayerContext must be used inside of a PlayerContextProvider")

    return context
}

async function fetchTrackDetails(id: number) {
    const response = await fetch(`/api/v1/track/${id}`)
    if (!response.ok)
        throw new Error(`Http request failed with status code: ${response.status}`)

    const track: ITrack = await response.json()

    return track
}

async function fetchTracksDetails(ids: number[]) {
    const response = await fetch(`/api/v1/tracks/ids`, {
        method: "POST",
        body: JSON.stringify(ids),
        headers: {
            "Content-Type": "application/json"
        }
    })

    if (!response.ok)
        throw new Error(`Http request failed with status code: ${response.status}`)

    const tracks: ITrack[] = await response.json()

    return tracks
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
    const queue: ITrack[] = useMemo(() => [], [])
    const [currentlyPlayingTrack, setCurrentlyPlayingTrack] = useState<ITrack>()
    const [state, setState] = useState<IPlaybackState>("paused")

    const playerRef = useRef(new Audio())
    const player = playerRef.current

    const queueTracks = useCallback((ids: number[]) => {
        fetchTracksDetails(ids)
            .then(tracks => {
                const sortedTracks = ids.reduce<ITrack[]>((acc, id) => {
                    const track = tracks.find(x => x.id === id)
                    if (track)
                        acc.push(track)

                    return acc
                }, [])

                queue.push(...sortedTracks)
            })
            .catch(error => {
                console.error("Failed to get track details", error)
            })
    }, [queue])

    const queueTrack = useCallback((id: number) =>
        queueTracks([id]), [queueTracks])

    const playTrack = useCallback((track: ITrack) => {
        setCurrentlyPlayingTrack(track)
        document.title = `${track.title} - ${track.artists.map(artist => artist.name).join(", ")}`
        player.src = `/api/v1/track/${track.id}/stream`
        player.play().then(() => setState("playing"))
    }, [player])

    const playTrackById = useCallback((id: number) =>
        fetchTrackDetails(id)
            .then(track => playTrack(track))
            .catch(error => {
                console.error("Something went wrong while loading track info", error)
            }),
        [playTrack]
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

        playTrackById(savedTrack.id)
            .then(_ => {
                player.fastSeek(savedTrack.timestamp)
                player.pause()
            })
    }

    function writePlaybackStatusToLocalStorageWhilePlaying() {
        if (!currentlyPlayingTrack)
            return

        if (state !== "playing")
            return

        const interval = setInterval(
            () => {
                writeTrackInfoToLocalStorage({
                    id: currentlyPlayingTrack.id,
                    timestamp: Math.round(player.currentTime)
                })
            },
            1000
        )

        return () => { clearInterval(interval) }
    }

    /* Effects */

    useEffect(
        writePlaybackStatusToLocalStorageWhilePlaying,
        [currentlyPlayingTrack, state, player.currentTime]
    )

    useEffect(
        loadInitalPlaybackState,
        [playTrackById, player]
    )

    useEffect(() => {
        const playNextTrackInQueue = () => {
            const nextTrack = queue.shift()
            if (!nextTrack)
                return

            playTrack(nextTrack)
        }

        player.addEventListener("ended", playNextTrackInQueue)

        return () => {
            player.removeEventListener("ended", playNextTrackInQueue)
        }
    })


    const values = {
        track: currentlyPlayingTrack,
        play: playTrackById,
        queueTrack: queueTrack,
        queueTracks: queueTracks,
        togglePlayback: togglePlayBack,
        state: state,
        player: player
    }

    return <PlayerContext.Provider
        value={values}
    >
        {children}
    </PlayerContext.Provider>
}