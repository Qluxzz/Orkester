import React, { useEffect, useState } from "react";
import styled from "styled-components";

import ITrack from "types/track";
import { secondsToHumanReadableFormat } from "utilities/secondsToTimeFormat";
import { useHistory } from "react-router";
import CenteredDotLoader from "CenteredDotLoader";
import { ArtistLink } from "utilities/Links";
import TrackList from "Features/TrackList/TrackList";
import AlbumImage from "utilities/AlbumImage";

interface IAlbum {
    id: number
    name: string
    urlName: string
    released: string
    artist: {
        id: number
        name: string
        urlName: string
    }
    tracks: ITrack[]
}

async function fetchAlbumInfo(id: number): Promise<IAlbum> {
    const response = await fetch(`/api/v1/album/${id}`)

    if (!response.ok)
        throw new Error(`Http request failed with status ${response.status}`)

    return await response.json()
}

interface IGetAlbumById {
    id: number
}

export function GetAlbumById({ id }: IGetAlbumById) {
    const history = useHistory()
    const [album, setAlbum] = useState<IAlbum>()

    useEffect(() => {
        if (id === album?.id)
            return

        let isCanceled = false

        fetchAlbumInfo(id)
            .then(album => {
                if (isCanceled)
                    return

                setAlbum(album)
                history.replace(`/album/${album.id}/${album.urlName}`)

            })
            .catch(error => {
                console.error("Failed to get album info!", error)
            })

        return () => { isCanceled = true }
    }, [id, history, album])

    if (!album)
        return <CenteredDotLoader />

    return <AlbumView album={album} />
}

const AlbumInfo = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: end;

    padding-left: 10px;
`

const Container = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
`

interface IAlbumView {
    album: IAlbum
}

function AlbumView({ album }: IAlbumView) {
    const totalPlayTime = album.tracks.reduce((acc, x) => (acc += x.length), 0)

    const releaseDate = new Date(album.released)

    return <Container>
        <header style={{ display: "flex", marginBottom: 20 }}>
            <AlbumImage album={album} size={192} />
            <AlbumInfo>
                <h1>{album.name}</h1>
                <div style={{ display: "flex", gap: 10 }}>
                    <ArtistLink {...album.artist} key={album.artist.id}>
                        <p>{album.artist.name}</p>
                    </ArtistLink>
                    <p>{formatReleaseDate(releaseDate)}</p>
                    <p>{album.tracks.length} song{album.tracks.length !== 1 && "s"}, {secondsToHumanReadableFormat(totalPlayTime)}</p>

                </div>
            </AlbumInfo>
        </header>
        <TrackList tracks={album.tracks} />
    </Container >
}

function formatReleaseDate(date: Date) {
    const year = date.getFullYear()

    return `${year}`
}