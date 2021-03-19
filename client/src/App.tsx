import { useEffect, useRef, useState } from "react";
import styled from "styled-components"



type IPlayBackState = "playing" | "paused"

function App() {
  const [playbackState, setPlaybackState] = useState<IPlayBackState>("paused")
  const audioRef = useRef<HTMLAudioElement>(null)

  useEffect(() => {
    switch (playbackState) {
      case "paused":
        audioRef?.current?.pause()
        break
      case "playing":
        audioRef?.current?.play()
        break
    }
  }, [playbackState])

  function togglePlayback() {
    setPlaybackState(
      currentState => {
        switch (currentState) {
          case "playing":
            return "paused"
          case "paused":
            return "playing"
        }
      }
    )
  }

  return (
    <div className="App">
      <PlaybackButton
        playbackState={playbackState}
        togglePlayback={togglePlayback}
      />
      <audio ref={audioRef} src="http://localhost:3000/stream" />
    </div>
  );
}

type IButton = {
  playbackState: IPlayBackState
}

const Button = styled.button<IButton>`
    width: 50%;
    height: 40px;
    border-radius: 0;
    border: 0;
    background-color: red;
    
    ${props => props.playbackState === "playing" && `
      background-color: green;
      color: white;
    `}
  `

function PlaybackButton({
  playbackState,
  togglePlayback
}: { playbackState: IPlayBackState, togglePlayback: () => void }) {
    return <Button
      onClick={togglePlayback}
      playbackState={playbackState}
    >
    {playbackState}
  </Button>
}

export default App;
