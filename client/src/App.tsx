import {
  BrowserRouter,
  Switch,
  Route,
  useParams
} from "react-router-dom"

import Player from "./Player";

import styled from "styled-components"

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
        <Route path="/">
          <div>Welcome home!</div>
        </Route>
      </Switch>
    </BrowserRouter>
  </AppStyle>
}


function PlayerWrapper() {
  const { id } = useParams<{ id: string }>()

  return <Player id={parseInt(id)} />
}

export default App;
