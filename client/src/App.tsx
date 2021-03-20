import React, { useEffect, useRef, useState } from "react";
import PlaybackButton from "./PlaybackButton";
import { IPlayBackState } from "./types/playbackState";


function App() {
  const [playbackState, setPlaybackState] = useState<IPlayBackState>("paused")
  const audio = new Audio("http://localhost:3000/stream")

  useEffect(() => {
    async function togglePlayback() {
      switch (playbackState) {
        case "paused":
          if (!audio.paused)
            audio.pause()
          break
        case "playing":
          await audio.play()
          break
      }
    }

    togglePlayback()
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
    </div>
  );
}



export default App;
