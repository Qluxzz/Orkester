import React, { useContext } from "react";
import ITrack from "types/track";

interface ITrackContext {
    track?: ITrack
}


const TrackContext = React.createContext<ITrackContext>({} as ITrackContext)

export const TrackContextProvider = TrackContext.Provider

export const useTrackContext = () => {
    const context = useContext(TrackContext)
    if (!context)
        throw new Error("TrackContext must be used inside of a TrackContextProvider")

    return context
}