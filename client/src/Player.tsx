import React, { useState, useEffect } from "react"
import ITrack from "./types/track"
import styled from "styled-components"


const PlayerDiv = styled.div`
    background-radius: 20%;
    border: 1px solid green;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 10px;
`


export default function Player({ id }: { id: number }) {
    const [track, setTrack] = useState<ITrack | "loading">("loading")

    useEffect(() => {
        let isCanceled = false

        async function fetchTrackDetails() {
            try {
                const response = await fetch(`http://localhost:3001/track/${id}`)
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
        <h2>{track.Title}</h2>
        <p>{track.Artist} - {track.Album}</p>
        <img style={{ padding: 10, width: 256 }} src={`http://localhost:3001/track/${id}/image`} alt={track.Album} />
        <Controls id={id} />
    </PlayerDiv>
}

function Controls({ id }: { id: number }) {
    return <audio
        src={`http://localhost:3001/stream/${id}`}
        controls
    />
}