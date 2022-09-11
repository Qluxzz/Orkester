port module JSPlayer exposing (Msg(..), pause, play, playTrack, playbackFailed, progressUpdated, seek, stateChange)


type Msg
    = PlaybackFailed String
    | PlayTrack { id : Int }
    | Seek { timestamp : Int }
    | ProgressUpdated Int
      -- Changes to the JavaScript Audio object
      -- Can be that the user used shortcuts to play next/prev, pause/play
    | ExternalStateChange String



-- PORTS
-- OUT


port playTrack : Int -> Cmd msg


port seek : { timestamp : Int } -> Cmd msg


port play : () -> Cmd msg


port pause : () -> Cmd msg



-- IN


port playbackFailed : (String -> msg) -> Sub msg


port progressUpdated : (Int -> msg) -> Sub msg


port stateChange : (String -> msg) -> Sub msg
