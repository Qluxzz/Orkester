import { PlaybackContextProvider } from "Contexts/PlaybackContext"
import { TrackContextProvider } from "Contexts/TrackContext"
import usePlayer from "hooks/usePlayer"
import React, { useState, useContext, useEffect, useCallback, useMemo } from "react"
import ILocalStorageTrack from "types/localStorageTrack"
import ITrack from "types/track"

interface IPlayerContext {
    play: (id: number, timestamp?: number) => void
}

const PlayerContext = React.createContext<IPlayerContext>({} as IPlayerContext)

export const usePlayerContext = () => {
    const context = useContext(PlayerContext)
    if (!context)
        throw new Error("PlayerContext must be used inside of a PlayerContextProvider")

    return context
}

async function fetchTrackDetails(id: number) {
    return fetchTracksDetails([id])
        .then(tracks => tracks[0])
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

function readTrackInfoFromLocalStorage(): ILocalStorageTrack | undefined {
    const savedTrackJson = localStorage.getItem("track")
    if (!savedTrackJson)
        return

    const savedTrack: ILocalStorageTrack = JSON.parse(savedTrackJson)
    if (!savedTrack)
        return

    return savedTrack
}

export function PlayerContextProvider({ children }: { children: React.ReactNode }) {
    const [currentlyPlayingTrack, setCurrentlyPlayingTrack] = useState<ITrack>()
    const { playTrack, progress, playbackState, pause, play, seek, repeatState, setRepeat } = usePlayer()

    const _playTrack = useCallback((track: ITrack, timestamp?: number) => {
        setCurrentlyPlayingTrack(track)
        document.title = `${track.title} - ${track.artists.map(artist => artist.name).join(", ")}`
        playTrack(track.id, timestamp)

    }, [playTrack])

    const playTrackById = useCallback((id: number, timestamp?: number) =>
        fetchTrackDetails(id)
            .then(track => _playTrack(track, timestamp))
            .catch(error => {
                console.error("Something went wrong while loading track info", error)
            }),
        [_playTrack]
    )

    const loadInitalPlaybackState = useCallback(() => {
        const savedTrack = readTrackInfoFromLocalStorage()

        if (!savedTrack)
            return

        playTrackById(savedTrack.id, savedTrack.timestamp)
            .then(() => pause())
    }, [playTrackById, pause])

    /* Effects */

    useEffect(loadInitalPlaybackState, [loadInitalPlaybackState])

    /* Memoized context provider object */

    const playerContextValues = useMemo(() => ({
        play: playTrackById,
    }), [playTrackById])

    const trackContextValues = useMemo(() => ({ track: currentlyPlayingTrack }), [currentlyPlayingTrack])

    const playbackContextValues = useMemo(() => ({
        duration: progress.duration,
        currentTime: progress.currentTime,
        play,
        pause,
        playbackState: playbackState,
        seek,
        repeatState,
        setRepeat
    }), [progress, play, pause, playbackState, seek, repeatState, setRepeat])

    return <PlayerContext.Provider
        value={playerContextValues}
    >
        <TrackContextProvider value={trackContextValues}>
            <PlaybackContextProvider value={playbackContextValues}>
                {children}
            </PlaybackContextProvider>
        </TrackContextProvider>
    </PlayerContext.Provider>
}