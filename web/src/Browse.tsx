import React, { useState, useEffect } from "react"
import { Link } from "react-router-dom"
import { Table } from "./Table"
import ITrack from "./types/track"


export function GenreView({ name }: { name: string }) {
    const [tracks, setTracks] = useState<ITrack[]>()

    useEffect(() => {
        let isCanceled = false

        async function getGenre() {
            const response = await fetch(`/api/v1/browse/genres/${name}`)

            if (!response.ok)
                throw new Error(`Http request failed with status ${response.status}`)

            return await response.json()
        }

        getGenre()
            .then(tracks => {
                if (isCanceled)
                    return

                setTracks(tracks)
            })


        return () => { isCanceled = true }
    }, [name])

    if (!tracks)
        return <div>Loading</div>

    return <Table
        headerColumns={["Title", "Artist", "Album"]}
        rows={tracks.map(track => [
            <Link to={`/track/${track.id}`}>{track.title}</Link>,
            track.artist && <Link to={`/artist/${track.artist.id}`}>{track.artist.name}</Link>,
            track.album && <Link to={`/album/${track.album.id}`}>{track.album.name}</Link>,
        ])}
    />
}


type IList = {
    Name: string
    Urlname: string
}[]

export function BrowseView({ type }: { type: "artists" | "genres" }) {
    const [list, setList] = useState<IList | "loading">("loading")

    console.log(type)

    useEffect(() => {
        let isCanceled = false

        async function getBrowseView() {
            const response = await fetch(`/api/v1/browse/${type}`)

            const list: IList = await response.json()

            return list
        }

        getBrowseView()
            .then(list => {
                if (isCanceled)
                    return

                setList(list)
            })

        return () => { isCanceled = true }
    }, [type])

    return <div>
        <h1>{type}</h1>
        {list === "loading"
            ? <p>Loading...</p>
            : <ol>
                {list.map(({ Name, Urlname }) =>
                    <Link to={`/browse/${type}/${Urlname}`}>
                        <li>{Name}</li>
                    </Link>
                )}
            </ol>
        }
    </div>
}
