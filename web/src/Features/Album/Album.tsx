import { useEffect, useState } from "react";
import { usePlayerContext } from "Context";
import Table from "Table";
import ITrack from "types/track";
import { secondsToTimeFormat } from "Utilities/secondsToTimeFormat";
import styled from "styled-components";

interface IAlbum {
    name: string
    artist: string
    year: Date
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
                throw error
            })

        return () => { isCanceled = true }
    }, [id])

    if (!album)
        return <div>Loading...</div>

    return <AlbumView {...album} />
}


const LinkButton = styled.button`
    background: none;
    color: white;
    border: none;
    text-decoration: underline;

    :hover {
        cursor: pointer;
    }
`

function AlbumView({ name, artist, year, tracks } : IAlbum) {
    const { play } = usePlayerContext()

    return <div>
        {name}
        {artist}
        {year}
        <Table
            headerColumns={[
                "#",
                "Name",
                "Length"
            ]}
            rows={tracks.map((track) => [
                track.trackNumber,
                <LinkButton 
                    type="button" 
                    onClick={() => play(track.id)}
                >
                    {track.title}
                </LinkButton>,
                secondsToTimeFormat(track.length)
            ])}
        />
    </div>
}

async function fetchAlbumInfo(id: number): Promise<IAlbum> {
    const response = await fetch(`/api/v1/album/${id}`)

    if (!response.ok)
        throw new Error(`Http request failed with status ${response.status}`)

    return await response.json()
}