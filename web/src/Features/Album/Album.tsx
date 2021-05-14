import { useEffect, useState } from "react";
import styled from "styled-components";

import ITrack from "types/track";
import { secondsToTimeFormat } from "Utilities/secondsToTimeFormat";
import { usePlayerContext } from "Context";
import { Redirect, Route, Switch } from "react-router";
import CenteredDotLoader from "CenteredDotLoader";

interface IAlbum {
    id: number
    name: string
    urlName: string
    tracks: ITrack[]
}

export function GetAlbumWithId({ id }: { id: number }) {
    const [album, setAlbum] = useState<IAlbum>()

    useEffect(() => {
        let isCanceled = false

        fetchAlbumInfo(id)
            .then(album => {
                if (isCanceled)
                    return

                setAlbum(album)
            })
            .catch(error => {
                console.error("Failed to get album info!", error)
            })

        return () => { isCanceled = true }
    }, [id])

    if (!album)
        return <CenteredDotLoader />

    return <Switch>
        <Route path={`/album/${album.id}/${album.urlName}`}>
            <AlbumView {...album} />
        </Route>
        <Redirect to={`/album/${album.id}/${album.urlName}`} />
    </Switch>
}

const Row = styled.div`
    display: flex;
    padding: 10px 20px;
`
const HeaderRow = styled(Row)`
    border-bottom: 1px solid #333;
    margin-bottom: 10px;
`

const TrackRow = styled(Row)`
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

`



type ISortColumn = "trackNumber" | "title" | "length"
type ISortDirection = "ascending" | "descending"

interface ISortOptions {
    column: ISortColumn
    direction: ISortDirection
}

function AlbumView(album: IAlbum) {
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

    const sortedTracks = [...album.tracks].sort((a, b) => {
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

    const totalPlayTime = album.tracks.reduce((acc, x) => (acc += x.length), 0)

    return <div
        style={{
            display: "flex",
            justifyContent: "center",
            flexDirection: "column"
        }}
    >
        <header style={{ display: "flex", padding: 10 }}>
            <img src={`/api/v1/album/${album.id}/image`} style={{ width: 192 }} alt={`Album cover for ${album.name}`} />
            <div style={{ padding: 10 }}>
                <h1>{album.name}</h1>
                <p>{album.tracks.length} track{album.tracks.length !== 1 && "s"}, {secondsToTimeFormat(totalPlayTime)}</p>
            </div>
        </header>
        <section>
            <HeaderRow>
                <TrackNumber onClick={() => sortByColumn("trackNumber")}>#</TrackNumber>
                <TrackTitle onClick={() => sortByColumn("title")}>TITLE</TrackTitle>
                <TrackLength onClick={() => sortByColumn("length")}>🕒</TrackLength>
            </HeaderRow>
            {sortedTracks.map(track =>
                <TrackRow
                    key={track.id}
                    onClick={() => play(track.id)}
                >
                    <TrackNumber>{track.trackNumber}</TrackNumber>
                    <TrackTitle>{track.title}</TrackTitle>
                    <TrackLength>{secondsToTimeFormat(track.length)}</TrackLength>
                </TrackRow>
            )}
        </section>
    </div>
}

async function fetchAlbumInfo(id: number): Promise<IAlbum> {
    const response = await fetch(`/api/v1/album/${id}`)

    if (!response.ok)
        throw new Error(`Http request failed with status ${response.status}`)

    return await response.json()
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