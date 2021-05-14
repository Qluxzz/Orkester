import { Link } from "react-router-dom"
import styled from "styled-components"
import { usePlayerContext } from "Context"
import { useEffect, useMemo, useRef } from "react"
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
            <div style={{ marginLeft: 10, overflow: "hidden" }}>
                <h1>{track.title}</h1>
                <h2 style={{ whiteSpace: "nowrap", textOverflow: "ellipsis" }}>
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
    const channel = useMemo(() => new BroadcastChannel("currently_playing"), [])

    useEffect(() => {
        channel.onmessage = _ => {
            playerRef.current?.pause()
        }

        return () => {
            channel.close()
        }
    }, [channel])

    return <audio
        ref={playerRef}
        src={`/api/v1/track/${track.id}/stream`}
        controls
        autoPlay
        onPlay={() => {
            window.localStorage.setItem("track", track.id.toString())

            channel.postMessage("playing")
        }}
    />
}