import React, { useState, useEffect } from "react"
import ITrack from "./types/track"
import styled from "styled-components"


const PlayerDiv = styled.div`
    display: flex;
    flex-direction: column;
    padding: 20px;
    flex-grow: 1;
`

const TrackImage = styled.img`
    width: 100%;
    height: auto;
    border-radius: 10px;
    border: 2px solid white;

`


export default function Player({ id }: { id: number }) {
    const [track, setTrack] = useState<ITrack | "loading">("loading")

    useEffect(() => {
        let isCanceled = false

        async function fetchTrackDetails() {
            try {
                const response = await fetch(`/api/v1/track/${id}`)
                if (!response.ok)
                    throw new Error(`Http request failed with status code: ${response.status}`)

                const trackInfo: ITrack = await response.json()

                return trackInfo
            } catch (error) {
                console.log("Failed to get track info!", error)
                throw error
            }
        }

        fetchTrackDetails()
            .then(track => {
                if (!isCanceled)
                    setTrack(track)
            })
            .catch(error => console.error(error))

        return () => { isCanceled = true }
    }, [id])

    if (track === "loading")
        return <div>Loading</div>

    return <PlayerDiv>
        <p style={{ margin: 5, textAlign: "center" }}>Playing from album<br />{track.album.name}</p>
        <TrackImage src={`/api/v1/track/${id}/image`} alt={track.album.name} />
        <h2 style={{ margin: 5 }}>{track.title}</h2>
        <p style={{ margin: 5 }}>{track.artist.name}</p>
        <Controls id={id} />
    </PlayerDiv>
}

function Controls({ id }: { id: number }) {
    return <audio
        src={`/api/v1/track/${id}/stream`}
        controls
        style={{ 
            margin: "0 -20px -20px -20px",
            flexGrow: 1
         }}
    />
}