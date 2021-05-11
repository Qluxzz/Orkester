import { Link } from "react-router-dom"
import styled from "styled-components"
import { usePlayerContext } from "Context"
import { useEffect, useRef } from "react"
import ITrack from "types/track"

const Bar = styled.div`
  display: flex;
  flex-direction: column;
  background: #333;
  padding: 10px;
`

export default function PlayerBar() {
    const { track } = usePlayerContext()


    if (!track)
        return <Bar>Nothing is currently playing...</Bar>

    return <Bar>
        <div style={{ display: "flex", marginBottom: 10 }}>
            <Link
                to={`/album/${track.album.id}/${track.album.urlName}`}
            >
                <img key={track.id} width="72" height="72" src={`/api/v1/album/${track.album.id}/image`} alt={track.album.name} />
            </Link>
            <div style={{ marginLeft: 10 }}>
                <h1>{track.title}</h1>
                <h2>
                    <Link
                        to={`/artist/${track.artist.id}/${track.artist.urlName}`}
                    >
                        {track.artist.name}
                    </Link>
                    {" - "}
                    <Link
                        to={`/album/${track.album.id}/${track.album.urlName}`}
                    >
                        {track.album.name}
                    </Link>
                </h2>
            </div>
        </div>
        <Controls track={track} />
    </Bar>
}

function Controls({ track }: { track: ITrack }) {
    const playerRef = useRef<HTMLAudioElement>(null)

    useEffect(() => {
        /*
            A StorageEvent is sent to a window when a storage area it 
            has access to is changed within the context of another document.
        */
        window.addEventListener("storage", () => {
            // If playing in other document, pause this playback
            playerRef.current?.pause()
        })
    }, [])

    return <audio
        ref={playerRef}
        src={`/api/v1/track/${track.id}/stream`}
        controls
        style={{
            flexGrow: 1
        }}
        autoPlay
        onPlay={() => {
            window.localStorage.setItem("track", track.id.toString())
            window.localStorage.setItem("playingSession", (Math.random() * 1000).toString())
        }}
    />
}