import React, { useContext } from "react"
import ITrack from "types/track"

const TrackContext = React.createContext<(ITrack | undefined)>(undefined)

export function useTrackContext() {
    return useContext(TrackContext)
}

export const TrackContextProvider = TrackContext.Provider