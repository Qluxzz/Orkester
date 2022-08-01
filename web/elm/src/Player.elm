port module Player exposing (Msg(..), playTrack, playbackFailed)


type Msg
    = PlaybackFailed String
    | PlayTrack { id : Int, timestamp : Int }



-- PORTS


port playbackFailed : (String -> msg) -> Sub msg


port playTrack : { id : Int, timestamp : Int } -> Cmd msg
