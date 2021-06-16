import React, { useState } from "react"
import styled from "styled-components"

import LikeButton from "Features/Album/LikeButton"
import ITrack from "types/track"
import AlbumImage from "utilities/AlbumImage"
import { AlbumLink, ArtistLink } from "utilities/Links"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"
import { usePlayerContext } from "Context"


type IColumnKey = ITrackKeys | "albumCover"

interface IColumn {
    display: string
    key: IColumnKey
    width?: number | "grow"
    centered?: boolean
}

interface ITrackListBase {
    tracks: ITrack[]
    initalSortColumn?: ITrackKeys
    columns: IColumn[]
    onLikedChanged?: (liked: boolean, trackId: number) => void
}

type ITrackKeys = keyof ITrack

type ISortDirection = "ascending" | "descending"

interface ISortOptions {
    column: ITrackKeys
    direction: ISortDirection
}

const defaultSortColumn: ITrackKeys = "trackNumber"

export default function TrackListBase({
    tracks,
    initalSortColumn = defaultSortColumn,
    columns,
    onLikedChanged
}: ITrackListBase) {
    const [sortOptions, setSortOptions] = useState<ISortOptions>({
        column: initalSortColumn,
        direction: "ascending"
    })

    function sortByColumn(column: IColumnKey) {
        if (column === "albumCover")
            return

        const sortDirection = (
            sortOptions.column === column
            && sortOptions.direction === "ascending"
        )
            ? "descending"
            : "ascending"

        setSortOptions({
            column: column,
            direction: sortDirection
        })
    }

    const sortedTracks = [...tracks].sort((a, b) => {
        const comparison = (() => {
            switch (sortOptions.direction) {
                case "ascending":
                    return greaterThan
                case "descending":
                    return lessThan
                default:
                    throw new Error(`Unknown sort directon ${sortOptions.direction}`)
            }
        })()

        const [colA, colB] = (() => {
            switch (sortOptions.column) {
                case "date":
                    return [new Date(Date.parse(a.date)), new Date(Date.parse(b.date))]
                case "album":
                    return [a.album.name.toLowerCase(), b.album.name.toLowerCase()]
                case "artists":
                    return [
                        a.artists.map(x => x.urlName).join(","),
                        b.artists.map(x => x.urlName).join(",")
                    ]
                default:
                    return [a[sortOptions.column], b[sortOptions.column]]
            }
        })()

        return comparison(colA, colB)
    })

    return <div>
        <header style={{ display: "flex", padding: "10px" }}>
            {columns.map(column => {
                const style: React.CSSProperties = {}

                if (typeof column.width === "number") {
                    style.width = column.width
                } else if (column.width === "grow") {
                    style.flexGrow = 1
                }

                if (column.centered) {
                    style.textAlign = "center"
                }

                return <div
                    style={style}
                    onClick={() => sortByColumn(column.key)}
                    key={column.key}
                >
                    {column.display}
                </div>
            })}
        </header>
        <StyledList>
            {sortedTracks.map(track =>
                <TrackRow
                    key={track.id}
                    track={track}
                    columns={columns}
                    onLikedChanged={onLikedChanged}
                />
            )}
        </StyledList>
    </div>
}

const StyledList = styled.ul`
    padding: 0;
    margin: 0;

    li:nth-child(odd) {
        background-color: #333;
    }
`

function formatDate(d: Date): string {
    return `${d.getFullYear()}-${(d.getMonth() + 1).toString(10).padStart(2, "0")}-${d.getDate()}`
}

interface ITrackRow {
    columns: IColumn[]
    track: ITrack
    onLikedChanged?: (liked: boolean, trackId: number) => void
}

function TrackRow({ columns, track, onLikedChanged }: ITrackRow) {
    const { play } = usePlayerContext()

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
                    case "artists":
                        return <>{track.artists.map(artist => <ArtistLink {...artist}>{artist.name}</ArtistLink>)}</>
                    case "albumCover":
                        return <div style={{ marginRight: 10 }}>
                            <AlbumImage album={track.album} size={72} />
                        </div>
                    case "liked":
                        return <LikeButton
                            trackId={track.id}
                            liked={track.liked}
                            onLikeChanged={onLikedChanged}
                        />
                    case "length":
                        return secondsToTimeFormat(track.length)
                    case "date":
                        return formatDate(new Date(Date.parse(track.date)))
                    default:
                        return track[column.key]
                }
            })()

            const style: React.CSSProperties = {}

            if (typeof column.width === "number") {
                style.width = column.width
            } else if (column.width === "grow") {
                style.flexGrow = 1
            }

            if (column.centered) {
                style.textAlign = "center"
            }

            return <div
                key={column.key}
                style={style}
            >
                {children}
            </div>
        })}
    </li>
}

function sortDirection(direction: ISortDirection) {
    return function <Type extends number | string | boolean | Date>(a: Type, b: Type) {
        if (a === b)
            return 0

        if (direction === "descending") {
            [a, b] = [b, a]
        }

        return a > b
            ? 1
            : -1
    }
}

const greaterThan = sortDirection("ascending")
const lessThan = sortDirection("descending")