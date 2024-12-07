module Api.Artist exposing (..)

import Json.Decode exposing (Decoder, int, list, string, succeed)
import Json.Decode.Pipeline exposing (required)
import Types.ReleaseDate exposing (ReleaseDate, releaseDateDecoder)


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    , albums : List Album
    }


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , released : ReleaseDate
    }


artistDecoder : Decoder Artist
artistDecoder =
    succeed Artist
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
        |> required "albums" (list albumDecoder)


albumDecoder : Decoder Album
albumDecoder =
    succeed Album
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
        |> required "released" releaseDateDecoder
