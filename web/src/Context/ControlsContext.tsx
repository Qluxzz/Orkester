import React, { useContext } from "react"
import IPlaybackState from "types/playbackState"

interface IControlsContext {
    play: () => void
    pause: () => void
    playbackState: IPlaybackState
    seek: (time: number) => void
}

const ControlsContext = React.createContext<IControlsContext>({} as IControlsContext)

export function useControlsContext() {
    const value = useContext(ControlsContext)

    if (!value)
        throw new Error("Tried to access ControlsContext outside of provider")

    return value
}

export const ControlsContextProvider = ControlsContext.Provider