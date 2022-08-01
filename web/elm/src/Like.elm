module Like exposing (Msg(..), likeTrackById)

import ApiBaseUrl exposing (apiBaseUrl)
import Http
import RemoteData exposing (WebData)
import TrackId exposing (TrackId)


type Msg
    = LikeTrackResponse TrackId (WebData ())


likeTrackById : TrackId -> Cmd Msg
likeTrackById trackId =
    Http.request
        { method = "PUT"
        , headers = []
        , url = apiBaseUrl ++ "/api/v1/track/" ++ String.fromInt trackId ++ "/like"
        , body = Http.emptyBody
        , expect = Http.expectWhatever (RemoteData.fromResult >> LikeTrackResponse trackId)
        , timeout = Nothing
        , tracker = Nothing
        }
