module Unlike exposing (..)

import ApiBaseUrl exposing (apiBaseUrl)
import Http
import RemoteData exposing (WebData)
import TrackId exposing (TrackId)


type Msg
    = UnlikeTrackResponse TrackId (WebData ())


unlikeTrackById : TrackId -> Cmd Msg
unlikeTrackById trackId =
    Http.request
        { method = "DELETE"
        , headers = []
        , url = apiBaseUrl ++ "/api/v1/track/" ++ String.fromInt trackId ++ "/like"
        , body = Http.emptyBody
        , expect = Http.expectWhatever (RemoteData.fromResult >> UnlikeTrackResponse trackId)
        , timeout = Nothing
        , tracker = Nothing
        }
