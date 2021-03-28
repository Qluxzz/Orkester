import {
  BrowserRouter,
  Switch,
  Route,
  useParams,
  Link
} from "react-router-dom"

import Player from "./Player";

import styled from "styled-components"
import { useEffect, useState } from "react";
import ITrack from "./types/track";
import { StripedTable } from "./Table";

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
        <Route path="/track/:id" children={<PlayerWrapper />} />
        <Route path="/browse/genres/:name" children={<GenreViewWrapper />} />
        <Route path="/browse/genres" children={<BrowseView type="genres" />} />
        <Route path="/browse/artists" children={<BrowseView type="artists" />} />
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

  return <StripedTable
    headerColumns={["Title", "Artist", "Album"]}
    rows={tracks.map(track => [
      <Link to={`/track/${track.Id}`}>{track.Title}</Link>,
      <Link to={`/album/${track.Album}`}>{track.Album}</Link>,
      <Link to={`/artist/${track.Artist}`}>{track.Artist}</Link>
    ])}
  />
}


type IList = {
  Name: string
  Urlname: string
}[]

function BrowseView({ type }: { type: "artists" | "genres" }) {
  const [list, setList] = useState<IList | "loading">("loading")

  console.log(type)

  useEffect(() => {
    async function getBrowseView() {
      const response = await fetch(`/api/v1/browse/${type}`)

      const list: IList = await response.json()

      return list
    }

    getBrowseView()
      .then(list => setList(list))

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


function PlayerWrapper() {
  const { id } = useParams<{ id: string }>()

  return <Player id={parseInt(id)} />
}

export default App;
