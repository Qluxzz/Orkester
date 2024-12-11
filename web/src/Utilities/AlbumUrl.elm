module Utilities.AlbumUrl exposing (..)


albumImageUrl : { r | id : Int } -> String
albumImageUrl { id } =
    "/api/v1/album/" ++ String.fromInt id ++ "/image"
