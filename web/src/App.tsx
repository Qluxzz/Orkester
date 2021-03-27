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

const AppStyle = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`


function App() {
  return <AppStyle>
    <BrowserRouter>
      <Switch>
        <Route path="/track/:id" children={<PlayerWrapper />} />
        <Route path="/browse/:type" children={<BrowseWrapper />} />
        <Route path="/">
          <Link to="/track/80">Press here plz</Link>
          <div>Welcome home!</div>
        </Route>
      </Switch>
    </BrowserRouter>
  </AppStyle>
}

function BrowseWrapper() {
  const { type } = useParams<{ type: "artists" | "genres" }>()
  const [list, setList] = useState<{ name: string, urlname: string }[]>([])

  console.log(type)

  useEffect(() => {
    async function getBrowseView() {
      const response = await fetch(`/api/v1/browse/${type}`)

      return await response.json()
    }

    getBrowseView()
      .then(list => setList(list))

  }, [type])

  return <div>
    <h1>{type}</h1>
    <ol>
    {list.map(({ name, urlname }) => {
      <Link to={`/api/v1/browse/${type}/${urlname}`}>
        <li>{name}</li>
      </Link>
    })}
    </ol>
  </div>
}


function PlayerWrapper() {
  const { id } = useParams<{ id: string }>()

  return <Player id={parseInt(id)} />
}

export default App;
