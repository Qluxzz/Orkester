import { useAppContext } from "Context/AppContext"
import LikeButton from "Features/Album/LikeButton"
import styled from "styled-components"
import ITrack from "types/track"
import AlbumImage from "utilities/AlbumImage"
import ArtistList from "utilities/ArtistList"
import ellipsisTextMixin from "utilities/ellipsisText"
import { AlbumLink } from "utilities/Links"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import getColumnStyle from "./getColumnStyle"
import { IColumn } from "./TrackListBase"


function formatDate(d: Date): string {
    return `${d.getFullYear()}-${(d.getMonth() + 1).toString(10).padStart(2, "0")}-${d.getDate().toString(10).padStart(2, "0")}`
}

const TrackTitle = styled.p`
    ${_ => ellipsisTextMixin}
`

interface ITrackRow {
    columns: IColumn[]
    track: ITrack
    onLikedChanged?: (liked: boolean, trackId: number) => void
}

export function TrackRow({ columns, track, onLikedChanged }: ITrackRow) {
    const { play } = useAppContext()

    return <li
        style={{
            display: "flex",
            padding: 10,
            alignItems: "center",
            cursor: "default"
        }}
        onClick={() => play(track.id)}
    >
        {columns.map(column => {

            const children = (() => {
                switch (column.key) {
                    case "album":
                        return <AlbumLink {...track.album}>{track.album.name}</AlbumLink>
                    case "albumCover":
                        return <div style={{ marginRight: 10 }}>
                            <AlbumImage album={track.album} size={48} />
                        </div>
                    case "liked":
                        return <LikeButton
                            trackId={track.id}
                            liked={track.liked}
                            onLikeChanged={onLikedChanged}
                        />
                    case "title":
                        return <div style={{ marginRight: 10 }}>
                            <TrackTitle>{track.title}</TrackTitle>
                            <p style={{ fontSize: 12 }}><ArtistList artists={track.artists} /></p>
                        </div>
                    case "length":
                        return secondsToTimeFormat(track.length)
                    case "date":
                        return formatDate(new Date(Date.parse(track.date)))
                    default:
                        return track[column.key]
                }
            })()

            return <div
                key={column.key}
                style={getColumnStyle(column)}
            >
                {children}
            </div>
        })}
    </li>
}