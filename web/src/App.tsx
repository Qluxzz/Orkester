import React, { useState } from "react"

import {
  BrowserRouter,
  Switch,
  Route,
  useParams
} from "react-router-dom"

import styled from "styled-components"
import { GetAlbumById } from "Features/Album/Album"
import { GetArtistWithId } from "Artist"
import PlayerBar from "Features/Player/PlayerBar"
import SearchBar from "Features/Search/SearchBar"
import SearchResults from "Features/Search/SearchResults"
import LikedTracks from "Features/Playlist/LikedTracks"
import SideBar from "Features/SideBar/SideBar"
import { AppContextProvider } from "Context/AppContext"

const Container = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1 1 0;
  overflow: hidden;
`

const Content = styled.div`
  overflow: auto;
  flex: 1 1 0;
  display: flex;
  flex-direction: row;
`

const MainContent = styled.main`
  display: flex;
  flex-direction: column;
  flex: 1 1 0;
  padding: 20px;
  gap: 20px;
  overflow: hidden;
`

const ScrollableContent = styled.div`
  overflow: auto;
`

type AlbumCoverSize = "large" | "small"

function App() {
  const [albumCoverSize, setAlbumCoverSize] = useState<AlbumCoverSize>("small")

  return <AppContextProvider>
    <Container>
      <BrowserRouter>
        <Content>
          <SideBar
            hideAlbumCover={() => setAlbumCoverSize("small")}
            showAlbumCover={albumCoverSize === "large"}
          />
          <MainContent>
            <Route path="/search/:query" children={() => <SearchBar />} />
            <ScrollableContent>
              <Switch>
                <Route path="/album/:id">
                  <AlbumViewWrapper />
                </Route>
                <Route path="/artist/:id">
                  <ArtistViewWrapper />
                </Route>
                <Route path="/search/:query">
                  <SearchViewWrapper />
                </Route>
                <Route path="/collection/tracks">
                  <LikedTracks />
                </Route>
              </Switch>
            </ScrollableContent>
          </MainContent>
        </Content>
        <PlayerBar
          showAlbumCover={albumCoverSize === "small"}
          hideAlbumCover={() => setAlbumCoverSize("large")}
        />
      </BrowserRouter>
    </Container>
  </AppContextProvider>
}

function AlbumViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetAlbumById id={parseInt(id)} />
}

function ArtistViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetArtistWithId id={parseInt(id)} />
}

function SearchViewWrapper() {
  const { query } = useParams<{ query: string }>()

  return <SearchResults query={query} />
}

export default App;
