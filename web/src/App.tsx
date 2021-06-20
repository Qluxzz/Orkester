import React from "react"

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
import { useState } from "react"
import ITrack from "types/track"
import usePlayer from "hooks/usePlayer"
import { useCallback } from "react"

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

async function fetchTrackDetails(id: number) {
  const response = await fetch(`/api/v1/track/${id}`, {
    headers: {
      "Content-Type": "application/json"
    }
  })

  if (!response.ok)
    throw new Error(`Http request failed with status code: ${response.status}`)

  const track: ITrack = await response.json()

  return track
}

function App() {
  const [currentTrack, setCurrentTrack] = useState<ITrack>()
  const { play, pause, playTrack, seek, progress, playbackState } = usePlayer()

  const playTrackById = useCallback((id: number) => {
    fetchTrackDetails(id)
      .then(track => setCurrentTrack(track))

    playTrack(id)
  }, [playTrack])


  return <Container>
    <BrowserRouter>
      <Content>
        <SideBar />
        <MainContent>
          <Route path="/search/:query" children={() => <SearchBar />} />
          <ScrollableContent>
            <Switch>
              <Route path="/album/:id">
                <AlbumViewWrapper
                  play={playTrackById}
                />
              </Route>
              <Route path="/artist/:id">
                <ArtistViewWrapper />
              </Route>
              <Route path="/search/:query">
                <SearchViewWrapper play={playTrackById} />
              </Route>
              <Route path="/collection/tracks">
                <LikedTracks play={playTrackById} />
              </Route>
            </Switch>
          </ScrollableContent>
        </MainContent>
      </Content>
      <PlayerBar
        play={play}
        pause={pause}
        track={currentTrack}
        seek={seek}
        duration={progress.duration}
        currentTime={progress.currentTime}
        playbackState={playbackState}
      />
    </BrowserRouter>
  </Container>
}

function AlbumViewWrapper({ play }: { play: (id: number) => void }) {
  const { id } = useParams<{ id: string }>()

  return <GetAlbumById id={parseInt(id)} play={play} />
}

function ArtistViewWrapper() {
  const { id } = useParams<{ id: string }>()

  return <GetArtistWithId id={parseInt(id)} />
}

function SearchViewWrapper({ play }: { play: (id: number) => void }) {
  const { query } = useParams<{ query: string }>()

  return <SearchResults query={query} play={play} />
}

export default App;
