import {
  BrowserRouter,
  Switch,
  Route,
  useParams,
  Link
} from "react-router-dom"

import Player from "./Player";

import styled from "styled-components"
import { BrowseView, GenreView } from "./Browse";
import { GetAlbumWithId } from "./Album";
import { GetArtistWithId } from "./Artist"

const AppStyle = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
`


function App() {
  return <AppStyle>
    <BrowserRouter>
      <Switch>
        <Route path="/track/:id" component={PlayerWrapper} />
        <Route path="/browse/genres/:name" component={GenreViewWrapper} />
        <Route path="/browse/genres" children={<BrowseView type="genres" />} />
        <Route path="/browse/artists" children={<BrowseView type="artists" />} />
        <Route path="/album/:id" component={AlbumViewWrapper} />
        <Route path="/artist/:id" component={ArtistViewWrapper} />
        <Route path="/">
          <Link to="/track/80">Press here plz</Link>
          <div>Welcome home!</div>
        </Route>
      </Switch>
    </BrowserRouter>
  </AppStyle>
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
