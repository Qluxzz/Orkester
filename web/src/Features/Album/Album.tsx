import React, { useEffect, useState } from "react";
import styled from "styled-components";

import ITrack from "types/track";
import { secondsToTimeFormat } from "utilities/secondsToTimeFormat";
import { useHistory } from "react-router";
import CenteredDotLoader from "CenteredDotLoader";
import { ArtistLink } from "utilities/Links";
import TrackList from "TrackList";

interface IAlbum {
    id: number
    name: string
    urlName: string
    date: string
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

export function GetAlbumWithId({ id }: { id: number }) {
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

    return <AlbumView {...album} />
}

const AlbumInfo = styled.div`
    padding: 10px;

    *:not(:last-child) {
        display: block;
        margin-bottom: 10px;
    }
`

const Container = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
`

function AlbumView(album: IAlbum) {
    const totalPlayTime = album.tracks.reduce((acc, x) => (acc += x.length), 0)

    return <Container>
        <header style={{ display: "flex" }}>
            <img
                src={`/api/v1/album/${album.id}/image`}
                style={{ width: 192, height: 192, alignSelf: "center" }}
                alt={`Album cover for ${album.name}`}
            />
            <AlbumInfo>
                <h1>{album.name}</h1>
                <p>{album.tracks.length} track{album.tracks.length !== 1 && "s"}, {secondsToTimeFormat(totalPlayTime)}</p>
                <ArtistLink {...album.artist} key={album.artist.id}>
                    <p>{album.artist.name}</p>
                </ArtistLink>
                <p>{album.date}</p>
            </AlbumInfo>
        </header>
        <TrackList tracks={album.tracks} />
    </Container>
}
