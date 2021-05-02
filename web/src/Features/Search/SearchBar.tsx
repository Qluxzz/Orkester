import { usePlayerContext } from "Context"
import React, { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import styled from "styled-components"
import ITrack, { IAlbum, IArtist } from "types/track"


const Bar = styled.div`
    display: flex;
    flex-direction: column;
    padding: 10px;
`

const Input = styled.input`
    flex-grow: 1;
`

const SearchResults = styled.div`

`

async function search(query: string) {
    const response = await fetch(`/api/v1/search/${query}`)

    if (response.ok)
        return response.json()

    throw new Error(`Http request failed with status code ${response.status}`)
}

interface ISearchResults {
    tracks: IArtist[]
    artists: IArtist[]
    albums: IAlbum[]
}

export default function SearchBar() {
    const [query, setQuery] = useState<string>()
    const [searchResults, setSearchResults] = useState<ISearchResults>()
    const { play } = usePlayerContext()

    useEffect(() => {
        if (!query)
            return

        let isCanceled = false

        search(query)
            .then(searchResults => {
                if (isCanceled)
                    return

                setSearchResults(searchResults)
            })
            .catch((error: Error) => {
                if (error.name === "AbortError")
                    return

                console.error("Failed to get search results", error)
            })

        return () => { isCanceled = true }
    }, [query])

    const noSearchResults = searchResults
        && searchResults.albums.length === 0
        && searchResults.tracks.length === 0
        && searchResults.artists.length === 0

    return <Bar>
        <Input onChange={e => setQuery(e.target.value)} value={query}></Input>
        {searchResults && <SearchResults>
            {noSearchResults
                ? <p>No results...</p>
                : <>
                    {searchResults.tracks.length > 0 && <>
                        <h1>Tracks</h1>
                        <ul>{searchResults.tracks.map(track => <li onClick={() => play(track.id)} key={track.id}>{track.name}</li>)}</ul>
                    </>}
                    {searchResults.albums.length > 0 && <>
                        <h1>Albums</h1>
                        <ul>{searchResults.albums.map(album => <Link to={`/album/${album.id}`}><li key={album.id}>{album.name}</li></Link>)}</ul>
                    </>}
                    {searchResults.artists.length > 0 && <>
                        <h1>Artists</h1>
                        <ul>{searchResults.artists.map(artist => <Link to={`/artist/${artist.id}`}><li key={artist.id}>{artist.name}</li></Link>)}</ul>
                    </>}
                </>
            }
        </SearchResults>}
    </Bar>
}

