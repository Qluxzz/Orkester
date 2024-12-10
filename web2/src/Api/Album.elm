module Api.Album exposing (..)

import Json.Decode
import Json.Decode.Pipeline exposing (required)
import Types.ReleaseDate exposing (ReleaseDate, releaseDateDecoder)
import Types.TrackId


albumDecoder : Json.Decode.Decoder Album
albumDecoder =
    Json.Decode.succeed Album
        |> required "id" Json.Decode.int
        |> required "name" Json.Decode.string
        |> required "urlName" Json.Decode.string
        |> required "tracks" (Json.Decode.list trackDecoder)
        |> required "released" releaseDateDecoder
        |> required "artist" artistDecoder


trackDecoder : Json.Decode.Decoder Track
trackDecoder =
    Json.Decode.succeed Track
        |> required "id" Types.TrackId.trackIdDecoder
        |> required "trackNumber" Json.Decode.int
        |> required "title" Json.Decode.string
        |> required "length" Json.Decode.int
        |> required "liked" Json.Decode.bool
        |> required "artists" (Json.Decode.list artistDecoder)


artistDecoder : Json.Decode.Decoder Artist
artistDecoder =
    Json.Decode.succeed Artist
        |> required "id" Json.Decode.int
        |> required "name" Json.Decode.string
        |> required "urlName" Json.Decode.string


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , tracks : List Track
    , released : ReleaseDate
    , artist : Artist
    }


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    }


type alias Seconds =
    Int


type alias Track =
    { id : Types.TrackId.TrackId
    , trackNumber : Int
    , title : String
    , length : Seconds
    , liked : Bool
    , artists : List Artist
    }
