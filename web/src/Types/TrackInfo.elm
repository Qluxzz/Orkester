module Types.TrackInfo exposing (Track, trackInfoDecoder)

import Json.Decode exposing (Decoder, bool, int, list, string, succeed)
import Json.Decode.Pipeline exposing (required)
import Types.TrackId exposing (TrackId, trackIdDecoder)


type alias Track =
    { id : TrackId
    , title : String
    , length : Int
    , liked : Bool
    , album : Album
    , artists : List Artist
    }


type alias Artist =
    { id : Int, name : String, urlName : String }


type alias Album =
    { id : Int, name : String, urlName : String }


trackInfoDecoder : Decoder Track
trackInfoDecoder =
    succeed Track
        |> required "id" trackIdDecoder
        |> required "title" string
        |> required "length" int
        |> required "liked" bool
        |> required "album" albumDecoder
        |> required "artists" (list artistDecoder)


artistDecoder : Decoder Artist
artistDecoder =
    succeed Artist
        |> required "id" int
        |> required "name" string
        |> required "urlName" string


albumDecoder : Decoder Album
albumDecoder =
    succeed Album
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
