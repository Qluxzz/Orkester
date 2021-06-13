import { useEffect, useState } from "react";
import { useHistory } from "react-router-dom"
import styled from "styled-components";

import CenteredDotLoader from "CenteredDotLoader"
import { AlbumLink } from "utilities/Links";
import AlbumImage from "utilities/AlbumImage";

interface IArtist {
    id: number
    name: string
    urlName: string
    albums: {
        id: number,
        name: string,
        urlName: string
    }[]
}

type Artist = "loading" | "loading-failed" | IArtist


export function GetArtistWithId({ id }: { id: number }) {
    const history = useHistory()
    const [artist, setArtist] = useState<Artist>("loading")

    useEffect(() => {
        let isCanceled = false

        async function fetchArtistInfo(): Promise<IArtist> {
            const response = await fetch(`/api/v1/artist/${id}`)

            if (!response.ok)
                throw new Error(`Http request failed with status ${response.status}`)

            return await response.json()
        }

        setArtist("loading")

        fetchArtistInfo()
            .then(artist => {
                if (isCanceled)
                    return

                setArtist(artist)
                history.replace(`/artist/${artist.id}/${artist.urlName}`)
            })
            .catch(error => {
                if (isCanceled)
                    return

                setArtist("loading-failed")

                console.error("Failed to get artist info!", error)
            })

        return () => { isCanceled = true }
    }, [id, history])

    if (artist === "loading")
        return <CenteredDotLoader />

    if (artist === "loading-failed")
        return <p>Failed to load artist</p>

    return <ArtistView {...artist} />
}

const Album = styled.div`
    display: flex;
    flex-direction: column;
    flex: 1 1 0;
    background: #333;
    padding: 10px;

    picture {
        position: relative;
        overflow: hidden;
        height: 0;
        padding-top: 100%;

        img {
            display: block;
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
        }
    }

    p {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-weight: bold;
        padding: 10px 0 5px 0;
        line-height: 1;
    }
`

const ArtistName = styled.h1`
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-bottom: 20px;
`


function ArtistView(artist: IArtist) {
    return <>
        <ArtistName>{artist.name}</ArtistName>
        <div style={{
            display: "grid",
            gap: 24,
            gridTemplateColumns: "repeat(auto-fill, minmax(240px, 1fr))",
            gridTemplateRows: "1fr"
        }}>
            {artist.albums.map(album =>
                <AlbumLink {...album}>
                    <Album>
                        <AlbumImage album={album} />
                        <p>{album.name}</p>
                    </Album>
                </AlbumLink>
            )}
        </div>
    </>
}