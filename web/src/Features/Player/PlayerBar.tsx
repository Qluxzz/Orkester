import styled from "styled-components"
import { usePlayerContext } from "Context"
import { useEffect, useMemo, useRef } from "react"
import ITrack from "types/track"
import { AlbumLink, ArtistLink } from "utilities/Links"

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
            <AlbumLink {...track.album}>
                <img key={track.id} width="72" height="72" src={`/api/v1/album/${track.album.id}/image`} alt={track.album.name} />
            </AlbumLink>
            <div style={{ marginLeft: 10, overflow: "hidden" }}>
                <h1>{track.title}</h1>
                <h2 style={{ whiteSpace: "nowrap", textOverflow: "ellipsis" }}>
                    {track.artists.map((artist, i, arr) => <>
                        <ArtistLink {...artist} key={artist.id}>
                            {artist.name}
                        </ArtistLink>
                        {i !== arr.length - 1 && ", "}
                    </>)}
                    {" - "}
                    <AlbumLink {...track.album}>
                        {track.album.name}
                    </AlbumLink>
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

    useEffect(() => {
        playerRef
            .current?.play().then(_ => { })
    }, [track])

    return <audio
        ref={playerRef}
        src={`/api/v1/track/${track.id}/stream`}
        controls
        onPlay={() => {
            window.localStorage.setItem("track", track.id.toString())

            channel.postMessage("playing")
        }}
    />
}