import { useEffect, useState } from "react";
import { Link } from "react-router-dom"
import styled from "styled-components";

interface IArtist {
    name: string
    albums: IAlbum[]
}

interface IAlbum {
    id: number
    name: string
    year: string
    urlName: string
}

export function GetArtistWithId({ id }: { id: number }) {
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
            })
            .catch(error => {
                console.error("Failed to get artist info!", error)
            })

        return () => { isCanceled = true }
    }, [id])

    if (!artist)
        return <div>Loading...</div>

    return <ArtistView {...artist} />
}

const Album = styled.div`
    display: flex;
    flex-direction: column;
    flex: 1 1 0;
    background: #333;
    padding: 10px;

    img {
        display: block;
        max-width: 100%;
    }

    p {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-weight: bold;
        padding: 10px;
    }
`

const ArtistName = styled.h1`
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
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
            {artist.albums.map((album, i) => 
                <Link to={`/album/${album.id}/${album.urlName}`}>
                    <Album>
                    <img src={`/api/v1/album/${album.id}/image`} alt={`Album cover for ${album.name} by ${artist.name}`} />
                    <p>{album.name}</p>
                    <p>{album.year}</p>
                    </Album>
                </Link>
            )}
        </div>
    </>
}