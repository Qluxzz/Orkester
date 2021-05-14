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

const HeaderRow = styled(Row)`
    border-bottom: 1px solid #333;
    margin-bottom: 10px;

    :hover {
        background: none;
    }
`

type ISorting = "trackNumber" | "title" | "length"
type ISortDirection = "ascending" | "descending"


function AlbumView({ name, tracks }: IAlbum) {
    const [sorting, setSorting] = useState<ISorting>("trackNumber")
    const [sortDirection, setSortDirection] = useState<ISortDirection>("ascending")
    const { play } = usePlayerContext()

    function sortByColumn(column: ISorting) {
        if (sorting === column)
            setSortDirection(
                sortDirection === "ascending"
                    ? "descending"
                    : "ascending"
            )
        else {
            setSorting(column)
            setSortDirection("ascending")
        }
    }

    const sortedTracks = [...tracks].sort((a, b) => {
        const comparison = sortDirection === "ascending"
            ? greaterThan
            : lesserThan

        switch (sorting) {
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
        }
    })


    return <div
        style={{
            display: "flex",
            justifyContent: "center",
            flexDirection: "column"
        }}
    >
        <h1>{name}</h1>
        <section>
            <HeaderRow>
                <TrackNumber onClick={() => sortByColumn("trackNumber")}>#</TrackNumber>
                <TrackTitle onClick={() => sortByColumn("title")}>TITLE</TrackTitle>
                <TrackLength onClick={() => sortByColumn("length")}>🕒</TrackLength>
            </HeaderRow>
            {sortedTracks.map(track =>
                <Row key={track.id} onClick={() => play(track.id)}>
                    <TrackNumber>{track.trackNumber}</TrackNumber>
                    <TrackTitle>{track.title}</TrackTitle>
                    <TrackLength>{secondsToTimeFormat(track.length)}</TrackLength>
                </Row>
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

function greaterThan<Type extends number | string>(a: Type, b: Type) {
    if (a === b)
        return 0
    return a > b
        ? 1
        : -1
}

function lesserThan<Type extends number | string>(a: Type, b: Type) {
    if (a === b)
        return 0
    return a < b
        ? 1
        : -1
}