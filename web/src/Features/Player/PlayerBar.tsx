import { Link } from "react-router-dom"
import styled from "styled-components"
import { usePlayerContext } from "Context"

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
            <img width="72" height="72" src={`/api/v1/track/${track.id}/image`} alt={track.album.name} />
            <div style={{ marginLeft: 10 }}>
                <h1 style={{ margin: 0 }}>{track.title}</h1>
                <h2 style={{ margin: 0 }}>
                    <Link to={`/artist/${track.artist.id}`}>{track.artist.name}</Link> - <Link to={`/album/${track.album.id}`}>{track.album.name}</Link>
                </h2>
            </div>
        </div>
        <Controls id={track.id} />
    </Bar>
}

function Controls({ id }: { id: number }) {
    return <audio
        src={`/api/v1/track/${id}/stream`}
        controls
        style={{
            flexGrow: 1
        }}
        autoPlay
    />
}