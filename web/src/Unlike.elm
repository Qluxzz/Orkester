module Unlike exposing (..)

import Http
import RemoteData exposing (WebData)
import Types.TrackId exposing (TrackId)
import Utilities.ApiBaseUrl exposing (apiBaseUrl)


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
