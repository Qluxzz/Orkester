import { useEffect, useState } from "react";
import { useHistory } from "react-router-dom"
import styled from "styled-components";

import CenteredDotLoader from "CenteredDotLoader"
import { AlbumLink } from "utilities/Links";

interface IArtist {
    id: number
    name: string
    urlName: string
    albums: IAlbum[]
}

interface IAlbum {
    id: number
    name: string
    urlName: string
}

export function GetArtistWithId({ id }: { id: number }) {
    const history = useHistory()
    const [artist, setArtist] = useState<IArtist>()

    useEffect(() => {
        let isCanceled = false

        async function fetchArtistInfo(): Promise<IArtist> {
            const response = await fetch(`/api/v1/artist/${id}`)

            if (!response.ok)
                throw new Error(`Http request failed with status ${response.status}`)

            return await response.json()
        }

        fetchArtistInfo()
            .then(artist => {
                if (isCanceled)
                    return

                setArtist(artist)
                history.replace(`/artist/${artist.id}/${artist.urlName}`)
            })
            .catch(error => {
                console.error("Failed to get artist info!", error)
            })

        return () => { isCanceled = true }
    }, [id, history])

    if (!artist)
        return <CenteredDotLoader />

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
        <div>
            <ArtistName>{artist.name}</ArtistName>
        </div>
        <div style={{
            display: "grid",
            gap: 24,
            gridTemplateColumns: "repeat(auto-fill, minmax(240px, 1fr))",
            gridTemplateRows: "1fr"
        }}>
            {artist.albums.map(album =>
                <AlbumLink {...album}>
                    <Album>
                        <picture>
                            <img src={`/api/v1/album/${album.id}/image`} alt={`Album cover for ${album.name} by ${artist.name}`} />
                        </picture>
                        <p>{album.name}</p>
                    </Album>
                </AlbumLink>
            )}
        </div>
    </>
}