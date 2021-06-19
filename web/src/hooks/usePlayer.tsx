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
        try {
            setTrackId(id)
            player.src = `/api/v1/track/${id}/stream`
            player.preload = timeStamp.toString()

            player
                .play()
                .then(() => player.fastSeek(timeStamp))
        } catch (e) {
            const error: Error = e

            if (error.name === "AbortError")
                return

            console.error(error)
        }
    }, [player])

    const pause = useCallback(() => {
        if (player.paused)
            return

        player.pause()
    }, [player])

    const play = useCallback(async () => {
        await player.play()
    }, [player])

    function writePlaybackStatusToLocalStorageWhilePlaying() {
        if (!trackId)
            return

        if (player.paused)
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

    useEffect(
        writePlaybackStatusToLocalStorageWhilePlaying,
        [trackId, playbackState, player]
    )

    useEffect(() => {
        const setPlaying = () => setPlaybackState("playing")
        const setPaused = () => setPlaybackState("paused")

        player.addEventListener("play", setPlaying)
        player.addEventListener("pause", setPaused)

        return () => {
            player.removeEventListener("play", setPlaying)
            player.removeEventListener("pause", setPaused)
        }
    }, [player])

    useEffect(() => {
        const updateProgress = () => setProgress({
            duration: Math.round(player.duration),
            currentTime: Math.round(player.currentTime)
        })

        player.addEventListener("timeupdate", updateProgress)

        return () => {
            player.removeEventListener("timeupdate", updateProgress)
        }
    })

    const seek = useCallback((time: number) => {
        player.fastSeek(time)
    }, [player])

    const values = useMemo(() => ({
        playTrack: playTrackAtMs,
        playbackState,
        play,
        pause,
        progress,
        seek
    }), [playTrackAtMs, playbackState, play, pause, progress, seek])

    return values
}

function writeTrackInfoToLocalStorage(trackInfo: ILocalStorageTrack) {
    localStorage.setItem("track", JSON.stringify(trackInfo))
}