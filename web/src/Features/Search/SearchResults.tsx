import React, { useEffect, useState } from "react"
import styled from "styled-components"
import { IArtist, IAlbum } from "types/track"
import CenteredDotLoader from "CenteredDotLoader"
import { AlbumLink, ArtistLink } from "utilities/Links"
import { useAppContext } from "Context/AppContext"

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

interface ITrack {
    id: number
    title: string
}

interface ISearchResults {
    tracks: ITrack[]
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
        return <CenteredDotLoader />

    const noSearchResults = searchResults
        && searchResults.albums.length === 0
        && searchResults.tracks.length === 0
        && searchResults.artists.length === 0

    if (noSearchResults)
        return <p>No results...</p>

    return <Container>
        <div>
            <h1>Tracks</h1>
            {searchResults.tracks.length === 0
                ? <p>No tracks found</p>
                : <UnorderedListWithNoDots>{
                    searchResults.tracks.map(track =>
                        <TrackRow key={track.id} track={track} />
                    )}
                </UnorderedListWithNoDots>
            }
        </div>
        <div>
            <h1>Albums</h1>
            {searchResults.albums.length === 0
                ? <p>No albums found</p>
                : <UnorderedListWithNoDots>
                    {searchResults.albums.map(album =>
                        <AlbumLink {...album} key={album.id}><li>{album.name}</li></AlbumLink>
                    )}
                </UnorderedListWithNoDots>
            }
        </div>
        <div>
            <h1>Artists</h1>
            {searchResults.artists.length === 0
                ? <p>No artists found</p>
                : <UnorderedListWithNoDots>
                    {searchResults.artists.map(artist =>
                        <ArtistLink {...artist} key={artist.id}><li>{artist.name}</li></ArtistLink>
                    )}
                </UnorderedListWithNoDots>
            }
        </div>
    </Container>
}

function TrackRow({ track }: { track: ITrack }) {
    const { play } = useAppContext()
    return <li onClick={() => play(track.id)} key={track.id}>{track.title}</li>
}