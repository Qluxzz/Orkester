port module JSPlayer exposing (Msg(..), playTrack, playbackFailed, progressUpdated, seek)


type Msg
    = PlaybackFailed String
    | PlayTrack { id : Int, timestamp : Int }
    | Seek { timestamp : Int }
    | ProgressUpdated Int



-- PORTS


port playbackFailed : (String -> msg) -> Sub msg


port playTrack : { id : Int, timestamp : Int } -> Cmd msg


port seek : { timestamp : Int } -> Cmd msg


port progressUpdated : (Int -> msg) -> Sub msg
