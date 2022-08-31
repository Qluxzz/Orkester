port module JSPlayer exposing (Msg(..), pause, play, playTrack, playbackFailed, progressUpdated, seek)


type Msg
    = PlaybackFailed String
    | PlayTrack { id : Int }
    | Seek { timestamp : Int }
    | ProgressUpdated Int
    | Pause
    | Play



-- PORTS


port playbackFailed : (String -> msg) -> Sub msg


port playTrack : Int -> Cmd msg


port seek : { timestamp : Int } -> Cmd msg


port progressUpdated : (Int -> msg) -> Sub msg


port play : () -> Cmd msg


port pause : () -> Cmd msg
