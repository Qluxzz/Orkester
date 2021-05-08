import React, { useEffect, useState } from "react"
import { usePlayerContext } from "Context"
import { Link } from "react-router-dom"
import styled from "styled-components"
import { IArtist, IAlbum } from "types/track"

const Container = styled.div`
    display: flex;
    gap: 20px;

    div {
        flex: 1 1 0;
        max-width: 300px;
    }
`

const UnorderedListWithNoDots = styled.ul`
    list-style: none;
    padding: 0;
    margin: 0;

    li {
        margin: 5px 0;
        padding: 5px 0;
        text-decoration: underline;
        
        :hover {
            cursor: pointer;
        }
    }
`

interface ISearchResults {
    tracks: IArtist[]
    artists: IArtist[]
    albums: IAlbum[]
}

async function search(query: string) {
    const response = await fetch(`/api/v1/search/${query}`)

    if (!response.ok)
        throw new Error(`Http request failed with status code ${response.status}`)

    return response.json()
}


export default function SearchResults({ query }: { query: string }) {
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
                console.error("Failed to get search results", error)
            })

        return () => { isCanceled = true }
    }, [query])

    if (!searchResults)
        return null

    const noSearchResults = searchResults
        && searchResults.albums.length === 0
        && searchResults.tracks.length === 0
        && searchResults.artists.length === 0

    if (noSearchResults)
        return <p>No results...</p>

    return <Container>
        {searchResults.tracks.length > 0 && <div>
            <h1>Tracks</h1>
            <UnorderedListWithNoDots>{searchResults.tracks.map(track =>
                <li onClick={() => play(track.id)} key={track.id}>{track.name}</li>
            )}</UnorderedListWithNoDots>
        </div>}
        {searchResults.albums.length > 0 && <div>
            <h1>Albums</h1>
            <UnorderedListWithNoDots>{searchResults.albums.map(album =>
                <Link to={`/album/${album.id}`} key={album.id}><li>{album.name}</li></Link>
            )}</UnorderedListWithNoDots>
        </div>}
        {searchResults.artists.length > 0 && <div>
            <h1>Artists</h1>
            <UnorderedListWithNoDots>{searchResults.artists.map(artist =>
                <Link to={`/artist/${artist.id}`} key={artist.id}><li>{artist.name}</li></Link>
            )}</UnorderedListWithNoDots>
        </div>}
    </Container>
}