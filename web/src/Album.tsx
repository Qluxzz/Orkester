import React, { useEffect, useState } from "react";
import { Table } from "./Table";
import ITrack from "./types/track";

type IAlbum = {
    name: string
    artist: string
    year: Date
    tracks: ITrack[]
}

function padWithTwoLeadingZeroes(value: number): string {
    return value.toString(10).padStart(2, "0")
}

// Convert seconds to (hours):(minutes):(seconds)
// Example: secondsToTimeFormat(1337) => 22:17
export function secondsToTimeFormat(value: number): string {
    const hours = Math.floor(value / 3600)
    value -= hours * 3600
    const minutes = Math.floor(value / 60)
    value -= minutes * 60
    const seconds = value

    const parts = []

    if (hours > 0)
        parts.push(padWithTwoLeadingZeroes(hours))

    parts.push(padWithTwoLeadingZeroes(minutes))
    parts.push(padWithTwoLeadingZeroes(seconds))

    return parts.join(":")
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