module Api.Search exposing (..)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (required)
import Types.TrackInfo



type alias SearchResult =
    { albums : List Album
    , artists : List Artist
    , tracks : List Types.TrackInfo.Track
    }

type alias IdNameAndUrlName =
    { id : Int
    , name : String
    , urlName : String
    }

type alias Album =
    IdNameAndUrlName


type alias Artist =
    IdNameAndUrlName


albumDecoder : Decoder Album
albumDecoder =
    Decode.succeed IdNameAndUrlName
        |> required "id" Decode.int
        |> required "name" Decode.string
        |> required "urlName" Decode.string


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed IdNameAndUrlName
        |> required "id" Decode.int
        |> required "name" Decode.string
        |> required "urlName" Decode.string


searchResultDecoder : Decoder SearchResult
searchResultDecoder =
    Decode.succeed SearchResult
        |> required "albums" (Decode.list albumDecoder)
        |> required "artists" (Decode.list artistDecoder)
        |> required "tracks" (Decode.list Types.TrackInfo.trackInfoDecoder)