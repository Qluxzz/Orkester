import React from "react"

import {
  BrowserRouter,
  Switch,
  Route,
  useParams,
  Link
} from "react-router-dom"

import styled from "styled-components"
import { GetAlbumWithId } from "Features/Album/Album"
import { GetArtistWithId } from "Artist"
import { PlayerContextProvider } from "Context"
import PlayerBar from "Features/Player/PlayerBar"
import SearchBar from "Features/Search/SearchBar"
import SearchResults from "Features/Search/SearchResults"

const Container = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1 1 0;
`

const Content = styled.div`
  padding: 10px;
  overflow: auto;
  flex: 1 1 0;
`


function App() {
  return <PlayerContextProvider>
    <Container>
      <BrowserRouter>
        <SearchBar />
        <Content>
          <Switch>
            <Route path="/album/:id" component={AlbumViewWrapper} />
            <Route path="/artist/:id" component={ArtistViewWrapper} />
            <Route path="/search/:query" component={SearchViewWrapper} />
            <Route path="/">
              <Link to="/track/80">Press here plz</Link>
              <div>Welcome home!</div>
            </Route>
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

export default App;
