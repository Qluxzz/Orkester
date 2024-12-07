module Utilities.ArtistUrl exposing (..)


artistUrl : { r | id : Int, urlName : String } -> String
artistUrl artist =
    "/artist/" ++ String.fromInt artist.id ++ "/" ++ artist.urlName
