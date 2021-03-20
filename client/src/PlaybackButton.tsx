import styled from "styled-components"
import { IPlayBackState } from "./types/playbackState"

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

type IPlaybackButton = {
    playbackState: IPlayBackState,
    togglePlayback: () => void
}

export default function PlaybackButton({
    playbackState,
    togglePlayback
}: IPlaybackButton) {
    return <Button
        onClick={togglePlayback}
        playbackState={playbackState}
    >
        {playbackState}
    </Button>
}