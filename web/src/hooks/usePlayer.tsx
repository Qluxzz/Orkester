import { useCallback, useEffect, useMemo, useRef, useState } from "react"
import ILocalStorageTrack from "types/localStorageTrack"
import IPlaybackState from "types/playbackState"

export default function usePlayer() {
    const [trackId, setTrackId] = useState<number>()
    const [playbackState, setPlaybackState] = useState<IPlaybackState>("paused")
    const [progress, setProgress] = useState<{ duration: number, currentTime: number }>({ duration: 0, currentTime: 0 })

    const playerRef = useRef(new Audio())
    const player = playerRef.current

    const playTrackAtMs = useCallback((id: number, timeStamp: number = 0) => {
        player.src = `/api/v1/track/${id}/stream`
        player
            .play()
            .then(() => player.fastSeek(timeStamp))
            .then(() => {
                setPlaybackState("playing")
                setTrackId(id)
            })
    }, [player])

    const pause = useCallback(() => {
        if (player.paused)
            return

        player.pause()
        setPlaybackState("paused")
    }, [player])

    const play = useCallback(async () => {
        await player.play()
        setPlaybackState("playing")
    }, [player])

    function writePlaybackStatusToLocalStorageWhilePlaying() {
        if (!trackId)
            return

        if (playbackState !== "playing")
            return

        const interval = setInterval(
            () => {
                writeTrackInfoToLocalStorage({
                    id: trackId,
                    timestamp: Math.round(player.currentTime)
                })
            },
            1000
        )

        return () => { clearInterval(interval) }
    }

    function updateProgress() {
        const interval = setInterval(() => {
            if (!player)
                return

            if (player.paused)
                return

            setProgress({
                duration: Math.round(player.duration),
                currentTime: Math.round(player.currentTime)
            })
        }, 1000)

        return () => {
            clearInterval(interval)
        }
    }

    useEffect(
        writePlaybackStatusToLocalStorageWhilePlaying,
        [trackId, playbackState, player]
    )

    useEffect(updateProgress, [player])

    const values = useMemo(() => ({
        playTrack: playTrackAtMs,
        playbackState,
        play,
        pause,
        progress
    }), [playTrackAtMs, playbackState, play, pause, progress])

    return values
}

function writeTrackInfoToLocalStorage(trackInfo: ILocalStorageTrack) {
    localStorage.setItem("track", JSON.stringify(trackInfo))
}