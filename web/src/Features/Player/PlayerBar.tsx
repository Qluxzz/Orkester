import styled from "styled-components"
import { usePlayerContext } from "Context"
import { AlbumLink, ArtistLink } from "utilities/Links"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { useEffect, useState } from "react"

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
        <Controls />
    </Bar>
}

function Controls() {
    const { togglePlayback, state } = usePlayerContext()

    return <div>
        <button onClick={togglePlayback}>{state === "paused" ? "play" : "pause"}</button>
        <ProgressBar />
    </div>
}

function ProgressBar() {
    const [data, setData] = useState<{ duration: number, timestamp: number }>()
    const { player } = usePlayerContext()

    useEffect(() => {
        const interval = setInterval(() => {
            if (!player)
                return

            setData({
                duration: player.duration,
                timestamp: player.currentTime
            })
        }, 1000)

        return () => {
            clearInterval(interval)
        }
    }, [player])

    if (!data || !data.timestamp || !data.duration)
        return null

    return <>{secondsToTimeFormat(data.timestamp)}/{secondsToTimeFormat(data.duration)}</>
}