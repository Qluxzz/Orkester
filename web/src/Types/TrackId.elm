module Types.TrackId exposing (TrackId, toString, trackIdDecoder)

import Json.Decode


type TrackId
    = TrackId String


trackIdDecoder : Json.Decode.Decoder TrackId
trackIdDecoder =
    Json.Decode.string
        |> Json.Decode.andThen
            (\s ->
                if String.startsWith "track-" s then
                    Json.Decode.succeed (TrackId s)

                else
                    Json.Decode.fail "This is not a track id"
            )


toString : TrackId -> String
toString (TrackId str) =
    str
