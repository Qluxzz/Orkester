module Api.LikedTracks exposing (Album, Artist, Track, tracksDecoder)

import Json.Decode as Decode
import Json.Decode.Pipeline
import Parser
import Types.TrackId
import Utilities.Date


type alias Track =
    { id : Types.TrackId.TrackId
    , trackNumber : Int
    , title : String
    , length : Int
    , liked : Bool
    , artists : List Artist
    , dateAdded : Utilities.Date.Date
    , album : Album
    }


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    }


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    }


tracksDecoder : Decode.Decoder (List Track)
tracksDecoder =
    Decode.list trackDecoder


trackDecoder : Decode.Decoder Track
trackDecoder =
    Decode.succeed Track
        |> Json.Decode.Pipeline.required "id" Types.TrackId.trackIdDecoder
        |> Json.Decode.Pipeline.required "trackNumber" Decode.int
        |> Json.Decode.Pipeline.required "title" Decode.string
        |> Json.Decode.Pipeline.required "length" Decode.int
        |> Json.Decode.Pipeline.required "liked" Decode.bool
        |> Json.Decode.Pipeline.required "artists" (Decode.list artistDecoder)
        |> Json.Decode.Pipeline.required "dateAdded" Utilities.Date.dateDecoder
        |> Json.Decode.Pipeline.required "album" albumDecoder


artistDecoder : Decode.Decoder Artist
artistDecoder =
    Decode.succeed Artist
        |> Json.Decode.Pipeline.required "id" Decode.int
        |> Json.Decode.Pipeline.required "name" Decode.string
        |> Json.Decode.Pipeline.required "urlName" Decode.string


albumDecoder : Decode.Decoder Album
albumDecoder =
    Decode.succeed Album
        |> Json.Decode.Pipeline.required "id" Decode.int
        |> Json.Decode.Pipeline.required "name" Decode.string
        |> Json.Decode.Pipeline.required "urlName" Decode.string
