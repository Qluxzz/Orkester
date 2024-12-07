module Utilities.AlbumUrl exposing (..)


albumUrl : { r | id : Int, urlName : String } -> String
albumUrl { id, urlName } =
    "/album/" ++ String.fromInt id ++ "/" ++ urlName


albumImageUrl : { r | id : Int } -> String
albumImageUrl { id } =
    "/api/v1/album/" ++ String.fromInt id ++ "/image"
