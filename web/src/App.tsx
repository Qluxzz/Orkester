import React from "react"

import {
  BrowserRouter,
  Switch,
  Route,
  useParams
} from "react-router-dom"

import styled from "styled-components"
import { GetAlbumWithId } from "Features/Album/Album"
import { GetArtistWithId } from "Artist"
import { PlayerContextProvider } from "Context"
import PlayerBar from "Features/Player/PlayerBar"
import SearchBar from "Features/Search/SearchBar"
import SearchResults from "Features/Search/SearchResults"
import LikedTracks from "Features/Playlist/LikedTracks"

const Container = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1 1 0;
  overflow: hidden;
`

const Content = styled.div`
  padding: 10px;
  overflow: auto;
  flex: 1 1 0;
  display: flex;
  flex-direction: column;
`


function App() {
  return <PlayerContextProvider>
    <Container>
      <BrowserRouter>
        <Route path={["/search/:query", "/"]} component={SearchBar} />
        <Content>
          <Switch>
            <Route path="/album/:id" component={AlbumViewWrapper} />
            <Route path="/artist/:id" component={ArtistViewWrapper} />
            <Route path="/search/:query" component={SearchViewWrapper} />
            <Route path="/playlist/liked" component={LikedTracksViewWrapper} />
          </Switch>
        </Content>
        <PlayerBar />
      </BrowserRouter>
    </Container>
  </PlayerContextProvider >
}

function AlbumViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetAlbumWithId id={parseInt(id)} />
}

function ArtistViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetArtistWithId id={parseInt(id)} />
}

function SearchViewWrapper() {
  const { query } = useParams<{ query: string }>()

  return <SearchResults query={query} />
}

function LikedTracksViewWrapper() {
  return <LikedTracks />
}

export default App;
