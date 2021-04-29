import { useState, useEffect } from "react"
import ITrack from "./types/track"
import styled from "styled-components"
import { Link } from "react-router-dom"

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

                document.title = `${track.title} - ${track.album.name}`
            })
            .catch(error => console.error(error))

        return () => { isCanceled = true }
    }, [id])

    if (track === "loading")
        return <div>Loading</div>

    return <PlayerDiv>
        <p style={{ margin: 5, textAlign: "center" }}>Playing from album<br /><Link to={`/album/${track.album.id}/${track.album.urlName}`}>{track.album.name}</Link></p>
        <TrackImage src={`/api/v1/track/${id}/image`} alt={track.album.name} />
        <h2 style={{ margin: 5 }}>{track.title}</h2>
        <Link to={`/artist/${track.artist.id}/${track.artist.urlName}`}><p style={{ margin: 5 }}>{track.artist.name}</p></Link>
    </PlayerDiv>
}

