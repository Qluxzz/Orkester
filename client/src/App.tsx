import React, { useEffect, useRef, useState } from "react";
import {
  BrowserRouter,
  Switch,
  Route,
  Link,
  useParams
} from "react-router-dom"

import Player from "./Player";


function App() {
  return <BrowserRouter>
    <Switch>
      <Route path="/stream/:id" children={<PlayerWrapper />} />
      <Route path="/">
        <div>Welcome home!</div>
      </Route>
    </Switch>
  
  </BrowserRouter>
}


function PlayerWrapper() {
  const { id } = useParams<{ id: string }>()

  return <Player id={parseInt(id)} />
}

export default App;
