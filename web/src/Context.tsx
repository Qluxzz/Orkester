import React, { useState, useContext, useEffect } from "react"
import ITrack from "./types/track"

interface IPlayerContext {
    track?: ITrack
    play: (id: number) => void
}

const PlayerContext = React.createContext<IPlayerContext>({
    track: undefined,
    play: () => { throw new Error("Tried to access context outside of provider") }
})

/*
 Track loading lifecycle,
 Inital value is undefined, there's no track at all, possibly one could store a track id in local storage that is fetched whenever the site is launched
 Then while the track is loading, it's status is set to loading
 If the request succedded, loading-finished and track is set correctly
 If the request failed, loading-failed with a modal or notification that the track failed to load
 
*/

export const usePlayerContext = () => useContext(PlayerContext)

async function fetchTrackDetails(id: number) {
    const response = await fetch(`/api/v1/track/${id}`)
    if (!response.ok)
        throw new Error(`Http request failed with status code: ${response.status}`)

    const track: ITrack = await response.json()

    return track
}


export function PlayerContextProvider({ children }: { children: React.ReactNode }) {
    const [track, setTrack] = useState<ITrack>()

    useEffect(() => {
        const savedTrack = localStorage.getItem("trackId")
        if (!savedTrack)
            return

        const trackId = parseInt(savedTrack)

        if (!trackId)
            return

        loadTrack(trackId)
    }, [])

    const loadTrack = (id: number) => {
        fetchTrackDetails(id)
            .then(track => {
                setTrack(track)
                localStorage.setItem("trackId", track.id.toString())
            })
            .catch(error => {
                console.error("Something went wrong while loading track info", error)
            })
    }

    return <PlayerContext.Provider value={{ track, play: loadTrack }}>
        {children}
    </PlayerContext.Provider>
}