import { usePlayerContext } from "Context"
import LikeButton from "Features/Album/LikeButton"
import React, { useState } from "react"
import styled from "styled-components"
import ITrack, { ILikeStatus } from "types/track"
import { ArtistLink } from "utilities/Links"
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"

const Row = styled.div`
    display: flex;
    padding: 10px 20px;
`
const HeaderRow = styled(Row)`
    border-bottom: 1px solid #333;
    margin-bottom: 10px;
`

const TrackRow = styled(Row)`
    align-items: center;

    a {
        font-size: 14px;
    }

    :hover {
        background: #333;
    }
`

const TrackNumber = styled.div`
    width: 50px;
`

const TrackTitle = styled.div`
    flex: 1 1 0px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    padding-right: 20px;
`

const TrackLength = styled.div`
    font-variant-numeric: tabular-nums;
`


type ISortColumn = "trackNumber" | "title" | "length"
type ISortDirection = "ascending" | "descending"

interface ISortOptions {
    column: ISortColumn
    direction: ISortDirection
}

interface ITrackList {
    tracks: ITrack[]
    onLikeStatusChanged?: (status: ILikeStatus, trackId: number) => void
}

export default function TrackList({ tracks, onLikeStatusChanged }: ITrackList) {
    const { play } = usePlayerContext()

    const [sortOptions, setSortOptions] = useState<ISortOptions>({
        column: "trackNumber",
        direction: "ascending"
    })

    function sortByColumn(column: ISortColumn) {
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

        switch (sortOptions.column) {
            case "length":
                return comparison(
                    a.length,
                    b.length
                )
            case "title":
                return comparison(
                    a.title.toLowerCase(),
                    b.title.toLowerCase()
                )
            case "trackNumber":
                return comparison(
                    a.trackNumber,
                    b.trackNumber
                )
            default:
                throw new Error("Unknown sort column")
        }
    })

    if (tracks.length === 0)
        return <div>No tracks</div>

    return <section>
        <HeaderRow>
            <TrackNumber onClick={() => sortByColumn("trackNumber")}>#</TrackNumber>
            <TrackTitle onClick={() => sortByColumn("title")}>TITLE</TrackTitle>
            <TrackLength onClick={() => sortByColumn("length")}>🕒</TrackLength>
        </HeaderRow>
        {sortedTracks.map(track =>
            <TrackRow
                key={track.id}
                onDoubleClick={() => play(track.id)}
            >
                <TrackNumber>{track.trackNumber}</TrackNumber>
                <TrackTitle>
                    <div>{track.title}</div>
                    {track.artists.map((artist, i, arr) => <>
                        <ArtistLink {...artist}>{artist.name}</ArtistLink>
                        {i !== arr.length - 1 && ", "}
                    </>)}
                </TrackTitle>
                <LikeButton
                    trackId={track.id}
                    likeStatus={track.likeStatus}
                    onLikeStatusChanged={onLikeStatusChanged}
                />
                <TrackLength>{secondsToTimeFormat(track.length)}</TrackLength>
            </TrackRow>
        )}
    </section>
}

function sortDirection(direction: ISortDirection) {
    return function <Type extends number | string>(a: Type, b: Type) {
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