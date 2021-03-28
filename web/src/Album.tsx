import React, { useEffect, useState } from "react";
import { Table } from "./Table";
import ITrack from "./types/track";

type IAlbum = {
    name: string
    artist: string
    year: Date
    tracks: ITrack[]
}

function secondsToTimeFormat(value: number): string {
    if (value < 60)
        return `00:${value.toString().padStart(2, "0")}`

    const oneHourInSeconds = 60 * 60

    if (value < oneHourInSeconds) {
        const minutes = Math.floor(value / 60)
        const seconds = Math.floor(value - Math.floor(value / 60) * 60)

        return `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`
    }

    throw new Error(`One hour and above is currently not handled`)
}


function AlbumView({ name, artist, year, tracks } : IAlbum) {
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
                track.title,
                secondsToTimeFormat(track.length)
            ])}
        />
    </div>
}

export function GetAlbumWithId({ id }: { id: number }) {
    const [album, setAlbum] = useState<IAlbum>()

    useEffect(() => {
        let isCanceled = false

        async function fetchAlbumInfo(): Promise<IAlbum> {
            const response = await fetch(`/api/v1/album/${id}`)

            if (!response.ok)
                throw new Error(`Http request failed with status ${response.status}`)

            return await response.json()
        }

        fetchAlbumInfo()
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
        return <div>Loading...</div>

    return <AlbumView {...album} />
}