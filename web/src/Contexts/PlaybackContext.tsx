import React, { useContext } from "react";
import IPlaybackState from "types/playbackState"


interface IPlaybackContext {
    currentTime: number
    duration: number
    play: () => void
    pause: () => void
    seekToMs: (time: number) => void
    playbackState: IPlaybackState
}


const PlaybackContext = React.createContext<IPlaybackContext>({} as IPlaybackContext)

export const PlaybackContextProvider = PlaybackContext.Provider

export const usePlaybackContext = () => {
    const context = useContext(PlaybackContext)
    if (!context)
        throw new Error("PlaybackContext must be used inside of a PlaybackContextProvider")

    return context
}