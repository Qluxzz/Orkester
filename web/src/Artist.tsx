import { useEffect, useState } from "react";
import { Link } from "react-router-dom"

interface IArtist {
    name: string
    albums: IAlbum[]
}

interface IAlbum {
    id: number
    name: string
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
                throw error
            })

        return () => { isCanceled = true }
    }, [id])

    if (!artist)
        return <div>Loading...</div>

    return <ArtistView {...artist} />
}

function ArtistView(artist: IArtist) {
    return <div>
        <h1>{artist.name}</h1>
        <ul>
            {artist.albums.map((album, i) => <li key={i}>
                <Link to={`/album/${album.id}/${album.urlName}`}><h2>{album.name}</h2></Link>
            </li>)}
        </ul>
    </div>
}