import React from "react"

import {
  BrowserRouter,
  Switch,
  Route,
  useParams,
  Link
} from "react-router-dom"

import Player from "./Features/Player/Player";

import styled from "styled-components"
import { BrowseView, GenreView } from "./Features/Browse/Browse";
import { GetAlbumWithId } from "./Album";
import { GetArtistWithId } from "./Artist"
import { PlayerContextProvider } from "./Context";
import PlayerBar from "./Features/Player/PlayerBar";

const AppStyle = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  flex: 1 1 100%;
`


function App() {
  return <PlayerContextProvider>
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100%"
      }}
    >
      <BrowserRouter>
        <AppStyle>
          <Switch>
            <Route path="/track/:id" component={PlayerWrapper} />
            <Route path="/browse/genres/:name" component={GenreViewWrapper} />
            <Route path="/browse/genres" children={<BrowseView type="genre" />} />
            <Route path="/browse/artists" children={<BrowseView type="artist" />} />
            <Route path="/album/:id" component={AlbumViewWrapper} />
            <Route path="/artist/:id" component={ArtistViewWrapper} />
            <Route path="/">
              <Link to="/track/80">Press here plz</Link>
              <div>Welcome home!</div>
            </Route>
          </Switch>
        </AppStyle>
        <PlayerBar />
      </BrowserRouter>
    </div>
  </PlayerContextProvider >
}

function GenreViewWrapper() {
  const { name } = useParams<{ name: string }>()

  return <GenreView name={name} />
}

function PlayerWrapper() {
  const { id } = useParams<{ id: string }>()

  return <Player id={parseInt(id)} />
}

function AlbumViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetAlbumWithId id={parseInt(id)} />
}

function ArtistViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetArtistWithId id={parseInt(id)} />
}

export default App;


