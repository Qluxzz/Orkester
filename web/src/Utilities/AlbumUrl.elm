module Utilities.AlbumUrl exposing (..)

import Utilities.ApiBaseUrl exposing (apiBaseUrl)


albumUrl : { r | id : Int, urlName : String } -> String
albumUrl { id, urlName } =
    "/album/" ++ String.fromInt id ++ "/" ++ urlName


albumImageUrl : { r | id : Int } -> String
albumImageUrl { id } =
    apiBaseUrl ++ "/api/v1/album/" ++ String.fromInt id ++ "/image"
