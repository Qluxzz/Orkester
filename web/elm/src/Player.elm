port module Player exposing (Msg(..), playTrack, playbackFailed, progressUpdated)


type Msg
    = PlaybackFailed String
    | PlayTrack { id : Int, timestamp : Int }
    | ProgressUpdated Int



-- PORTS


port playbackFailed : (String -> msg) -> Sub msg


port playTrack : { id : Int, timestamp : Int } -> Cmd msg


port progressUpdated : (Int -> msg) -> Sub msg
