import usePlayer from "hooks/usePlayer"
import React, { useCallback, useContext, useMemo, useState } from "react"
import ITrack from "types/track"
import { ControlsContextProvider } from "./ControlsContext"
import { ProgressContextProvider } from "./ProgressContext"
import { TrackContextProvider } from "./TrackContext"

interface IAppContext {
    play: (id: number) => void
}

const AppContext = React.createContext<IAppContext>({} as IAppContext)

export function useAppContext() {
    const value = useContext(AppContext)

    if (!value)
        throw new Error("Tried to access ProgressContext outside of provider")

    return value
}

export function AppContextProvider({ children }: { children: React.ReactNode }) {
    const [currentTrack, setCurrentTrack] = useState<ITrack>()
    const { play, pause, playTrack, seek, playbackState, progress } = usePlayer()

    const playTrackById = useCallback((id: number) => {
        fetchTrackDetails(id)
            .then(track => setCurrentTrack(track))

        playTrack(id)
    }, [playTrack])

    return <AppContext.Provider value={useMemo(() => ({ play: playTrackById }), [playTrackById])}>
        <TrackContextProvider value={useMemo(() => currentTrack, [currentTrack])}>
            <ProgressContextProvider value={useMemo(() => progress, [progress])}>
                <ControlsContextProvider value={useMemo(() => ({ play, pause, seek, playbackState }), [play, pause, seek, playbackState])}>
                    {children}
                </ControlsContextProvider>
            </ProgressContextProvider>
        </TrackContextProvider>
    </AppContext.Provider>
}

async function fetchTrackDetails(id: number) {
    const response = await fetch(`/api/v1/track/${id}`, {
        headers: {
            "Content-Type": "application/json"
        }
    })

    if (!response.ok)
        throw new Error(`Http request failed with status code: ${response.status}`)

    const track: ITrack = await response.json()

    return track
}